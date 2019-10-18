// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01sq contains all the functions
// and data to interact with the SQL Database.

// Generated: Fri Oct 18, 2019 14:51

package hndlrApp01sqCustomer

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "strings"
	"testing"

    "github.com/2kranki/go_util"
)

//============================================================================
//                              Test Data
//============================================================================

type TestData_App01sqCustomer struct {
    T           *testing.T
    bt          *App01sqCustomerTestData
    db          *IO_App01sqCustomer
    H           *HandlersApp01sqCustomer
    Mux         *http.ServeMux
    w           *httptest.ResponseRecorder
    Req         *http.Request
    Resp        *http.Response
}

//----------------------------------------------------------------------------
//                            Check Status Code
//----------------------------------------------------------------------------

// CheckStatus checks the request status code for a specific status.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01sqCustomer) CheckStatus(status int) {

    
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
func (td *TestData_App01sqCustomer) GetReq(target string, body string) {

    
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
func (td *TestData_App01sqCustomer) PostReq(target string, body string) {

    
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
func (td *TestData_App01sqCustomer) ResponseBody() string {
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
func (td *TestData_App01sqCustomer) ServeHttp( ) {

    
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
func (td *TestData_App01sqCustomer) Setup(t *testing.T) {

    td.T = t
    td.SetupIO()
    td.SetupHandlers()

}

//----------------------------------------------------------------------------
//                             Set up I/O
//----------------------------------------------------------------------------

// SetupFakeDB initializes the DB with 2 records.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01sqCustomer) SetupIO( ) {
    var err         error
    var rcd         App01sqCustomer

    td.bt = NewTestApp01sqCustomer()
    if td.bt == nil {
        td.T.Fatalf("Error: Unable to allocate ioApp01sq Test!\n")
    } else {
        td.bt.Setup(td.T)
    }

    td.db = NewIoApp01sqCustomer(td.bt.io)
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
func (td *TestData_App01sqCustomer) SetupHandlers( ) {

	// Set up main Handler which parses the templates.
    hndlrsApp01sq = NewTmplsApp01sq()
    hndlrsApp01sq.SetTmplsDir("../tmpl")
    hndlrsApp01sq.SetupTmpls()

    // Set up the Handler object.
    td.H = &HandlersApp01sqCustomer{db:td.db, rowsPerPage:2}
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

func TestApp01sqCustomerHndlrDB(t *testing.T) {
    var err         error
    var td          *TestData_App01sqCustomer
    var rcd         App01sqCustomer
    var rcd2        App01sqCustomer

    t.Logf("TestCustomer.DB()...\n")
    td = &TestData_App01sqCustomer{}
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

func TestApp01sqCustomerHndlrListIndex(t *testing.T) {
    var err         error
    var td          *TestData_App01sqCustomer
    var r           string

    t.Logf("TestCustomer.HndlrListIndex()...\n")
    td = &TestData_App01sqCustomer{}
    td.Setup(t)

    if err != nil {
        t.Fatalf("Error: Cannot connect: %s\n", err.Error())
    }

    r = td.ResponseBody()

    t.Logf("TestCustomer.HndlrListIndex() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             List Show
//----------------------------------------------------------------------------

func TestApp01sqCustomerHndlrListShow(t *testing.T) {
    var td          *TestData_App01sqCustomer

    t.Logf("TestListShow()...\n")
    td = &TestData_App01sqCustomer{}
    td.Setup(t)

    // First try a blank record.

    t.Logf("TestListShow() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

func TestApp01sqCustomerHndlrRowDelete(t *testing.T) {
    var err         error
    var td          *TestData_App01sqCustomer
    var rcd         App01sqCustomer
    //expectedBody    := ""

    t.Logf("TestRowDelete()...\n")
    td = &TestData_App01sqCustomer{}
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

func TestApp01sqCustomerHndlrRowEmpty(t *testing.T) {
    var td          *TestData_App01sqCustomer
/*****
    expectedBody    := ""
 *****/

    t.Logf("TestCustomer.RowEmpty()...\n")
    td = &TestData_App01sqCustomer{}
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

func TestApp01sqCustomerHndlrRowFirst(t *testing.T) {
    var td          *TestData_App01sqCustomer
    var rcd         App01sqCustomer
/*****
    expectedBody    := ""
 *****/

    t.Logf("TestCustomer.RowFirst()...\n")
    td = &TestData_App01sqCustomer{}
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

func TestApp01sqCustomerHndlrRowInsert(t *testing.T) {
    var td          *TestData_App01sqCustomer
    var rcd         App01sqCustomer
    //expectedBody    := ""

    t.Logf("TestCustomerRowInsert()...\n")
    td = &TestData_App01sqCustomer{}
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

func TestApp01sqCustomerHndlrRowNext(t *testing.T) {
    var td          *TestData_App01sqCustomer
    var rcd         App01sqCustomer

    t.Logf("TestCustomer.RowNext()...\n")
    td = &TestData_App01sqCustomer{}
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

func TestApp01sqCustomerHndlrRowLastPrev(t *testing.T) {
    var td          *TestData_App01sqCustomer

    t.Logf("TestCustomer.RowPrev()...\n")
    td = &TestData_App01sqCustomer{}
    td.Setup(t)

    t.Logf("TestCustomer.RowPrev() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Display
//----------------------------------------------------------------------------

func TestApp01sqCustomerHndlrRowDisplay(t *testing.T) {
    var td          *TestData_App01sqCustomer

    t.Logf("TestCustomer.RowDisplay()...\n")
    td = &TestData_App01sqCustomer{}
    td.Setup(t)

    t.Logf("TestCustomerRowShow() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

func TestApp01sqCustomerHndlrRowUpdate(t *testing.T) {
    var td          *TestData_App01sqCustomer

    t.Logf("TestCustomer.RowUpdate()...\n")
    td = &TestData_App01sqCustomer{}
    td.Setup(t)

    t.Logf("TestCustomer.RowUpdate() - End of Test\n\n\n")
}

