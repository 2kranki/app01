// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01sq contains all the functions
// and data to interact with the SQL Database.

// Generated: Mon Oct 28, 2019 08:40 for sqlite Database

package ioApp01sq

import (
	"testing"
)


var ioApp01sq   *IO_App01sq


//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                              Connect
//----------------------------------------------------------------------------

func TestApp01sqConnect(t *testing.T) {
    var err         error

	t.Logf("TestConnect()...\n")


	ioApp01sq = NewIoApp01sq()
	ioApp01sq.DefaultParms()
    err = ioApp01sq.Connect("")
    if err == nil {
	    err = ioApp01sq.Disconnect()
        if err != nil {
            t.Fatalf("Error: %s\n\n", err)
        }
        ioApp01sq = nil
    } else {
            t.Fatalf("Error: %s\n\n", err)
    }

	t.Logf("TestConnect() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Disconnect
//----------------------------------------------------------------------------

func TestApp01sqDisconnect(t *testing.T) {
    var err         error

	t.Logf("TestDisconnect()...\n")
	ioApp01sq = NewIoApp01sq()
	ioApp01sq.DefaultParms()

	// Disconnect before a connection has been made.
    err = ioApp01sq.Disconnect()
    if err == nil {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    if ioApp01sq.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    // Now connect then disconnect.
    err = ioApp01sq.Connect("")
    if err != nil {
        t.Fatal("Error: Cannot connect: ", err)
    }

    if !ioApp01sq.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    err = ioApp01sq.Disconnect()
    if err != nil {
        t.Fatal("Error: Cannot disconnect: ", err)
    }
    ioApp01sq = nil

	t.Logf("TestDisconnect() - End of Test\n\n\n")
}



//----------------------------------------------------------------------------
//                              DatabaseCreate
//----------------------------------------------------------------------------

func TestApp01sqDatabaseCreate(t *testing.T) {
    var err         error

	t.Logf("TestDatabaseCreate()...\n")


	ioApp01sq = NewIoApp01sq()
	ioApp01sq.DefaultParms()

    err = ioApp01sq.DatabaseCreate("App01sq")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01sq.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01sq = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
//----------------------------------------------------------------------------
//                              QueryRow
//----------------------------------------------------------------------------

func TestApp01sqQueryRow(t *testing.T) {
    var err         error

	t.Logf("TestQueryRow()...\n")
	//DockerRun(t)

	ioApp01sq = NewIoApp01sq()
	ioApp01sq.DefaultParms()

    err = ioApp01sq.DatabaseCreate("App01sq")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01sq.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01sq = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
