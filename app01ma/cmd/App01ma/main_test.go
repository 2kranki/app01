// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ma contains all the functions
// and data to interact with the SQL Database.

// Generated: Sat Nov 23, 2019 00:27 for mariadb Database

package main

import (
	"testing"
	"time"
	//"github.com/2kranki/go_util"
)

//----------------------------------------------------------------------------
//                              Docker Run - mariadb
//----------------------------------------------------------------------------

// DockerRun executes the dbs/mariadb/run.sh to create a fresh SQL Server.
func DockerRun(t *testing.T) {
	var err error
	var exec *util.ExecCmd
	var output string

	t.Logf("DockerRun()...\n")

	exec = util.NewExecArgs("", "../dbs/mariadb/run.sh")
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
