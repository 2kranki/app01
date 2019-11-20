// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ms contains all the functions
// and data to interact with the SQL Database.

// Generated: Wed Nov 20, 2019 16:06 for mssql Database

package main

import (
	"testing"
    "time"
	"github.com/2kranki/go_util"
)


//----------------------------------------------------------------------------
//                              Docker Run - mssql
//----------------------------------------------------------------------------

// DockerRun executes the dbs/mssql/run.sh to create a fresh SQL Server.
func DockerRun(t *testing.T) {
    var err         error
    var exec        *util.ExecCmd
    var output      string

	t.Logf("DockerRun()...\n")

	exec = util.NewExecArgs("", "../dbs/mssql/run.sh")
	if exec == nil {
        t.Fatalf("Error: Failed to create util.ExecCmd instance!\n\n")
	}

    output, err = exec.RunWithOutput()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    t.Logf("%s\n", output)
    time.Sleep(5000 * time.Millisecond)

	t.Logf("DockerRun() - End\n\n\n")
}


//----------------------------------------------------------------------------
//                         Test Certificate Creation
//----------------------------------------------------------------------------

func TestCertApp01ms(t *testing.T) {
    var err         error

	t.Logf("TestCertApp01ms()...\n")

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

	t.Logf("TestCertApp01ms() - End\n\n\n")
}

