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

// Generated: Mon Jan  6, 2020 09:54

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"app01sq/pkg/hndlrApp01sq"

	"app01sq/pkg/hndlrApp01sqCustomer"
	"app01sq/pkg/hndlrApp01sqVendor"
	"app01sq/pkg/httpServer"
	"app01sq/pkg/ioApp01sq"
	"app01sq/pkg/ioApp01sqCustomer"
	"app01sq/pkg/ioApp01sqVendor"
)

const (
	RowsPerPage = 15
)

var hndlrsApp01sq *hndlrApp01sq.TmplsApp01sq

var hndlrsApp01sqCustomer *hndlrApp01sqCustomer.HandlersApp01sqCustomer
var hndlrsApp01sqVendor *hndlrApp01sqVendor.HandlersApp01sqVendor

var app01sqIO *ioApp01sq.IO_App01sq

var app01sqCustomerIO *ioApp01sqCustomer.IO_App01sqCustomer
var app01sqVendorIO *ioApp01sqVendor.IO_App01sqVendor

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

	// App01sq.Customer URL handlers for table maintenance
	hndlrsApp01sqCustomer = hndlrApp01sqCustomer.NewHandlersApp01sqCustomer(app01sqCustomerIO, RowsPerPage, h.Mux)
	hndlrsApp01sqCustomer.Tmpls = hndlrsApp01sq
	if hndlrsApp01sqCustomer.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01sq\n\n\n")
	}
	// App01sq.Vendor URL handlers for table maintenance
	hndlrsApp01sqVendor = hndlrApp01sqVendor.NewHandlersApp01sqVendor(app01sqVendorIO, RowsPerPage, h.Mux)
	hndlrsApp01sqVendor.Tmpls = hndlrsApp01sq
	if hndlrsApp01sqVendor.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01sq\n\n\n")
	}

	// Start the HTTP Server.
	h.Serve(true)

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
	go func() {
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
