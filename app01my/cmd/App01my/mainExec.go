// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Wed Nov 20, 2019 16:06

package main

import (
    "fmt"
    "log"
	"net/http"
	"os"
    "os/signal"

    "app01my/pkg/hndlrApp01my"
	
        "app01my/pkg/hndlrApp01myCustomer"
        "app01my/pkg/ioApp01myCustomer"
        "app01my/pkg/hndlrApp01myVendor"
        "app01my/pkg/ioApp01myVendor"
    "app01my/pkg/ioApp01my"
    "github.com/2kranki/go_util"
)

const (
    RowsPerPage = 15
)

var     hndlrsApp01my    *hndlrApp01my.TmplsApp01my

    var hndlrsApp01myCustomer     *hndlrApp01myCustomer.HandlersApp01myCustomer
    var hndlrsApp01myVendor     *hndlrApp01myVendor.HandlersApp01myVendor

var app01myIO *ioApp01my.IO_App01my

    var app01myCustomerIO  *ioApp01myCustomer.IO_App01myCustomer
    var app01myVendorIO  *ioApp01myVendor.IO_App01myVendor


type PemControl struct {
	certPath    *util.Path
	certPem     *util.Path
	keyPem      *util.Path
}

func (p *PemControl) CertPemPath() string {
    return p.certPem.String()
}

func (p *PemControl) KeyPemPath() string {
    return p.KeyPem.String()
}

// Gen generates the Certificates needed for HTTPS.
func (p *PemControl) Gen() {
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

    certPem = certPath.Append("cert.pem")
    if certPem == nil {
        log.Fatalf("Error: Creating %s/cert.pem path\n\n", certPath.String())
    }
    keyPem = certPath.Append("key.pem")
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

// certsIsPresent checks to see if the Certificates needed for HTTPS
// are present. If certificates seem ok, nil is returned. Otherwise,
// an error is returned.
func (p *PemControl) IsPresent(force bool) error {
    var err         error
    var out         string

    log.Printf("\tChecking for HTTPS Certificates...\n")
    if p.certPem.String() == "" {
        return fmt.Errorf("Error: Missing cert certificate path!\n\n")
    }
    if p.keyPem.String() == "" {
        return fmt.Errorf("Error: Missing key certificate path!\n\n")
    }

    if p.certPem.IsPathRegularFile() && p.keyPem.IsPathRegularFile() && !force {
        return nil
    }

    return fmt.Errorf("Error: Certificates need to be (re)built!\n\n")
}

// Setup checks to see if the Certificates needed for HTTPS
// are present.
func (p *PemControl) Setup() error {
    var err         error
    var out         string

    log.Printf("\tSetting up for the HTTPS Certificates...\n")
    if certDir == "" {
        return fmt.Errorf("Error: Missing certificate path!\n\n")
    }

    log.Printf("\tChecking for HTTPS Certificates in %s...\n", certDir)
    p.certPath = util.NewPath(certDir)
    if p.certPath == nil {
        return fmt.Errorf("Error: Creating %s path\n\n", certPath.String())
    }
    if err = p.certPath.CreateDir(); err != nil {
        return fmt.Errorf("Error: Create %s : %s\n\n", certPath.String(), err.Error())
    }

    p.certPem = p.certPath.Append("cert.pem")
    if certPem == nil {
        return fmt.Errorf("Error: Creating %s/cert.pem path\n\n", certPath.String())
    }
    p.keyPem = p.certPath.Append("key.pem")
    if keyPem == nil {
        return fmt.Errorf("Error: Creating %s/key.pem path\n\n", certPath.String())
    }

    return nil
}

func NewPem() *PemControl {
    return &PemControl{}
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

    fmt.Printf("\tHndlrHome Serving File: ./html/App01my.menu.html\n")
    hndlrsApp01my.MainDisplay(w, "")
    //http.ServeFile(w, r, baseDir+"/html/App01my.menu.html")

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

	
	    // App01my.Customer URL handlers for table maintenance
	    hndlrsApp01myCustomer = hndlrApp01myCustomer.NewHandlersApp01myCustomer(app01myCustomerIO, RowsPerPage, mux)
	    hndlrsApp01myCustomer.Tmpls = hndlrsApp01my
        if hndlrsApp01myCustomer.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01my\n\n\n")
        }
	    // App01my.Vendor URL handlers for table maintenance
	    hndlrsApp01myVendor = hndlrApp01myVendor.NewHandlersApp01myVendor(app01myVendorIO, RowsPerPage, mux)
	    hndlrsApp01myVendor.Tmpls = hndlrsApp01my
        if hndlrsApp01myVendor.Tmpls == nil {
            log.Fatalf("ERROR - Failed to load templates from hndlrsApp01my\n\n\n")
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
    go func(){
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