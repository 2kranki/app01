// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Sun Nov 17, 2019 06:49

package main

import (
    "fmt"
    "log"
	"net/http"
	"os"
    "os/exec"
    "os/signal"

    "app01ms/pkg/hndlrApp01ms"
	
        "app01ms/pkg/hndlrApp01msCustomer"
        "app01ms/pkg/ioApp01msCustomer"
        "app01ms/pkg/hndlrApp01msVendor"
        "app01ms/pkg/ioApp01msVendor"
    "app01ms/pkg/ioApp01ms"
    "github.com/2kranki/go_util"
)

const (
    RowsPerPage = 15
)

var     hndlrsApp01ms    *hndlrApp01ms.TmplsApp01ms

    var hndlrsApp01msCustomer     *hndlrApp01msCustomer.HandlersApp01msCustomer
    var hndlrsApp01msVendor     *hndlrApp01msVendor.HandlersApp01msVendor

var app01msIO *ioApp01ms.IO_App01ms

    var app01msCustomerIO  *ioApp01msCustomer.IO_App01msCustomer
    var app01msVendorIO  *ioApp01msVendor.IO_App01msVendor


// genCerts generates the Certificates needed for HTTPS.
func genCerts() {
    var err         error
    var cmd         string

    log.Printf("\tGenerating HTTPS Certificates if needed...\n")
    if certDir == "" {
        log.Fatalf("Error: Missing certificate path!\n\n")
    }

    log.Printf("\tChecking for HTTPS Certificates in %s...\n", certDir)
    certPath := util.NewPath(certDir)
    if certPath == nil {
        log.Fatalf("Error: Creating %s path\n\n", certPath.String())
    }
    if err = certPath.CreateDir(); err != nil {
        log.Fatalf("Error: Create %s : %s\n\n", certPath.String(), err.Error())
    }

    certPem := certPath.Append("cert.pem")
    if certPem == nil {
        log.Fatalf("Error: Creating %s/cert.pem path\n\n", certPath.String())
    }
    keyPem := certPath.Append("key.pem")
    if keyPem == nil {
        log.Fatalf("Error: Creating %s/key.pem path\n\n", certPath.String())
    }
    if certPem.IsPathRegularFile() && keyPem.IsPathRegularFile() {
        return
    }

    log.Printf("\tMissing HTTPS Certificates will now be generated...\n")
    // NOTE - The cmd to create the certificates may need to be massaged for
    //      a more specific installation.
    cmd  = "openssl req -x509 -nodes -days 365 -newkey rsa:2048 "
    cmd += fmt.Sprintf("-keyout %s -out %s", keyPem, certPem)
    cmd += " -passout pass:xyzzy"
    cmd += " -subj \"/C=US/ST=Florida/L=Tampa/O=De/OU=Dev/CN=example.com\""
    cmdr := exec.Command(cmd)
    if cmdr == nil {
        log.Fatalf("Error: Could not create command object!\n")
    }
    err = cmdr.Run()
    if err != nil {
        log.Fatalf("Error: Could not create HTTPS Certificates : %s!\n",
                    err.Error())
    }
    if certPem.IsPathRegularFile() && keyPem.IsPathRegularFile() {
        return
    }

    log.Fatalf("Error: OpenSSL could not create the certificates!\n")
}

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

    fmt.Printf("\tHndlrHome Serving File: ./html/App01ms.menu.html\n")
    hndlrsApp01ms.MainDisplay(w, "")
    //http.ServeFile(w, r, baseDir+"/html/App01ms.menu.html")

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

func mainExec() {

    // Generate HTTPS Certificates if needed.
    genCerts()

    // Setup the I/O.
    setupIO()

    // Set up templates.
    setupTmpls()

    // Set up default URL handlers
    log.Printf("\tSetting up the Mux Handlers...\n")
    mux := http.NewServeMux()
	mux.HandleFunc("/", HndlrHome)
	mux.HandleFunc("/favicon.ico", HndlrFavIcon)

	
	    // App01ms.Customer URL handlers for table maintenance
	    hndlrsApp01msCustomer = hndlrApp01msCustomer.NewHandlersApp01msCustomer(app01msCustomerIO, RowsPerPage, mux)
	    hndlrsApp01msCustomer.Tmpls = hndlrsApp01ms
        if hndlrsApp01msCustomer.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01ms\n\n\n")
        }
	    // App01ms.Vendor URL handlers for table maintenance
	    hndlrsApp01msVendor = hndlrApp01msVendor.NewHandlersApp01msVendor(app01msVendorIO, RowsPerPage, mux)
	    hndlrsApp01msVendor.Tmpls = hndlrsApp01ms
        if hndlrsApp01msVendor.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01ms\n\n\n")
        }

	// Start the HTTP Server.
    log.Printf("\tStarting Server at %s:%s...\n", http_srvr, http_port)
	srvrStr := fmt.Sprintf("%s:%s", http_srvr, http_port)
    s := &http.Server{
            Addr:    srvrStr,
            Handler: MuxHandlerWrapper(mux),
        }
        log.Fatal(s.ListenAndServe())
    

}

// setupIO connects to the datatbase.
func setupIO() {

    // Connect the databases.
    log.Printf("\tConnecting to the Database...\n")
    app01msIO = ioApp01ms.NewIoApp01ms()
    //app01msIO.SetName(db_name)
    app01msIO.SetPort(db_port)
    app01msIO.SetPW(db_pw)
    app01msIO.SetPort(db_port)
    app01msIO.SetServer(db_srvr)
    app01msIO.SetUser(db_user)
    err := app01msIO.DatabaseCreate(db_name)
    if err != nil {
        log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
    }

    // Set up to disconnect the database upon program interrupt.
    chnl := make(chan os.Signal, 1)
    signal.Notify(chnl, os.Interrupt)
    go func(){
        <-chnl
        if app01msIO.IsConnected() {
            err = app01msIO.Disconnect()
            if err != nil {
                log.Fatal(err)
            }
        }
        os.Exit(1)
    }()

    // Set up the Table I/O.
	
	    app01msCustomerIO = ioApp01msCustomer.NewIoApp01msCustomer(app01msIO)
        if app01msCustomerIO == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01msCustomer\n\n\n")
        }
	    app01msVendorIO = ioApp01msVendor.NewIoApp01msVendor(app01msIO)
        if app01msVendorIO == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01msVendor\n\n\n")
        }

}

func setupTmpls() {

    log.Printf("\tSetting up the Templates...\n")
    hndlrsApp01ms = hndlrApp01ms.NewTmplsApp01ms(baseDir + "/tmpl")
    hndlrsApp01ms.SetupTmpls()

}