// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Notes:
//  1.  When working with package main, please keep in mind that the
//      more functionality that you can move into functions, the easier
//      testing will be. This allows you to test the functionality in
//      small portions. Moving common functionality to packages that are
//      easily tested is even better.
//  2.  All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Mon Jan  6, 2020 11:09

package main

import (
    "fmt"
    "log"
	"net/http"
	"os"
    "os/signal"

    "app01ms/pkg/hndlrApp01ms"
	
        "app01ms/pkg/hndlrApp01msCustomer"
        "app01ms/pkg/ioApp01msCustomer"
        "app01ms/pkg/hndlrApp01msVendor"
        "app01ms/pkg/ioApp01msVendor"
    "app01ms/pkg/ioApp01ms"
    "app01ms/pkg/httpServer"
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

func mainExec() {

    h := httpServer.NewHttp(http_srvr, http_port, https_port)
    if h == nil {
        log.Fatalf("Error: Unable to create HTTP/HTTPS server!\n")
    }
    /*
        err := h.SetupCerts(certDir)
        if err != nil {
            log.Fatalf("Error: Unable to create HTTPS Certificates!\n")
        }
    */

    // Setup the I/O.
    setupIO()

    // Set up templates.
    setupTmpls()

    // Set up default URL handlers
    log.Printf("\tSetting up the Mux Handlers...\n")
	h.Mux.HandleFunc("/", HndlrHome)
	h.Mux.HandleFunc("/favicon.ico", HndlrFavIcon)

	
	    // App01ms.Customer URL handlers for table maintenance
	    hndlrsApp01msCustomer = hndlrApp01msCustomer.NewHandlersApp01msCustomer(app01msCustomerIO, RowsPerPage, h.Mux)
	    hndlrsApp01msCustomer.Tmpls = hndlrsApp01ms
        if hndlrsApp01msCustomer.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01ms\n\n\n")
        }
	    // App01ms.Vendor URL handlers for table maintenance
	    hndlrsApp01msVendor = hndlrApp01msVendor.NewHandlersApp01msVendor(app01msVendorIO, RowsPerPage, h.Mux)
	    hndlrsApp01msVendor.Tmpls = hndlrsApp01ms
        if hndlrsApp01msVendor.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01ms\n\n\n")
        }

	// Start the HTTP Server.
h.Serve(true)

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