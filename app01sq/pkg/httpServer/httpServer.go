// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// HTTP/HTTPS Server Package

// This package sets up either a single HTTP mux server which
// handles all requests or an HTTP mux server that redirects
// all traffic to an HTTPS mux server and the HTTPS mux server.
// So, in the first case, one mux server is created. In the
// second case, two servers are created.

// An example of how to use this package is:
//  h := NewHttp("localhost", "80", "443")
//  h.SetupCerts("/tmp/certs")          // <== if using HTTPS
//  h.Mux.HandleFunc("/", HndlrHome)
//  h.Serve(false)      // <== true == wrap handlers for debugging
//  NOTE: Any code here will never be executed.

// Generated: Mon Jan  6, 2020 11:09

package httpServer

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/2kranki/go_util"
	"app01sq/pkg/cert"
)

//----------------------------------------------------------------------------
//                         HTTP/HTTPS Server
//----------------------------------------------------------------------------

type HttpServer struct {
	CertDir    string
	Host       string
	Port       string
	SecureHost string
	SecurePort string
	Mux        *http.ServeMux
	Certs      *cert.CertControl // if != nil, assume HTTPS
}

// HndlrRedirect redirects all HTTP requests to HTTPS requests.
func (h *HttpServer) HndlrRedirect(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("HndlrRedirect(%s)\n", r.Method)

	redirectUrl := fmt.Sprintf("https://%s:%s%s", h.SecureHost, h.SecurePort,
		r.URL.Path)
	http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)

	fmt.Printf("...end HndlrRedirect()\n")
}

// Serve sets up the server(s) which handle the HTTP/HTTPS
// requests.
func (h *HttpServer) Serve(UseMuxWrapper bool) error {
	var serverString string

	if h.Certs != nil {
		log.Printf("\tStarting HTTP Redirect Server at %s:%s -> %s:%s...\n",
			h.Host, h.Port, h.SecureHost, h.SecurePort)
		serverString = fmt.Sprintf("%s:%s", h.Host, h.Port)
		go http.ListenAndServe(serverString, http.HandlerFunc(h.HndlrRedirect))
	}

	if h.Certs != nil {
		log.Printf("\tStarting HTTPS Server at %s:%s...\n",
			h.SecureHost, h.SecurePort)
		serverString = fmt.Sprintf("%s:%s", h.SecureHost, h.SecurePort)
	} else {
		log.Printf("\tStarting HTTP Server at %s:%s...\n", h.Host, h.Port)
		serverString = fmt.Sprintf("%s:%s", h.Host, h.Port)
	}
	if UseMuxWrapper {
		s := &http.Server{
			Addr:    serverString,
			Handler: MuxHandlerWrapper(h.Mux),
		}
		log.Fatal(s.ListenAndServe())
	} else {
		log.Fatal(http.ListenAndServe(serverString, h.Mux))
	}

	return nil
}

// SetupCerts insures that the HTTPS Certificates are present
// creating them if needed. It initializes the Certs variable
// which indicates that we will be using HTTPS instead of HTTP.
// Warning: This must be run before Serve().
func (h *HttpServer) SetupCerts(certDir string) error {

	// Generate HTTPS Certificates if needed.
	h.Certs = cert.NewCert(certDir)
	if h.Certs.IsPresent(false) != nil {
		err := h.Certs.Generate()
		if err != nil {
			return err
		}
	}

	return nil
}

//----------------------------------------------------------------------------
//                         New HTTP/HTTPS Server
//----------------------------------------------------------------------------

func NewHttp(host, port, securePort string) *HttpServer {

	h := &HttpServer{}
	h.Host = host
	h.Port = port
	h.SecureHost = host
	h.SecurePort = securePort
	h.Mux = http.NewServeMux()

	return h
}

//----------------------------------------------------------------------------
//               HTTP/HTTPS Server Handlers and Intercepts
//----------------------------------------------------------------------------

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
