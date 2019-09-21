// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01my contains all the functions
// and data to interact with the SQL Database.

// Generated: Sat Sep 21, 2019 06:41

package main

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

type TestData_App01myVendor struct {
	T    *testing.T
	bt   *App01myVendorTestData
	db   *IO_App01myVendor
	H    *HandlersApp01myVendor
	Mux  *http.ServeMux
	w    *httptest.ResponseRecorder
	Req  *http.Request
	Resp *http.Response
}

//----------------------------------------------------------------------------
//                            Check Status Code
//----------------------------------------------------------------------------

// CheckStatus checks the request status code for a specific status.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01myVendor) CheckStatus(status int) {

	td.T.Logf("Vendor.CheckStatus()\n")

	if td.Resp == nil {
		td.T.Fatalf("Error: Missing HTTP Response\n")
	}

	if td.Resp.StatusCode != status {
		td.T.Fatalf("Error: Invalid Status Code of %d, needed %d\n", td.Resp.StatusCode, status)
	}

	td.T.Logf("...end Vendor.Setup\n")

}

//----------------------------------------------------------------------------
//                              GET Request
//----------------------------------------------------------------------------

// GetReq initializes the http.Request for a GET.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01myVendor) GetReq(target string, body string) {

	td.T.Logf("Vendor.Setup()\n")

	if target == "" {
		td.T.Fatalf("Error: Missing Target String\n")
	}

	td.Req = httptest.NewRequest(http.MethodGet, target, strings.NewReader(body))
	td.ServeHttp() // Perform the test through the mux.

	td.T.Logf("...end Vendor.Setup\n")

}

//----------------------------------------------------------------------------
//                            POST Request
//----------------------------------------------------------------------------

// SetupPostReq initializes the http.Request for a POST.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01myVendor) PostReq(target string, body string) {

	td.T.Logf("Vendor.Setup()\n")

	td.Req = httptest.NewRequest(http.MethodPost, target, strings.NewReader(body))
	td.Req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	td.ServeHttp() // Perform the test through the mux.

	td.T.Logf("...end Vendor.Setup\n")

}

//----------------------------------------------------------------------------
//                            Response Body
//----------------------------------------------------------------------------

// ResponseBody returns the response body converted to a string.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01myVendor) ResponseBody() string {
	var str string

	td.T.Logf("Vendor.ResponseBody()\n")

	if td.Resp == nil {
		td.T.Fatalf("Error: Missing HTTP Response\n")
	}

	body, err := ioutil.ReadAll(td.Resp.Body)
	if err != nil {
		td.T.Fatal(err)
	}
	str = string(body)
	td.T.Logf("\tResponse Body: %s\n", body)

	td.T.Logf("...end Vendor.ResponseBody\n")

	return str
}

//----------------------------------------------------------------------------
//                             Serve HTTP
//----------------------------------------------------------------------------

// ServeHttp executes the handler through the mux.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01myVendor) ServeHttp() {

	td.T.Logf("Vendor.ServeHttp()\n")

	td.w = httptest.NewRecorder()
	td.Mux.ServeHTTP(td.w, td.Req)
	td.Resp = td.w.Result()

	td.T.Logf("...end Vendor.ServeHttp\n")

}

//----------------------------------------------------------------------------
//                             Set up
//----------------------------------------------------------------------------

// Setup initializes the Test Data.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01myVendor) Setup(t *testing.T) {

	td.T = t
	td.SetupIO()
	td.SetupHandlers()

}

//----------------------------------------------------------------------------
//                             Set up I/O
//----------------------------------------------------------------------------

// SetupFakeDB initializes the DB with 2 records.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData_App01myVendor) SetupIO() {
	var err error
	var rcd App01myVendor

	td.bt = NewTestApp01myVendor()
	if td.bt == nil {
		td.T.Fatalf("Error: Unable to allocate ioApp01my Test!\n")
	} else {
		td.bt.Setup(td.T)
	}

	td.db = NewIoApp01myVendor(td.bt.io)
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

	for i := 0; i < 2; i++ {
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
func (td *TestData_App01myVendor) SetupHandlers() {

	// Set up main Handler which parses the templates.
	hndlrsApp01my = NewTmplsApp01my()
	hndlrsApp01my.SetTmplsDir("../tmpl")
	hndlrsApp01my.SetupTmpls()

	// Set up the Handler object.
	td.H = &HandlersApp01myVendor{db: td.db, rowsPerPage: 2}
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

func TestApp01myVendorHndlrDB(t *testing.T) {
	var err error
	var td *TestData_App01myVendor
	var rcd App01myVendor
	var rcd2 App01myVendor

	t.Logf("TestVendor.DB()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	t.Logf("\tChecking First()...\n")
	if err = td.db.RowFirst(&rcd2); err != nil {
		t.Fatalf("Error - Read First failed: %s\n", err.Error())
	}
	rcd.TestData(0)
	if 0 != rcd.CompareKeys(&rcd2) {
		t.Fatalf("Error - First did not work, need A, got %+v\n", rcd2)
	}

	t.Logf("TestVendor.DB() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             List Index
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrListIndex(t *testing.T) {
	var err error
	var td *TestData_App01myVendor

	t.Logf("TestVendor.HndlrListIndex()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	//TODO: err = db.Ping()
	if err != nil {
		t.Fatalf("Error: Cannot connect: %s\n", err.Error())
	}

	t.Logf("TestVendor.HndlrListIndex() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             List Show
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrListShow(t *testing.T) {
	var td *TestData_App01myVendor

	t.Logf("TestListShow()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	// First try a blank record.

	t.Logf("TestListShow() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowDelete(t *testing.T) {
	var err error
	var td *TestData_App01myVendor
	var rcd App01myVendor
	//expectedBody    := ""

	t.Logf("TestRowDelete()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	// Delete a record.
	rcd.TestData(1) // "B"
	keys := rcd.KeysToValue()
	t.Logf("\tSetting up to delete (%d)\"%s\" row...\n", len(keys), keys)
	urlStr := fmt.Sprintf("/Vendor/delete?%s", keys)
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

	rcd.TestData(1) // "B"
	err = td.db.RowFind(&rcd)
	if err == nil {
		t.Fatalf("Expected Not Found error from RowFind, got ok\n")
	}

	t.Logf("TestRowDelete() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Empty
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowEmpty(t *testing.T) {
	var td *TestData_App01myVendor
	/*****
	   expectedBody    := ""
	*****/

	t.Logf("TestVendor.RowEmpty()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	// Display empty record.
	t.Logf("\tSetting up for Empty...\n")
	urlStr := fmt.Sprintf("/Vendor/empty")
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

	t.Logf("TestVendor.RowEmpty() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowFirst(t *testing.T) {
	var td *TestData_App01myVendor
	var rcd App01myVendor
	/*****
	   expectedBody    := ""
	*****/

	t.Logf("TestVendor.RowFirst()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	// Delete a record.
	rcd.TestData(2) // "C"
	keys := rcd.KeysToValue()
	t.Logf("\tSetting up to find first (%d)\"%s\" row...\n", len(keys), keys)
	urlStr := fmt.Sprintf("/Vendor/first?%s", keys)
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

	t.Logf("TestVendor.RowFirst() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowInsert(t *testing.T) {
	var td *TestData_App01myVendor
	var rcd App01myVendor
	//expectedBody    := ""

	t.Logf("TestVendorRowInsert()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	// Insert a "Z" record.
	rcd.TestData(25) // "Z"
	keys := rcd.KeysToValue()
	data := rcd.FieldsToValue()
	urlStr := fmt.Sprintf("/Vendor/insert?%s", keys)
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

	t.Logf("TestVendorRowInsert() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowNext(t *testing.T) {
	var td *TestData_App01myVendor
	var rcd App01myVendor

	t.Logf("TestVendor.RowNext()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	// Build and execute a URL.
	rcd.TestData(0) // "A"
	keys := rcd.KeysToValue()
	t.Logf("\tSetting up for next with keys of (%d)\"%s\"\n", len(keys), keys)
	urlStr := fmt.Sprintf("/Vendor/next?%s", keys)
	td.GetReq(urlStr, "")

	t.Logf("TestVendor.RowNext() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowLastPrev(t *testing.T) {
	var td *TestData_App01myVendor

	t.Logf("TestVendor.RowPrev()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	t.Logf("TestVendor.RowPrev() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Display
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowDisplay(t *testing.T) {
	var td *TestData_App01myVendor

	t.Logf("TestVendor.RowDisplay()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	t.Logf("TestVendorRowShow() - End of Test\n\n\n")
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

func TestApp01myVendorHndlrRowUpdate(t *testing.T) {
	var td *TestData_App01myVendor

	t.Logf("TestVendor.RowUpdate()...\n")
	td = &TestData_App01myVendor{}
	td.Setup(t)

	t.Logf("TestVendor.RowUpdate() - End of Test\n\n\n")
}
