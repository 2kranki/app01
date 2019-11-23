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

// Generated: Sat Nov 23, 2019 00:27

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"app01ma/pkg/hndlrApp01ma"

	"app01ma/pkg/hndlrApp01maCustomer"
	"app01ma/pkg/hndlrApp01maVendor"
	"app01ma/pkg/httpServer"
	"app01ma/pkg/ioApp01ma"
	"app01ma/pkg/ioApp01maCustomer"
	"app01ma/pkg/ioApp01maVendor"
)

const (
	RowsPerPage = 15
)

var hndlrsApp01ma *hndlrApp01ma.TmplsApp01ma

var hndlrsApp01maCustomer *hndlrApp01maCustomer.HandlersApp01maCustomer
var hndlrsApp01maVendor *hndlrApp01maVendor.HandlersApp01maVendor

var app01maIO *ioApp01ma.IO_App01ma

var app01maCustomerIO *ioApp01maCustomer.IO_App01maCustomer
var app01maVendorIO *ioApp01maVendor.IO_App01maVendor

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

	fmt.Printf("\tHndlrHome Serving File: ./html/App01ma.menu.html\n")
	hndlrsApp01ma.MainDisplay(w, "")
	//http.ServeFile(w, r, baseDir+"/html/App01ma.menu.html")

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

	// App01ma.Customer URL handlers for table maintenance
	hndlrsApp01maCustomer = hndlrApp01maCustomer.NewHandlersApp01maCustomer(app01maCustomerIO, RowsPerPage, h.Mux)
	hndlrsApp01maCustomer.Tmpls = hndlrsApp01ma
	if hndlrsApp01maCustomer.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01ma\n\n\n")
	}
	// App01ma.Vendor URL handlers for table maintenance
	hndlrsApp01maVendor = hndlrApp01maVendor.NewHandlersApp01maVendor(app01maVendorIO, RowsPerPage, h.Mux)
	hndlrsApp01maVendor.Tmpls = hndlrsApp01ma
	if hndlrsApp01maVendor.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01ma\n\n\n")
	}

	// Start the HTTP Server.
	h.Serve(true)

}

// setupIO connects to the datatbase.
func setupIO() {

	// Connect the databases.
	log.Printf("\tConnecting to the Database...\n")
	app01maIO = ioApp01ma.NewIoApp01ma()
	//app01maIO.SetName(db_name)
	app01maIO.SetPort(db_port)
	app01maIO.SetPW(db_pw)
	app01maIO.SetPort(db_port)
	app01maIO.SetServer(db_srvr)
	app01maIO.SetUser(db_user)
	err := app01maIO.DatabaseCreate(db_name)
	if err != nil {
		log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
	}

	// Set up to disconnect the database upon program interrupt.
	chnl := make(chan os.Signal, 1)
	signal.Notify(chnl, os.Interrupt)
	go func() {
		<-chnl
		if app01maIO.IsConnected() {
			err = app01maIO.Disconnect()
			if err != nil {
				log.Fatal(err)
			}
		}
		os.Exit(1)
	}()

	// Set up the Table I/O.

	app01maCustomerIO = ioApp01maCustomer.NewIoApp01maCustomer(app01maIO)
	if app01maCustomerIO == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01maCustomer\n\n\n")
	}
	app01maVendorIO = ioApp01maVendor.NewIoApp01maVendor(app01maIO)
	if app01maVendorIO == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01maVendor\n\n\n")
	}

}

func setupTmpls() {

	log.Printf("\tSetting up the Templates...\n")
	hndlrsApp01ma = hndlrApp01ma.NewTmplsApp01ma(baseDir + "/tmpl")
	hndlrsApp01ma.SetupTmpls()

}
