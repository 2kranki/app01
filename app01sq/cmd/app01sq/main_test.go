// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01sq contains all the functions
// and data to interact with the SQL Database.

// Generated: Wed Nov 20, 2019 16:06 for sqlite Database

package main

import (
	"testing"
    
	"github.com/2kranki/go_util"
)


//----------------------------------------------------------------------------
//                              Docker Run - sqlite
//----------------------------------------------------------------------------

// DockerRun executes the dbs/sqlite/run.sh to create a fresh SQL Server.
func DockerRun(t *testing.T) {

	t.Logf("DockerRun()...\n")

	t.Logf("DockerRun() - End\n\n\n")
}

//----------------------------------------------------------------------------
//                         Test Certificate Creation
//----------------------------------------------------------------------------

func TestCertApp01sq(t *testing.T) {
    var err         error

	t.Logf("TestCertApp01sq()...\n")

    certDir = "/tmp/certs"
    certDirPath := util.NewPath(certDir)
    err = certDirPath.RemoveDir()
    t.Logf("\tError: (Can be ignored) Deleting directory, %s - %s",
            certDirPath.String(), util.ErrorString(err))

    genCerts()

    certPath := certDirPath.Append("cert.pem")
    if !certPath.IsPathRegularFile() {
        t.Errorf("\tError: Missing %s!\n", certPath.String())
    }

    keyPath := certDirPath.Append("key.pem")
    if !keyPath.IsPathRegularFile() {
        t.Errorf("\tError: Missing %s!\n", keyPath.String())
    }

    // Clean up
    err = certDirPath.RemoveDir()
    t.Logf("\tError: (Can be ignored) Deleting directory, %s - %s",
            certDirPath.String(), util.ErrorString(err))

	t.Logf("TestCertApp01sq() - End\n\n\n")
}

