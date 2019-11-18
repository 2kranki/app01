// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ms contains all the functions
// and data to interact with the SQL Database.

// Generated: Sun Nov 17, 2019 06:49 for mssql Database

package ioApp01ms

import (
	"testing"
)


var ioApp01ms   *IO_App01ms


//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                              Connect
//----------------------------------------------------------------------------

func TestApp01msConnect(t *testing.T) {
    var err         error

	t.Logf("TestConnect()...\n")
// DockerRun(t)

	ioApp01ms = NewIoApp01ms()
	ioApp01ms.DefaultParms()
    err = ioApp01ms.Connect("")
    if err == nil {
	    err = ioApp01ms.Disconnect()
        if err != nil {
            t.Fatalf("Error: %s\n\n", err)
        }
        ioApp01ms = nil
    } else {
            t.Fatalf("Error: %s\n\n", err)
    }

	t.Logf("TestConnect() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Disconnect
//----------------------------------------------------------------------------

func TestApp01msDisconnect(t *testing.T) {
    var err         error

	t.Logf("TestDisconnect()...\n")
	ioApp01ms = NewIoApp01ms()
	ioApp01ms.DefaultParms()

	// Disconnect before a connection has been made.
    err = ioApp01ms.Disconnect()
    if err == nil {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    if ioApp01ms.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    // Now connect then disconnect.
    err = ioApp01ms.Connect("")
    if err != nil {
        t.Fatal("Error: Cannot connect: ", err)
    }

    if !ioApp01ms.IsConnected() {
        t.Fatal("Error: Never Connected!\n\n\n")
    }

    err = ioApp01ms.Disconnect()
    if err != nil {
        t.Fatal("Error: Cannot disconnect: ", err)
    }
    ioApp01ms = nil

	t.Logf("TestDisconnect() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              IsDatabaseDefined
//----------------------------------------------------------------------------

func TestApp01msIsDatabaseDefined(t *testing.T) {
    var err         error

	t.Logf("TestIsDatabaseDefined()...\n")
//DockerRun(t)

	ioApp01ms = NewIoApp01ms()
	ioApp01ms.DefaultParms()
    err = ioApp01ms.Connect("")
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }

    if !ioApp01ms.IsDatabaseDefined("App01ms") {
        err = ioApp01ms.DatabaseCreate("App01ms")
        if err != nil {
            t.Fatalf("\tError - Database Create failed: %s\n", err.Error())
        }
    }

    err = ioApp01ms.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01ms = nil

	t.Logf("TestIsDatabaseDefined() - End of Test\n\n\n")
}


//----------------------------------------------------------------------------
//                              DatabaseCreate
//----------------------------------------------------------------------------

func TestApp01msDatabaseCreate(t *testing.T) {
    var err         error

	t.Logf("TestDatabaseCreate()...\n")
//DockerRun(t)

	ioApp01ms = NewIoApp01ms()
	ioApp01ms.DefaultParms()

    err = ioApp01ms.DatabaseCreate("App01ms")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01ms.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01ms = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
//----------------------------------------------------------------------------
//                              QueryRow
//----------------------------------------------------------------------------

func TestApp01msQueryRow(t *testing.T) {
    var err         error

	t.Logf("TestQueryRow()...\n")
	//DockerRun(t)

	ioApp01ms = NewIoApp01ms()
	ioApp01ms.DefaultParms()

    err = ioApp01ms.DatabaseCreate("App01ms")
    if err != nil {
        t.Errorf("\tError - Database Create failed: %s\n", err.Error())
    }

    err = ioApp01ms.Disconnect()
    if err != nil {
        t.Fatalf("Error: %s\n\n", err)
    }
    ioApp01ms = nil

	t.Logf("TestQueryRow() - End of Test\n\n\n")
}
