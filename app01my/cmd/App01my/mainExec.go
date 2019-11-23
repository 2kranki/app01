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

	"app01my/pkg/hndlrApp01my"

	"app01my/pkg/hndlrApp01myCustomer"
	"app01my/pkg/hndlrApp01myVendor"
	"app01my/pkg/httpServer"
	"app01my/pkg/ioApp01my"
	"app01my/pkg/ioApp01myCustomer"
	"app01my/pkg/ioApp01myVendor"
)

const (
	RowsPerPage = 15
)

var hndlrsApp01my *hndlrApp01my.TmplsApp01my

var hndlrsApp01myCustomer *hndlrApp01myCustomer.HandlersApp01myCustomer
var hndlrsApp01myVendor *hndlrApp01myVendor.HandlersApp01myVendor

var app01myIO *ioApp01my.IO_App01my

var app01myCustomerIO *ioApp01myCustomer.IO_App01myCustomer
var app01myVendorIO *ioApp01myVendor.IO_App01myVendor

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

	// App01my.Customer URL handlers for table maintenance
	hndlrsApp01myCustomer = hndlrApp01myCustomer.NewHandlersApp01myCustomer(app01myCustomerIO, RowsPerPage, h.Mux)
	hndlrsApp01myCustomer.Tmpls = hndlrsApp01my
	if hndlrsApp01myCustomer.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01my\n\n\n")
	}
	// App01my.Vendor URL handlers for table maintenance
	hndlrsApp01myVendor = hndlrApp01myVendor.NewHandlersApp01myVendor(app01myVendorIO, RowsPerPage, h.Mux)
	hndlrsApp01myVendor.Tmpls = hndlrsApp01my
	if hndlrsApp01myVendor.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01my\n\n\n")
	}

	// Start the HTTP Server.
	h.Serve(true)

}

// setupIO connects to the datatbase.
func setupIO() {

	// Connect the databases.
	log.Printf("\tConnecting to the Database...\n")
	app01myIO = ioApp01my.NewIoApp01my()
	//app01myIO.SetName(db_name)
	app01myIO.SetPort(db_port)
	app01myIO.SetPW(db_pw)
	app01myIO.SetPort(db_port)
	app01myIO.SetServer(db_srvr)
	app01myIO.SetUser(db_user)
	err := app01myIO.DatabaseCreate(db_name)
	if err != nil {
		log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
	}

	// Set up to disconnect the database upon program interrupt.
	chnl := make(chan os.Signal, 1)
	signal.Notify(chnl, os.Interrupt)
	go func() {
		<-chnl
		if app01myIO.IsConnected() {
			err = app01myIO.Disconnect()
			if err != nil {
				log.Fatal(err)
			}
		}
		os.Exit(1)
	}()

	// Set up the Table I/O.

	app01myCustomerIO = ioApp01myCustomer.NewIoApp01myCustomer(app01myIO)
	if app01myCustomerIO == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01myCustomer\n\n\n")
	}
	app01myVendorIO = ioApp01myVendor.NewIoApp01myVendor(app01myIO)
	if app01myVendorIO == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01myVendor\n\n\n")
	}

}

func setupTmpls() {

	log.Printf("\tSetting up the Templates...\n")
	hndlrsApp01my = hndlrApp01my.NewTmplsApp01my(baseDir + "/tmpl")
	hndlrsApp01my.SetupTmpls()

}
