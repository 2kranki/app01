// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01pg contains all the functions
// and data to interact with the SQL Database.

// Generated: Sat Nov 23, 2019 00:27 for postgres Database

package cert

import (
	"github.com/2kranki/go_util"
	"testing"
	"time"
)

//----------------------------------------------------------------------------
//                         Test Certificate Creation
//----------------------------------------------------------------------------

func TestCertApp01pg(t *testing.T) {
	var err error
	var c *CertControl
	var tmpDir string = "/tmp/certs"

	t.Logf("TestCertApp01pg()...\n")
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

	t.Logf("TestCertApp01pg() - End\n\n\n")
}
