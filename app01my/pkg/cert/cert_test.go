// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01my contains all the functions
// and data to interact with the SQL Database.

// Generated: Mon Jan  6, 2020 11:09 for mysql Database

package cert

import (
	"testing"
    "time"
	"github.com/2kranki/go_util"
)

//----------------------------------------------------------------------------
//                         Test Certificate Creation
//----------------------------------------------------------------------------

func TestCertApp01my(t *testing.T) {
    var err         error
    var c           *CertControl
    var tmpDir      string  = "/tmp/certs"

	t.Logf("TestCertApp01my()...\n")
	c = NewCert(tmpDir)
	if c == nil {
        t.Fatalf("Error: Could not create CertControl object!\n")
	}

    err = c.IsPresent(true)
	if err == nil {
        t.Fatalf("Error: IsPresent(true) is nil!\n")
	}

    err = c.Generate()
	if err != nil {
        t.Fatalf("Error: Generate: %s!\n", err.Error())
	}

    if !c.CertPem().IsPathRegularFile() {
        t.Errorf("\tError: Missing %s!\n", c.CertPemPath())
    }

    if !c.KeyPem().IsPathRegularFile() {
        t.Errorf("\tError: Missing %s!\n", c.KeyPemPath())
    }

    // Clean up
    err = util.NewPath(tmpDir).RemoveDir()
	if err != nil {
        t.Logf("Clean up error: %s!\n", err.Error())
	}

	t.Logf("TestCertApp01my() - End\n\n\n")
}

