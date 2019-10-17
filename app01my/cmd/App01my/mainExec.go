// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Thu Oct 17, 2019 11:49

package main

import (
    "fmt"
    "log"
	"net/http"
	"os"
    "os/signal"
)

const (
    RowsPerPage = 15
)

var     hndlrsApp01my    *TmplsApp01my
	
	    var hndlrsApp01myCustomer     *HandlersApp01myCustomer
	    var hndlrsApp01myVendor     *HandlersApp01myVendor

// HndlrFavIcon is the default Favorite Icon Handler.  It defaults to
// returning a 405 status to indicate that no Icon is available.
func HndlrFavIcon(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("HndlrFavIcon(%s)\n", r.Method)

    if r.Method != "GET" {
	    http.NotFound(w, r)
	}
    http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)

    fmt.Printf("...end HndlrFavIcon(Error:405)\n")
}

// HndlrHome responds to a URL with no sub-elements.  It defaults to
// providing the default Menu to the browser/caller.
func HndlrHome(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("HndlrHome(%s)\n", r.Method)

    if r.URL.Path != "/" {
        fmt.Printf("...end HndlrHome(Error 404) Not '/' URL\n")
        http.NotFound(w, r)
        return
    }

    fmt.Printf("\tHndlrHome Serving File: ./html/App01my.menu.html\n")
    hndlrsApp01my.MainDisplay(w, "")
    //http.ServeFile(w, r, baseDir+"/html/App01my.menu.html")

    fmt.Printf("...end HndlrHome()\n")
}

// To understand the following, review packages net/http and net/url and review:
// https://stackoverflow.com/questions/25456390/how-to-log-http-server-errors-in-golang

// MuxResponseWriterWrap provides a wrapper around the Response
// Writer so that we can intercept data being written out if
// needed.
type MuxResponseWriterWrap struct {
	http.ResponseWriter
	status int
}

func (r *MuxResponseWriterWrap) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (r *MuxResponseWriterWrap) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// MuxHandlerWrapper will intercept each mux request and give
// us access both, before and after, the request is handled.
func MuxHandlerWrapper(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		record := &MuxResponseWriterWrap{
			ResponseWriter: w,
		}

        // Intercept before the request is handled.
		log.Println("mux input: (", r.Method, ") ", r.URL.String())

		f.ServeHTTP(record, r)

        // Intercept after the request is handled.
		log.Println("Bad Request ", record.status)

		if record.status == http.StatusBadRequest {
			log.Println("Bad Request ", r)
		}
	}
}

func exec() {

    // Connect the databases.
    log.Printf("\tConnecting to the Database...\n")
    ioApp01my := NewIoApp01my()
    //ioApp01my.SetName(db_name)
    ioApp01my.SetPort(db_port)
    ioApp01my.SetPW(db_pw)
    ioApp01my.SetPort(db_port)
    ioApp01my.SetServer(db_srvr)
    ioApp01my.SetUser(db_user)
    err := ioApp01my.DatabaseCreate(db_name)
    if err != nil {
        log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
    }
    chnl := make(chan os.Signal, 1)
    signal.Notify(chnl, os.Interrupt, os.Kill)
    go func(){
        <-chnl
        if ioApp01my.IsConnected() {
            ioApp01my.Disconnect()
        }
        os.Exit(1)
    }()

    // Set up the Table I/O.
	
	    ioApp01myCustomer := NewIoApp01myCustomer(ioApp01my)
        if ioApp01myCustomer == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01myCustomer\n\n\n")
        }
	    ioApp01myVendor := NewIoApp01myVendor(ioApp01my)
        if ioApp01myVendor == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01myVendor\n\n\n")
        }

    // Set up templates.
    log.Printf("\tSetting up the Templates...\n")
    hndlrsApp01my = NewTmplsApp01my("")
    hndlrsApp01my.SetTmplsDir(baseDir + "/tmpl")
    hndlrsApp01my.SetupTmpls()

    // Set up default URL handlers
    log.Printf("\tSetting up the Mux Handlers...\n")
    mux := http.NewServeMux()
	mux.HandleFunc("/", HndlrHome)
	mux.HandleFunc("/favicon.ico", HndlrFavIcon)

	
	    // App01my.Customer URL handlers for table maintenance
	    hndlrsApp01myCustomer = NewHandlersApp01myCustomer(ioApp01myCustomer, RowsPerPage, mux)
	    hndlrsApp01myCustomer.Tmpls = hndlrsApp01my
        if hndlrsApp01myCustomer.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01my\n\n\n")
        }
	    // App01my.Vendor URL handlers for table maintenance
	    hndlrsApp01myVendor = NewHandlersApp01myVendor(ioApp01myVendor, RowsPerPage, mux)
	    hndlrsApp01myVendor.Tmpls = hndlrsApp01my
        if hndlrsApp01myVendor.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01my\n\n\n")
        }

    // mkdir ssl
    // openssl req -x509 -days 365 -nodes -newkey rsa:2048 -keyout ./ssl/ssl_key.pem -out ./ssl/ssl_cert.pem

	// Start the HTTP Server.
    log.Printf("\tStarting Server at %s:%s...\n", http_srvr, http_port)
	srvrStr := fmt.Sprintf("%s:%s", http_srvr, http_port)
    s := &http.Server{
            Addr:    srvrStr,
            Handler: MuxHandlerWrapper(mux),
        }
        log.Fatal(s.ListenAndServe())
    

}

