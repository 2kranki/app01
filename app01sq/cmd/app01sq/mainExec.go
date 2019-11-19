// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Tue Nov 19, 2019 15:46

package main

import (
    "fmt"
    "log"
	"net/http"
	"os"
    "os/signal"

    "app01sq/pkg/hndlrApp01sq"
	
        "app01sq/pkg/hndlrApp01sqCustomer"
        "app01sq/pkg/ioApp01sqCustomer"
        "app01sq/pkg/hndlrApp01sqVendor"
        "app01sq/pkg/ioApp01sqVendor"
    "app01sq/pkg/ioApp01sq"
    "github.com/2kranki/go_util"
)

const (
    RowsPerPage = 15
)

var     hndlrsApp01sq    *hndlrApp01sq.TmplsApp01sq

    var hndlrsApp01sqCustomer     *hndlrApp01sqCustomer.HandlersApp01sqCustomer
    var hndlrsApp01sqVendor     *hndlrApp01sqVendor.HandlersApp01sqVendor

var app01sqIO *ioApp01sq.IO_App01sq

    var app01sqCustomerIO  *ioApp01sqCustomer.IO_App01sqCustomer
    var app01sqVendorIO  *ioApp01sqVendor.IO_App01sqVendor


// genCerts generates the Certificates needed for HTTPS.
func genCerts() {
    var err         error
    var out         string

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
    if certPem.IsPathRegularFile() && keyPem.IsPathRegularFile() && !force {
        return
    }

    log.Printf("\tMissing HTTPS Certificates will now be generated...\n")
    // NOTE - The cmd to create the certificates may need to be massaged for
    //      a more specific installation.
    //TODO: Allow for password to be substituted.
    cmd := util.NewExecArgs("openssl", "req", "-x509", "-nodes",
     "-days", "365", "-newkey", "rsa:2048", "-keyout", keyPem.String(),
     "-out", certPem.String(), "-passout", "pass:xyzzy",
     "-subj", "/C=US/ST=Florida/L=Tampa/O=De/OU=Dev/CN=example.com")
    log.Printf("\tExecuting %s...\n", cmd.CommandString())
    if cmd == nil {
        log.Fatalf("Error: Could not create cmd object!\n")
    }
    out, err = cmd.RunWithOutput()
    if err != nil {
        log.Fatalf("Error: Did not create HTTPS Certificates : %s : %s!\n",
                    err.Error(), out)
    }
    if certPem.IsPathRegularFile() && keyPem.IsPathRegularFile() {
        return
    }

    log.Fatalf("Error: OpenSSL did not create the certificates!\n")
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

	
	    // App01sq.Customer URL handlers for table maintenance
	    hndlrsApp01sqCustomer = hndlrApp01sqCustomer.NewHandlersApp01sqCustomer(app01sqCustomerIO, RowsPerPage, mux)
	    hndlrsApp01sqCustomer.Tmpls = hndlrsApp01sq
        if hndlrsApp01sqCustomer.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01sq\n\n\n")
        }
	    // App01sq.Vendor URL handlers for table maintenance
	    hndlrsApp01sqVendor = hndlrApp01sqVendor.NewHandlersApp01sqVendor(app01sqVendorIO, RowsPerPage, mux)
	    hndlrsApp01sqVendor.Tmpls = hndlrsApp01sq
        if hndlrsApp01sqVendor.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01sq\n\n\n")
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
    app01sqIO = ioApp01sq.NewIoApp01sq()
    //app01sqIO.SetName(db_name)
    app01sqIO.SetPort(db_port)
    app01sqIO.SetPW(db_pw)
    app01sqIO.SetPort(db_port)
    app01sqIO.SetServer(db_srvr)
    app01sqIO.SetUser(db_user)
    err := app01sqIO.DatabaseCreate(db_name)
    if err != nil {
        log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
    }

    // Set up to disconnect the database upon program interrupt.
    chnl := make(chan os.Signal, 1)
    signal.Notify(chnl, os.Interrupt)
    go func(){
        <-chnl
        if app01sqIO.IsConnected() {
            err = app01sqIO.Disconnect()
            if err != nil {
                log.Fatal(err)
            }
        }
        os.Exit(1)
    }()

    // Set up the Table I/O.
	
	    app01sqCustomerIO = ioApp01sqCustomer.NewIoApp01sqCustomer(app01sqIO)
        if app01sqCustomerIO == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01sqCustomer\n\n\n")
        }
	    app01sqVendorIO = ioApp01sqVendor.NewIoApp01sqVendor(app01sqIO)
        if app01sqVendorIO == nil {
            log.Fatalf("ERROR - Failed to Connect to Table, App01sqVendor\n\n\n")
        }

}

func setupTmpls() {

    log.Printf("\tSetting up the Templates...\n")
    hndlrsApp01sq = hndlrApp01sq.NewTmplsApp01sq(baseDir + "/tmpl")
    hndlrsApp01sq.SetupTmpls()

}