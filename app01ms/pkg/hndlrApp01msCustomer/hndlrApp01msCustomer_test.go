// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ms contains all the functions
// and data to interact with the SQL Database.

// Generated: Mon Jan  6, 2020 11:09

package hndlrApp01msCustomer

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "strings"
	"testing"

    "github.com/2kranki/go_util"
	"app01ms/pkg/App01msCustomer"
	"app01ms/pkg/hndlrApp01ms"
	"app01ms/pkg/ioApp01ms"
	"app01ms/pkg/ioApp01msCustomer"
)

//============================================================================
//                          App01msCustomerTestData
//============================================================================

type App01msCustomerTestData struct {
    T           *testing.T
    Port        string
    PW          string
    Server      string
    User        string
    NameDB      string
    io          *ioApp01ms.IO_App01ms
}

//----------------------------------------------------------------------------
//                            Check Status Code
//----------------------------------------------------------------------------

// CheckRcd compares the given record to the needed one and issues an error if
// they do not match.
func (td *App01msCustomerTestData) CheckRcd(need int, rcd *App01msCustomer.App01msCustomer) {
    var rcd2        App01msCustomer.App01msCustomer

    rcd2.TestData(need)

    if rcd.Compare(&rcd2) != 0 {
        td.T.Fatalf("Error: Record Mismatch: needed:%+v have:%+v\n", rcd2, rcd)
    }

}

//----------------------------------------------------------------------------
//                             Disconnect
//----------------------------------------------------------------------------

// Disconnect disconnects the ioApp01ms server.
func (td *App01msCustomerTestData) Disconnect() {
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
func (td *App01msCustomerTestData) Setup(t *testing.T) {

    td.T = t
    td.SetupDB()

}

//----------------------------------------------------------------------------
//                             Set up DB
//----------------------------------------------------------------------------

// SetupDB initializes the DB with test records.
// If it fails at something, it must issue a t.Fatalf().
func (td *App01msCustomerTestData) SetupDB( ) {
    var err         error

    // Set connection parameters based on database SQL type.
    td.io = ioApp01ms.NewIoApp01ms()
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
func NewTestApp01msCustomer() *App01msCustomerTestData {
    td := App01msCustomerTestData{}
    return &td
}

//----------------------------------------------------------------------------
//                          TestData_App01msCustomer
//----------------------------------------------------------------------------

type TestData_App01msCustomer struct {
    T           *testing.T
    bt          *App01msCustomerTestData
    db          *ioApp01msCustomer.IO_App01msCustomer
    H           *HandlersApp01msCustomer
    Mux         *http.ServeMux
    w           *httptest.ResponseRecorder
    Req         *http.Request
    Resp        *http.Response
    tmpls       *hndlrApp01ms.TmplsApp01ms
}

//----------------------------------------------------------------------------
//                            Check Status Code
//----------------------------------------------------------------------------

// CheckStatus checks the request status code for a specific status.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) CheckStatus(status int) {

    
        td.T.Logf("Customer.CheckStatus()\n")
    
    if td.Resp == nil {
        td.T.Fatalf("Error: Missing HTTP Response\n")
    }

    if td.Resp.StatusCode != status {
        td.T.Fatalf("Error: Invalid Status Code of %d, needed %d\n", td.Resp.StatusCode, status)
    }

    
        td.T.Logf("...end Customer.Setup\n")
    
}

//----------------------------------------------------------------------------
//                              GET Request
//----------------------------------------------------------------------------

// GetReq initializes the http.Request for a GET.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) GetReq(target string, body string) {

    
        td.T.Logf("Customer.Setup()\n")
    
    if target == "" {
        td.T.Fatalf("Error: Missing Target String\n")
    }

    td.Req = httptest.NewRequest(http.MethodGet, target, strings.NewReader(body))
    td.ServeHttp()          // Perform the test through the mux.

    
        td.T.Logf("...end Customer.Setup\n")
    
}

//----------------------------------------------------------------------------
//                            POST Request
//----------------------------------------------------------------------------

// SetupPostReq initializes the http.Request for a POST.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) PostReq(target string, body string) {

    
        td.T.Logf("Customer.Setup()\n")
    

    td.Req = httptest.NewRequest(http.MethodPost, target, strings.NewReader(body))
    td.Req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    td.ServeHttp()          // Perform the test through the mux.

    
        td.T.Logf("...end Customer.Setup\n")
    
}

//----------------------------------------------------------------------------
//                            Response Body
//----------------------------------------------------------------------------

// ResponseBody returns the response body converted to a string.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) ResponseBody() string {
    var str     string

    
        td.T.Logf("Customer.ResponseBody()\n")
    
    if td.Resp == nil {
        td.T.Fatalf("Error: Missing HTTP Response\n")
    }

    body, err := ioutil.ReadAll(td.Resp.Body)
    if err != nil {
        td.T.Fatal(err)
    }
    str = string(body)
    td.T.Logf("\tResponse Body: %s\n", body)

    
        td.T.Logf("...end Customer.ResponseBody\n")
    
    return str
}

//----------------------------------------------------------------------------
//                             Serve HTTP
//----------------------------------------------------------------------------

// ServeHttp executes the handler through the mux.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) ServeHttp( ) {

    
        td.T.Logf("Customer.ServeHttp()\n")
    

    td.w = httptest.NewRecorder()
    td.Mux.ServeHTTP(td.w, td.Req)
    td.Resp = td.w.Result()

    
        td.T.Logf("...end Customer.ServeHttp\n")
    
}

//----------------------------------------------------------------------------
//                             Set up
//----------------------------------------------------------------------------

// Setup initializes the Test Data.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) Setup(t *testing.T) {

    td.T = t
    td.SetupIO()
    td.SetupHandlers()

}

//----------------------------------------------------------------------------
//                             Set up I/O
//----------------------------------------------------------------------------

// SetupFakeDB initializes the DB with 2 records.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) SetupIO( ) {
    var err         error
    var rcd         App01msCustomer.App01msCustomer

    td.bt = NewTestApp01msCustomer()
    if td.bt == nil {
        td.T.Fatalf("Error: Unable to allocate ioApp01ms Test!\n")
    } else {
        td.bt.Setup(td.T)
    }

    td.db = ioApp01msCustomer.NewIoApp01msCustomer(td.bt.io)
    if td.db == nil {
        td.T.Fatalf("Error: Unable to allocate FakeDB!\n")
    }

    err = td.db.TableDelete()
    if err != nil {
        td.T.Fatalf("Error: Table Deletion Failure: %s\n\n\n", err.Error())
    }

    err = td.db.TableCreate()
    if err != nil {
        td.T.Fatalf("Error: Cannot create table: %s\n\n\n", err)
    }

    for i:=0; i<2; i++ {
        rcd.TestData(i)
        err = td.db.RowInsert(&rcd)
        if err != nil {
            td.T.Fatalf("Error: Insert %d Failed: %s \n", i, util.ErrorString(err))
        }
    }

}

//----------------------------------------------------------------------------
//                             Set up Handlers
//----------------------------------------------------------------------------

// SetupHandlers initializes HTTP Test Handlers.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01msCustomer) SetupHandlers( ) {

	// Set up main Handler which parses the templates.
    td.tmpls = hndlrApp01ms.NewTmplsApp01ms("../../tmpl")
    td.tmpls.SetupTmpls()

    // Set up the Handler object.
    td.H = &HandlersApp01msCustomer{db:td.db, rowsPerPage:2}
    if td.H == nil {
        td.T.Fatalf("Error: Unable to allocate Handlers\n")
    }

    // Now set up the Server mux for the test.
    td.Mux = http.NewServeMux()
    if td.Mux == nil {
        td.T.Fatalf("Error: Unable to allocate HTTP mux\n")
    }
    td.H.SetupHandlers(td.Mux)

}

//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                          DB
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrDB(t *testing.T) {
    var err         error
    var td          *TestData_App01msCustomer
    var rcd         App01msCustomer.App01msCustomer
    var rcd2        App01msCustomer.App01msCustomer

    t.Logf("TestCustomer.DB()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    t.Logf("\tChecking First()...\n")
    if err = td.db.RowFirst(&rcd2); err != nil {
        t.Fatalf("Error - Read First failed: %s\n", err.Error())
    }
    rcd.TestData(0)
    if 0 != rcd.CompareKeys(&rcd2) {
        t.Fatalf("Error - First did not work, need A, got %+v\n", rcd2)
    }

    t.Logf("TestCustomer.DB() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             List Index
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrListIndex(t *testing.T) {
    var err         error
    var td          *TestData_App01msCustomer
    //var r           string

    t.Logf("TestCustomer.HndlrListIndex()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    if err != nil {
        t.Fatalf("Error: Cannot connect: %s\n", err.Error())
    }

    // Issue a request for ???.
    //TODO: Create a first() request followed by next()'s'.

    // Check response.
    /*TODO: Uncomment when requests are actually being performed.
    r = td.ResponseBody()
    if r != "" {
        t.Logf("\t%s\n", r)
    }
    */

    // Parse response to verify
    //TODO: Parse the response.

    t.Logf("TestCustomer.HndlrListIndex() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             List Show
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrListShow(t *testing.T) {
    var td          *TestData_App01msCustomer

    t.Logf("TestListShow()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    // First try a blank record.
    //TODO: Perform Show()

    // Get the response.
    //TODO: get the response with initial error checking.

    // Parse response to verify
    //TODO: Parse the response.

    t.Logf("TestListShow() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowDelete(t *testing.T) {
    var err         error
    var td          *TestData_App01msCustomer
    var rcd         App01msCustomer.App01msCustomer
    //expectedBody    := ""

    t.Logf("TestRowDelete()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    // Delete a record.
    rcd.TestData(1)             // "B"
    keys := rcd.KeysToValue()
    t.Logf("\tSetting up to delete (%d)\"%s\" row...\n", len(keys), keys)
    urlStr := fmt.Sprintf("/Customer/delete?%s", keys)
    td.GetReq(urlStr, "")

    // Now get the Response and check it.
    td.CheckStatus(http.StatusOK)
    t.Logf("\t actualHeader: %q\n", td.Resp.Header)
    actualBody := td.ResponseBody()
    t.Logf("\t actualBody: %s\n", string(actualBody))
    //TODO: Update this (right now, output is too much.)
    //if expectedBody != string(actualBody) {
        //t.Errorf("Expected the message '%s'\n", expectedBody)
    //}

    rcd.TestData(1)             // "B"
    err = td.db.RowFind(&rcd)
    if err == nil {
        t.Fatalf("Expected Not Found error from RowFind, got ok\n")
    }

    t.Logf("TestRowDelete() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Empty
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowEmpty(t *testing.T) {
    var td          *TestData_App01msCustomer
/*****
    expectedBody    := ""
 *****/

    t.Logf("TestCustomer.RowEmpty()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    // Display empty record.
    t.Logf("\tSetting up for Empty...\n")
    urlStr := fmt.Sprintf("/Customer/empty")
    td.GetReq(urlStr, "")

    // Now get the Response and check it.
    td.CheckStatus(http.StatusOK)
    t.Logf("\t actualHeader: %q\n", td.Resp.Header)
    actualBody := td.ResponseBody()
    t.Logf("\t actualBody: %s\n", string(actualBody))
/*****
    if expectedBody != string(actualBody) {
        t.Errorf("Expected the message '%s'\n", expectedBody)
    }
 *****/

    t.Logf("TestCustomer.RowEmpty() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowFirst(t *testing.T) {
    var td          *TestData_App01msCustomer
    var rcd         App01msCustomer.App01msCustomer
/*****
    expectedBody    := ""
 *****/

    t.Logf("TestCustomer.RowFirst()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    // Delete a record.
    rcd.TestData(2)             // "C"
    keys := rcd.KeysToValue()
    t.Logf("\tSetting up to find first (%d)\"%s\" row...\n", len(keys), keys)
    urlStr := fmt.Sprintf("/Customer/first?%s", keys)
    td.GetReq(urlStr, "")

    // Now get the Response and check it.
    td.CheckStatus(http.StatusOK)
    t.Logf("\t actualHeader: %q\n", td.Resp.Header)
    actualBody := td.ResponseBody()
    t.Logf("\t actualBody: %s\n", string(actualBody))
/*****
    if expectedBody != string(actualBody) {
        t.Errorf("Expected the message '%s'\n", expectedBody)
    }
 *****/

    t.Logf("TestCustomer.RowFirst() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowInsert(t *testing.T) {
    var td          *TestData_App01msCustomer
    var rcd         App01msCustomer.App01msCustomer
    //expectedBody    := ""

    t.Logf("TestCustomerRowInsert()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    // Insert a "Z" record.
    rcd.TestData(25)        // "Z"
    keys := rcd.KeysToValue()
    data := rcd.FieldsToValue()
    urlStr := fmt.Sprintf("/Customer/insert?%s", keys)
    t.Logf("\tSetting up to insert (%d)\"%s\" row...\n", len(keys), keys)
    td.PostReq(urlStr, data)

    // Now get the Response and check it.
    td.CheckStatus(http.StatusOK)
    t.Logf("\t actualHeader: %q\n", td.Resp.Header)
    actualBody := td.ResponseBody()
    t.Logf("\t actualBody: %s\n", string(actualBody))
    //TODO: Update this (right now, output is too much.)
    //if expectedBody != string(actualBody) {
        //t.Errorf("Expected the message '%s'\n", expectedBody)
    //}

    t.Logf("TestCustomerRowInsert() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowNext(t *testing.T) {
    var td          *TestData_App01msCustomer
    var rcd         App01msCustomer.App01msCustomer

    t.Logf("TestCustomer.RowNext()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    // Build and execute a URL.
    rcd.TestData(0)             // "A"
    keys := rcd.KeysToValue()
    t.Logf("\tSetting up for next with keys of (%d)\"%s\"\n", len(keys), keys)
    urlStr := fmt.Sprintf("/Customer/next?%s", keys)
    td.GetReq(urlStr, "")

    t.Logf("TestCustomer.RowNext() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowLastPrev(t *testing.T) {
    var td          *TestData_App01msCustomer

    t.Logf("TestCustomer.RowPrev()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    t.Logf("TestCustomer.RowPrev() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Display
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowDisplay(t *testing.T) {
    var td          *TestData_App01msCustomer

    t.Logf("TestCustomer.RowDisplay()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    t.Logf("TestCustomerRowShow() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

func TestApp01msCustomerHndlrRowUpdate(t *testing.T) {
    var td          *TestData_App01msCustomer

    t.Logf("TestCustomer.RowUpdate()...\n")
    td = &TestData_App01msCustomer{}
    td.Setup(t)

    t.Logf("TestCustomer.RowUpdate() - End of Test\n\n\n")
}

