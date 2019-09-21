// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01sq contains all the functions
// and data to interact with the SQL Database.

// Generated: Sat Sep 21, 2019 06:41 for sqlite Database

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
	var err error
	var exec *util.ExecCmd

	t.Logf("DockerRun()...\n")

	exec = util.NewExecCmd("../dbs/sqlite/run.sh")
	if exec == nil {
		t.Fatalf("Error: Failed to create util.ExecCmd instance!\n\n")
	}

	//err = exec.Run()      // Not needed for sqlite!
	if err != nil {
		t.Fatalf("Error: %s\n\n", err)
	}

	t.Logf("DockerRun() - End\n\n\n")
}
