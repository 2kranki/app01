// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ma contains all the functions
// and data to interact with the SQL Database.

// Generated: Thu Nov 14, 2019 11:17 for mariadb Database

package ioApp01ma

import (
	"testing"
)


var ioApp01ma   *IO_App01ma


//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                              Connect
//----------------------------------------------------------------------------

func TestApp01maConnect(t *testing.T) {
    var err         error

	t.Logf("TestConnect()...\n")
// DockerRun(t)

	ioApp01ma = NewIoApp01ma()
	ioApp01ma.DefaultParms()
    err = ioApp01ma.Connect("")
    if err == nil {
	    err = ioApp01ma.Disconnect()
        if err != nil {
            t.Fatalf("Error: %s\n\n", err)
        }
        ioApp01ma = nil
    } else {
            t.Fatalf("Error: %s\n\n", err)
    }

	t.Logf("TestConnect() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Disconnect
//----------------------------------------------------------------------------

func TestApp01maDisconnect(t *testing.T) {
    var err         error

	t.Logf("TestDisconnect()...\n")
	ioApp01ma = NewIoApp01ma()
	ioApp01ma.DefaultParms()

	// Disconnect before a connection has been made.
    err = ioApp01ma.Disconnect()
    if err == nil {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    if ioApp01ma.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    // Now connect then disconnect.
    err = ioApp01ma.Connect("")
    if err != nil {
        t.Fatal("Error: Cannot connect: ", err)
    }

    if !ioApp01ma.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    err = ioApp01ma.Disconnect()
    if err != nil {
        t.Fatal("Error: Cannot disconnect: ", err)
    }
    ioApp01ma = nil

	t.Logf("TestDisconnect() - End of Test\n\n\n")
}



//----------------------------------------------------------------------------
//                              DatabaseCreate
//----------------------------------------------------------------------------

func TestApp01maDatabaseCreate(t *testing.T) {
    var err         error

	t.Logf("TestDatabaseCreate()...\n")
//DockerRun(t)

	ioApp01ma = NewIoApp01ma()
	ioApp01ma.DefaultParms()

    err = ioApp01ma.DatabaseCreate("App01ma")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01ma.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01ma = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
//----------------------------------------------------------------------------
//                              QueryRow
//----------------------------------------------------------------------------

func TestApp01maQueryRow(t *testing.T) {
    var err         error

	t.Logf("TestQueryRow()...\n")
	//DockerRun(t)

	ioApp01ma = NewIoApp01ma()
	ioApp01ma.DefaultParms()

    err = ioApp01ma.DatabaseCreate("App01ma")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01ma.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01ma = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
