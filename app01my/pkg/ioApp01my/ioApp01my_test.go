// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01my contains all the functions
// and data to interact with the SQL Database.

// Generated: Sun Nov 17, 2019 06:49 for mysql Database

package ioApp01my

import (
	"testing"
)


var ioApp01my   *IO_App01my


//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                              Connect
//----------------------------------------------------------------------------

func TestApp01myConnect(t *testing.T) {
    var err         error

	t.Logf("TestConnect()...\n")
// DockerRun(t)

	ioApp01my = NewIoApp01my()
	ioApp01my.DefaultParms()
    err = ioApp01my.Connect("")
    if err == nil {
	    err = ioApp01my.Disconnect()
        if err != nil {
            t.Fatalf("Error: %s\n\n", err)
        }
        ioApp01my = nil
    } else {
            t.Fatalf("Error: %s\n\n", err)
    }

	t.Logf("TestConnect() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Disconnect
//----------------------------------------------------------------------------

func TestApp01myDisconnect(t *testing.T) {
    var err         error

	t.Logf("TestDisconnect()...\n")
	ioApp01my = NewIoApp01my()
	ioApp01my.DefaultParms()

	// Disconnect before a connection has been made.
    err = ioApp01my.Disconnect()
    if err == nil {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    if ioApp01my.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    // Now connect then disconnect.
    err = ioApp01my.Connect("")
    if err != nil {
        t.Fatal("Error: Cannot connect: ", err)
    }

    if !ioApp01my.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    err = ioApp01my.Disconnect()
    if err != nil {
        t.Fatal("Error: Cannot disconnect: ", err)
    }
    ioApp01my = nil

	t.Logf("TestDisconnect() - End of Test\n\n\n")
}



//----------------------------------------------------------------------------
//                              DatabaseCreate
//----------------------------------------------------------------------------

func TestApp01myDatabaseCreate(t *testing.T) {
    var err         error

	t.Logf("TestDatabaseCreate()...\n")
//DockerRun(t)

	ioApp01my = NewIoApp01my()
	ioApp01my.DefaultParms()

    err = ioApp01my.DatabaseCreate("App01my")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01my.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01my = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
//----------------------------------------------------------------------------
//                              QueryRow
//----------------------------------------------------------------------------

func TestApp01myQueryRow(t *testing.T) {
    var err         error

	t.Logf("TestQueryRow()...\n")
	//DockerRun(t)

	ioApp01my = NewIoApp01my()
	ioApp01my.DefaultParms()

    err = ioApp01my.DatabaseCreate("App01my")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01my.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01my = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
