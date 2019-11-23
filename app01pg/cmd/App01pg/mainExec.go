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

	"app01pg/pkg/hndlrApp01pg"

	"app01pg/pkg/hndlrApp01pgCustomer"
	"app01pg/pkg/hndlrApp01pgVendor"
	"app01pg/pkg/httpServer"
	"app01pg/pkg/ioApp01pg"
	"app01pg/pkg/ioApp01pgCustomer"
	"app01pg/pkg/ioApp01pgVendor"
)

const (
	RowsPerPage = 15
)

var hndlrsApp01pg *hndlrApp01pg.TmplsApp01pg

var hndlrsApp01pgCustomer *hndlrApp01pgCustomer.HandlersApp01pgCustomer
var hndlrsApp01pgVendor *hndlrApp01pgVendor.HandlersApp01pgVendor

var app01pgIO *ioApp01pg.IO_App01pg

var app01pgCustomerIO *ioApp01pgCustomer.IO_App01pgCustomer
var app01pgVendorIO *ioApp01pgVendor.IO_App01pgVendor

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

	fmt.Printf("\tHndlrHome Serving File: ./html/App01pg.menu.html\n")
	hndlrsApp01pg.MainDisplay(w, "")
	//http.ServeFile(w, r, baseDir+"/html/App01pg.menu.html")

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

	// App01pg.Customer URL handlers for table maintenance
	hndlrsApp01pgCustomer = hndlrApp01pgCustomer.NewHandlersApp01pgCustomer(app01pgCustomerIO, RowsPerPage, h.Mux)
	hndlrsApp01pgCustomer.Tmpls = hndlrsApp01pg
	if hndlrsApp01pgCustomer.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01pg\n\n\n")
	}
	// App01pg.Vendor URL handlers for table maintenance
	hndlrsApp01pgVendor = hndlrApp01pgVendor.NewHandlersApp01pgVendor(app01pgVendorIO, RowsPerPage, h.Mux)
	hndlrsApp01pgVendor.Tmpls = hndlrsApp01pg
	if hndlrsApp01pgVendor.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01pg\n\n\n")
	}

	// Start the HTTP Server.
	h.Serve(true)

}

// setupIO connects to the datatbase.
func setupIO() {

	// Connect the databases.
	log.Printf("\tConnecting to the Database...\n")
	app01pgIO = ioApp01pg.NewIoApp01pg()
	//app01pgIO.SetName(db_name)
	app01pgIO.SetPort(db_port)
	app01pgIO.SetPW(db_pw)
	app01pgIO.SetPort(db_port)
	app01pgIO.SetServer(db_srvr)
	app01pgIO.SetUser(db_user)
	err := app01pgIO.DatabaseCreate(db_name)
	if err != nil {
		log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
	}

	// Set up to disconnect the database upon program interrupt.
	chnl := make(chan os.Signal, 1)
	signal.Notify(chnl, os.Interrupt)
	go func() {
		<-chnl
		if app01pgIO.IsConnected() {
			err = app01pgIO.Disconnect()
			if err != nil {
				log.Fatal(err)
			}
		}
		os.Exit(1)
	}()

	// Set up the Table I/O.

	app01pgCustomerIO = ioApp01pgCustomer.NewIoApp01pgCustomer(app01pgIO)
	if app01pgCustomerIO == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01pgCustomer\n\n\n")
	}
	app01pgVendorIO = ioApp01pgVendor.NewIoApp01pgVendor(app01pgIO)
	if app01pgVendorIO == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01pgVendor\n\n\n")
	}

}

func setupTmpls() {

	log.Printf("\tSetting up the Templates...\n")
	hndlrsApp01pg = hndlrApp01pg.NewTmplsApp01pg(baseDir + "/tmpl")
	hndlrsApp01pg.SetupTmpls()

}
