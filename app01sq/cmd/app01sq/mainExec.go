// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Wed Oct 16, 2019 20:04

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

var     hndlrsApp01sq    *TmplsApp01sq
	
	    var hndlrsApp01sqCustomer     *HandlersApp01sqCustomer
	    var hndlrsApp01sqVendor     *HandlersApp01sqVendor

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

    fmt.Printf("\tHndlrHome Serving File: ./html/App01sq.menu.html\n")
    hndlrsApp01sq.MainDisplay(w, "")
    //http.ServeFile(w, r, baseDir+"/html/App01sq.menu.html")

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
    ioApp01sq := NewIoApp01sq()
    //ioApp01sq.SetName(db_name)
    ioApp01sq.SetPort(db_port)
    ioApp01sq.SetPW(db_pw)
    ioApp01sq.SetPort(db_port)
    ioApp01sq.SetServer(db_srvr)
    ioApp01sq.SetUser(db_user)
    err := ioApp01sq.DatabaseCreate(db_name)
    if err != nil {
        log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
    }
    chnl := make(chan os.Signal, 1)
    signal.Notify(chnl, os.Interrupt, os.Kill)
    go func(){
        <-chnl
        if ioApp01sq.IsConnected() {
            ioApp01sq.Disconnect()
        }
        os.Exit(1)
    }()

    // Set up the Table I/O.
	
	    ioApp01sqCustomer := NewIoApp01sqCustomer(ioApp01sq)
        if ioApp01sqCustomer == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01sqCustomer\n\n\n")
        }
	    ioApp01sqVendor := NewIoApp01sqVendor(ioApp01sq)
        if ioApp01sqVendor == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01sqVendor\n\n\n")
        }

    // Set up templates.
    log.Printf("\tSetting up the Templates...\n")
    hndlrsApp01sq = NewTmplsApp01sq("")
    hndlrsApp01sq.SetTmplsDir(baseDir + "/tmpl")
    hndlrsApp01sq.SetupTmpls()

    // Set up default URL handlers
    log.Printf("\tSetting up the Mux Handlers...\n")
    mux := http.NewServeMux()
	mux.HandleFunc("/", HndlrHome)
	mux.HandleFunc("/favicon.ico", HndlrFavIcon)

	
	    // App01sq.Customer URL handlers for table maintenance
	    hndlrsApp01sqCustomer = NewHandlersApp01sqCustomer(ioApp01sqCustomer, RowsPerPage, mux)
	    hndlrsApp01sqCustomer.Tmpls = hndlrsApp01sq
        if hndlrsApp01sqCustomer.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01sq\n\n\n")
        }
	    // App01sq.Vendor URL handlers for table maintenance
	    hndlrsApp01sqVendor = NewHandlersApp01sqVendor(ioApp01sqVendor, RowsPerPage, mux)
	    hndlrsApp01sqVendor.Tmpls = hndlrsApp01sq
        if hndlrsApp01sqVendor.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01sq\n\n\n")
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

