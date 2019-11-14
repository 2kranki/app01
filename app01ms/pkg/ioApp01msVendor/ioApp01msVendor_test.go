// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioVendor_test tests various functions of
// the Table SQL Maintenance methods.

// Generated: Thu Nov 14, 2019 11:17 for mssql Database

package ioApp01msVendor

import (
	"testing"
    

	_ "github.com/denisenkom/go-mssqldb"
)


//============================================================================
//                              Test Data
//============================================================================

type App01msVendorTestData struct {
    T           *testing.T
    Port        string
    PW          string
    Server      string
    User        string
    NameDB      string
    io          *IO_App01ms
}

//----------------------------------------------------------------------------
//                            Check Status Code
//----------------------------------------------------------------------------

// CheckRcd compares the given record to the needed one and issues an error if
// they do not match.
func (td *App01msVendorTestData) CheckRcd(need int, rcd *App01msVendor) {
    var rcd2        App01msVendor

    rcd2.TestData(need)

    if rcd.Compare(&rcd2) != 0 {
        td.T.Fatalf("Error: Record Mismatch: needed:%+v have:%+v\n", rcd2, rcd)
    }

}

//----------------------------------------------------------------------------
//                             Disconnect
//----------------------------------------------------------------------------

// Disconnect disconnects the ioApp01ms server.
func (td *App01msVendorTestData) Disconnect() {
    var err         error

    err = td.io.Disconnect()
    if err != nil {
        td.T.Fatalf("Error: Disconnect Failure: %s\n", err.Error())
    }

}

//----------------------------------------------------------------------------
//                             Set up
//----------------------------------------------------------------------------

// Setup initializes the Test Data.
// If it fails at something, it must issue a t.Fatalf().
func (td *App01msVendorTestData) Setup(t *testing.T) {

    td.T = t
    td.SetupDB()

}

//----------------------------------------------------------------------------
//                             Set up DB
//----------------------------------------------------------------------------

// SetupDB initializes the DB with test records.
// If it fails at something, it must issue a t.Fatalf().
func (td *App01msVendorTestData) SetupDB( ) {
    var err         error

    // Set connection parameters based on database SQL type.
    td.io = NewIoApp01ms()
    td.io.DefaultParms()
    err = td.io.DatabaseCreate("App01ms")
    if err != nil {
        td.T.Fatalf("Error: Creation Failure: %s\n", err.Error())
    }

}

//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new io struct.
func NewTestApp01msVendor() *App01msVendorTestData {
    td := App01msVendorTestData{}
    return &td
}

//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                              Create Table
//----------------------------------------------------------------------------

func TestApp01msVendorCreateDeleteTable(t *testing.T) {
    var err         error
    var td          *App01msVendorTestData
    var io          *IO_App01msVendor

	t.Logf("TestCreateTable()...\n")
	DockerRun(t)
	td = NewTestApp01msVendor()
	td.Setup(t)
	io = NewIoApp01msVendor(td.io)

    err = io.TableDelete()
    if err != nil {
        t.Fatalf("Error: Table Deletion Failure: %s\n\n\n", err.Error())
    }

    err = io.TableCreate()
    if err != nil {
        t.Fatalf("Error: Cannot create table: %s\n\n\n", err)
    }

    err = io.TableDelete()
    if err != nil {
        t.Fatalf("Error: Table Deletion Failure: %s\n\n\n", err.Error())
    }

    td.Disconnect()
	t.Logf("TestCreateTable() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Row Insert
//----------------------------------------------------------------------------

func TestApp01msVendorRowInsert(t *testing.T) {
    var err         error
    var td          *App01msVendorTestData
    var io          *IO_App01msVendor
    var rcd         App01msVendor

    t.Logf("TestVendor.RowInsert()...\n")
	DockerRun(t)
	td = NewTestApp01msVendor()
	td.Setup(t)
	io = NewIoApp01msVendor(td.io)

    // Start clean with new empty tables.
    err = io.TableCreate()
    if err != nil {
        t.Fatal("Error: Cannot create tables: ", err)
    }

    // Now add some records.
    for i := 0; i < 5; i++ {
        t.Logf("\tInserting Record %d\n", i)
        rcd.TestData(i)
        err = io.RowInsert(&rcd)
        if err != nil {
            t.Fatalf("Error: : Record Insertion Failed: %s\n\n\n", err)
        }
    }

    // Now read the first record.
    t.Logf("\tReading First Record\n")
    err = io.RowFirst(&rcd)
    if err != nil {
        t.Fatalf("Error: : Record First Failed: %s\n\n\n", err)
    }
    td.CheckRcd(0, &rcd)

    // Now read the last record.
    t.Logf("\tReading Last Record\n")
    err = io.RowLast(&rcd)
    if err != nil {
        t.Fatalf("Error: : Record Last Failed: %s\n\n\n", err)
    }
    td.CheckRcd(4, &rcd)

    // Now read the middle record.
    t.Logf("\tReading via Find the Middle Record\n")
    rcd.TestData(2)
    err = io.RowFind(&rcd)
    if err != nil {
        t.Fatalf("Error: : Record Middle Failed: %s\n\n\n", err)
    }
    td.CheckRcd(2, &rcd)

    // Now read the first record via Find.
    t.Logf("\tReading via Find the First Record\n")
    rcd.TestData(0)
    err = io.RowFind(&rcd)
    if err != nil {
        t.Fatalf("Error: : Record First Failed: %s\n\n\n", err)
    }
    td.CheckRcd(0, &rcd)

    // Now read the last record via Find.
    t.Logf("\tReading via Find the Last Record\n")
    rcd.TestData(2)
    err = io.RowFind(&rcd)
    if err != nil {
        t.Fatalf("Error: : Record Last Failed: %s\n\n\n", err)
    }
    td.CheckRcd(2, &rcd)

    err = io.TableDelete()
    if err != nil {
        t.Fatal("Error: Cannot delete tables: ", err)
    }

    td.Disconnect()
    t.Logf("TestVendor RowInsert() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Row Page
//----------------------------------------------------------------------------

func TestApp01msVendorRowPage(t *testing.T) {
    var err         error
    var td          *App01msVendorTestData
    var io          *IO_App01msVendor
    var rcd         App01msVendor
    var rcds        []App01msVendor

    t.Logf("TestVendorRowPage()...\n")
	DockerRun(t)
	td = NewTestApp01msVendor()
	td.Setup(t)
	io = NewIoApp01msVendor(td.io)

    // Start clean with new empty tables.
    err = io.TableCreate()
    if err != nil {
        t.Fatal("Error: Cannot create tables: ", err)
    }

    // Now add some records.
    for i := 0; i < 10; i++ {
        chr := 'A' + i
        t.Logf("\tInserting Row %d - %c\n", i, chr)
        rcd.TestData(i)
        err = io.RowInsert(&rcd)
        if err != nil {
            t.Fatalf("Error: : Row Insertion Failed: %s\n\n\n", err)
        }
    }

    t.Logf("\tReading First Set of 4 Records\n")
    rcds, err = io.RowPage(0, 4)
    if err != nil {
        t.Fatalf("Error: First Record Set Failed: %s\n\n\n", err)
    }
    t.Logf("1 rcds(%d): %+v\n",len(rcds),rcds)
    if len(rcds) != 4 {
        t.Fatalf("Error: Number of Record Verification Failed\n\n\n")
    }
    for i:=0; i<4; i++ {
        td.CheckRcd(i, &rcds[i])
    }

    t.Logf("\tReading Second set of 4 Records\n")
    rcds, err = io.RowPage(4, 4)
    if err != nil {
        t.Fatalf("Error: : First Record Set Failed: %s\n\n\n", err)
    }
    t.Logf("2 rcds(%d): %+v\n",len(rcds),rcds)
    if len(rcds) != 4 {
        t.Fatalf("Error: : Number of Record Verification Failed\n\n\n")
    }
    for i:=0; i<4; i++ {
        td.CheckRcd(i+4, &rcds[i])
    }

    t.Logf("\tReading Third set of Records\n")
    rcds, err = io.RowPage(8, 4)
    if err != nil {
        t.Fatalf("Error: : First Record Set Failed: %s\n\n\n", err)
    }
    t.Logf("3 rcds(%d): %+v\n",len(rcds),rcds)
    if len(rcds) != 2 {
        t.Fatalf("Error: : Number of Record Verification Failed\n\n\n")
    }
    for i:=0; i<2; i++ {
        td.CheckRcd(i+8, &rcds[i])
    }

    // Now read the Fourth set of records. (That don't exist!)
    t.Logf("\tReading Fourth set of Records\n")
    rcds, err = io.RowPage(13, 4)
    if err != nil {
        t.Fatalf("Error: : Fourth Record Set Failed: %s\n\n\n", err)
    }
    if len(rcds) != 0 {
        t.Fatalf("Error: : Number of Record Verification Failed\n\n\n")
    }

    err = io.TableDelete()
    if err != nil {
        t.Fatal("Error: Cannot delete tables: ", err)
    }

    td.Disconnect()
    t.Logf("TestVendorRowInsert() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                              Table Scanner
//----------------------------------------------------------------------------

func TestApp01msVendorTableScan(t *testing.T) {
    var err         error
    var td          *App01msVendorTestData
    var io          *IO_App01msVendor
    var rcd         App01msVendor
    var cnt         int

	t.Logf("TestTableScan()...\n")
	DockerRun(t)
	td = NewTestApp01msVendor()
	td.Setup(t)
	io = NewIoApp01msVendor(td.io)

    // Start clean with new empty tables.
    err = io.TableCreate()
    if err != nil {
        t.Fatal("Error: Cannot create tables: ", err)
    }

    // Now add some records.
    for i := 0; i < 10; i++ {
        chr := 'A' + i
        t.Logf("\tInserting Row %d - %c\n", i, chr)
        rcd.TestData(i)
        err = io.RowInsert(&rcd)
        if err != nil {
            t.Fatalf("Error: : Row Insertion Failed: %s\n\n\n", err)
        }
    }

    apply := func (rcd App01msVendor) error {
                t.Logf("\tScan Row %d\n", cnt)
                cnt++
                return nil
             }
    err = io.TableScan(apply)
    if err != nil {
        t.Fatal("Error: Scanner: ", err)
    }
    if cnt != 10 {
        t.Fatalf("Error: Scanner Count: %d - should be 10", cnt)
    }


    td.Disconnect()
	t.Logf("TestCreateTable() - End of Test\n\n\n")
}

