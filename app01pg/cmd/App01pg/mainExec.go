// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Thu Nov 14, 2019 11:17

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

// HndlrFavIcon is the default Favorite Icon Handler.  It defaults to
// returning a 405 status to indicate that no Icon is available.
func HndlrFavIcon(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.NotFound(w, r)
	}
	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)

}

// HndlrHome responds to a URL with no sub-elements.  It defaults to
// providing the default Menu to the browser/caller.
func HndlrHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {

		http.NotFound(w, r)
		return
	}

	hndlrsApp01pg.MainDisplay(w, "")
	//http.ServeFile(w, r, baseDir+"/html/App01pg.menu.html")

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

	// Generate HTTPS Certificates if needed.

	// Connect the databases.

	ioApp01pg := ioApp01pg.NewIoApp01pg()
	//ioApp01pg.SetName(db_name)
	ioApp01pg.SetPort(db_port)
	ioApp01pg.SetPW(db_pw)
	ioApp01pg.SetPort(db_port)
	ioApp01pg.SetServer(db_srvr)
	ioApp01pg.SetUser(db_user)
	err := ioApp01pg.DatabaseCreate(db_name)
	if err != nil {
		log.Fatalf("ERROR - Failed to Connect Database\n\n\n")
	}

	// Set up to disconnect the database upon program interrupt.
	chnl := make(chan os.Signal, 1)
	signal.Notify(chnl, os.Interrupt)
	go func() {
		<-chnl
		if ioApp01pg.IsConnected() {
			err = ioApp01pg.Disconnect()
			if err != nil {
				log.Fatal(err)
			}
		}
		os.Exit(1)
	}()

	// Set up the Table I/O.

	ioApp01pgCustomer := ioApp01pgCustomer.NewIoApp01pgCustomer(ioApp01pg)
	if ioApp01pgCustomer == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01pgCustomer\n\n\n")
	}
	ioApp01pgVendor := ioApp01pgVendor.NewIoApp01pgVendor(ioApp01pg)
	if ioApp01pgVendor == nil {
		log.Fatalf("ERROR - Failed to Connect to Table, App01pgVendor\n\n\n")
	}

	// Set up templates.

	hndlrsApp01pg = hndlrApp01pg.NewTmplsApp01pg("")
	hndlrsApp01pg.SetTmplsDir(baseDir + "/tmpl")
	hndlrsApp01pg.SetupTmpls()

	// Set up default URL handlers

	mux := http.NewServeMux()
	mux.HandleFunc("/", HndlrHome)
	mux.HandleFunc("/favicon.ico", HndlrFavIcon)

	// App01pg.Customer URL handlers for table maintenance
	hndlrsApp01pgCustomer = hndlrApp01pgCustomer.NewHandlersApp01pgCustomer(ioApp01pgCustomer, RowsPerPage, mux)
	hndlrsApp01pgCustomer.Tmpls = hndlrsApp01pg
	if hndlrsApp01pgCustomer.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01pg\n\n\n")
	}
	// App01pg.Vendor URL handlers for table maintenance
	hndlrsApp01pgVendor = hndlrApp01pgVendor.NewHandlersApp01pgVendor(ioApp01pgVendor, RowsPerPage, mux)
	hndlrsApp01pgVendor.Tmpls = hndlrsApp01pg
	if hndlrsApp01pgVendor.Tmpls == nil {
		log.Fatalf("ERROR - Failed to load templates from hndlrsApp01pg\n\n\n")
	}

	// mkdir ssl
	// openssl req -x509 -days 365 -nodes -newkey rsa:2048 -keyout ./ssl/ssl_key.pem -out ./ssl/ssl_cert.pem

	// Start the HTTP Server.

	srvrStr := fmt.Sprintf("%s:%s", http_srvr, http_port)
	s := &http.Server{
		Addr:    srvrStr,
		Handler: MuxHandlerWrapper(mux),
	}
	log.Fatal(s.ListenAndServe())

}
