// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// HTTPS Certificate Package

// Generated: Wed Nov 20, 2019 16:06

package cert

import (
    "fmt"
    "log"
	//"net/http"

    "github.com/2kranki/go_util"
)



type CertControl struct {
    CertDir     string
	certPath    *util.Path
	certPem     *util.Path
	keyPem      *util.Path
}

func (c *CertControl) CertPem() *util.Path {
    return c.certPem
}

func (c *CertControl) CertPemPath() string {
    return c.certPem.String()
}

func (c *CertControl) KeyPem() *util.Path {
    return c.keyPem
}

func (c *CertControl) KeyPemPath() string {
    return c.keyPem.String()
}

// Gen generates the Certificates needed for HTTPS.
func (c *CertControl) Generate() error {
    var err         error
    var out         string

    log.Printf("\tGenerating HTTPS Certificates if needed...\n")

    log.Printf("\tChecking for HTTPS Certificates in %s...\n", c.CertDir)
    if err = c.certPath.CreateDir(); err != nil {
        return fmt.Errorf("Error: Create %s : %s\n\n", c.certPath.String(), err.Error())
    }

    log.Printf("\tMissing HTTPS Certificates will now be generated...\n")
    // NOTE - The cmd to create the certificates may need to be massaged for
    //      a more specific installation.
    //TODO: Allow for 'password' to be substituted.
    //TODO: Allow for the fields of the 'subject' to be substituted.
    cmd := util.NewExecArgs("openssl", "req", "-x509", "-nodes",
     "-days", "365", "-newkey", "rsa:2048", "-keyout", c.keyPem.String(),
     "-out", c.certPem.String(), "-passout", "pass:xyzzy",
     "-subj", "/C=US/ST=Florida/L=Tampa/O=De/OU=Dev/CN=example.com")
    log.Printf("\tExecuting %s...\n", cmd.CommandString())
    if cmd == nil {
        log.Fatalf("Error: Could not create cmd object!\n")
    }
    out, err = cmd.RunWithOutput()
    if err != nil {
            log.Printf("\tError: %s:%s\n", err.Error(), out)
        } else {
            log.Printf("\tWorked!\n")
        }
    if err != nil {
        return fmt.Errorf("Error: Did not create HTTPS Certificates : %s : %s!\n",
                    err.Error(), out)
    }
    if c.certPem.IsPathRegularFile() && c.keyPem.IsPathRegularFile() {
        return nil
    }

    return fmt.Errorf("Error: OpenSSL did not create the certificates!\n")
}

// IsPresent checks to see if the Certificates needed for HTTPS
// are present. If certificates seem ok, nil is returned. Otherwise,
// an error is returned.
func (c *CertControl) IsPresent(force bool) error {

    log.Printf("\tChecking for HTTPS Certificates...\n")

    if !c.certPath.IsPathDir() {
        return fmt.Errorf("Error: Missing cert directory path!\n\n")
    }
    if c.certPem.String() == "" {
        return fmt.Errorf("Error: Missing cert certificate path!\n\n")
    }
    if c.keyPem.String() == "" {
        return fmt.Errorf("Error: Missing key certificate path!\n\n")
    }

    if c.certPem.IsPathRegularFile() && c.keyPem.IsPathRegularFile() && !force {
        return nil
    }

    return fmt.Errorf("Error: Certificates need to be (re)built!\n\n")
}

// Setup sets up the various variables to access/generate certificates.
func (c *CertControl) Setup() error {

    log.Printf("\tSetting up for the HTTPS Certificates...\n")
    if c.CertDir == "" {
        return fmt.Errorf("Error: Missing certificate path!\n\n")
    }

    log.Printf("\tChecking for HTTPS Certificates in %s...\n", c.CertDir)
    c.certPath = util.NewPath(c.CertDir)
    if c.certPath == nil {
        return fmt.Errorf("Error: Creating %s path\n\n", c.certPath.String())
    }

    c.certPem = c.certPath.Append("cert.pem")
    if c.certPem == nil {
        return fmt.Errorf("Error: Creating %s/cert.pem path\n\n", c.certPath.String())
    }
    c.keyPem = c.certPath.Append("key.pem")
    if c.keyPem == nil {
        return fmt.Errorf("Error: Creating %s/key.pem path\n\n", c.certPath.String())
    }

    return nil
}

func NewCert(certDir string) *CertControl {
    c := &CertControl{}
    c.CertDir = certDir
    err := c.Setup()
    if err != nil {
        return nil
    }
    return c
}

