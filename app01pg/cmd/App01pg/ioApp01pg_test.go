// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01pg contains all the functions
// and data to interact with the SQL Database.

// Generated: Thu Oct 17, 2019 11:49 for postgres Database

package main

import (
	"testing"
)


var ioApp01pg   *IO_App01pg


//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                              Connect
//----------------------------------------------------------------------------

func TestApp01pgConnect(t *testing.T) {
    var err         error

	t.Logf("TestConnect()...\n")
	DockerRun(t)

	ioApp01pg = NewIoApp01pg()
	ioApp01pg.DefaultParms()
    err = ioApp01pg.Connect("")
    if err == nil {
	    err = ioApp01pg.Disconnect()
        if err != nil {
            t.Fatalf("Error: %s\n\n", err)
        }
        ioApp01pg = nil
    } else {
            t.Fatalf("Error: %s\n\n", err)
    }

	t.Logf("TestConnect() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Disconnect
//----------------------------------------------------------------------------

func TestApp01pgDisconnect(t *testing.T) {
    var err         error

	t.Logf("TestDisconnect()...\n")
	ioApp01pg = NewIoApp01pg()
	ioApp01pg.DefaultParms()

	// Disconnect before a connection has been made.
    err = ioApp01pg.Disconnect()
    if err == nil {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    if ioApp01pg.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    // Now connect then disconnect.
    err = ioApp01pg.Connect("")
    if err != nil {
        t.Fatal("Error: Cannot connect: ", err)
    }

    if !ioApp01pg.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    err = ioApp01pg.Disconnect()
    if err != nil {
        t.Fatal("Error: Cannot disconnect: ", err)
    }
    ioApp01pg = nil

	t.Logf("TestDisconnect() - End of Test\n\n\n")
}



//----------------------------------------------------------------------------
//                              DatabaseCreate
//----------------------------------------------------------------------------

func TestApp01pgDatabaseCreate(t *testing.T) {
    var err         error

	t.Logf("TestDatabaseCreate()...\n")
	DockerRun(t)

	ioApp01pg = NewIoApp01pg()
	ioApp01pg.DefaultParms()

    err = ioApp01pg.DatabaseCreate("App01pg")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01pg.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01pg = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
//----------------------------------------------------------------------------
//                              QueryRow
//----------------------------------------------------------------------------

func TestApp01pgQueryRow(t *testing.T) {
    var err         error

	t.Logf("TestQueryRow()...\n")
	//DockerRun(t)

	ioApp01pg = NewIoApp01pg()
	ioApp01pg.DefaultParms()

    err = ioApp01pg.DatabaseCreate("App01pg")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01pg.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01pg = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
