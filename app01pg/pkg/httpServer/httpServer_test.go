// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Functions to test the httpServer package.

// Generated: Mon Jan  6, 2020 11:09

package httpServer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	//"github.com/2kranki/go_util"
)

//----------------------------------------------------------------------------
//                         Test Support Functions
//----------------------------------------------------------------------------

//----------------------------------------------------------------------------
//                          TestData_App01sqCustomer
//----------------------------------------------------------------------------

type TestData struct {
    T           *testing.T
    //FIXME: H           *Handlers
    Mux         *http.ServeMux
    w           *httptest.ResponseRecorder
    Req         *http.Request
    Resp        *http.Response
    //FIXME: tmpls       *hndlrApp01sq.TmplsApp01sq
}

//----------------------------------------------------------------------------
//                            Check Status Code
//----------------------------------------------------------------------------

// CheckStatus checks the request status code for a specific status.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData) CheckStatus(status int) {


    td.T.Logf("TestData.CheckStatus()\n")

    if td.Resp == nil {
        td.T.Fatalf("Error: Missing HTTP Response\n")
    }

    if td.Resp.StatusCode != status {
        td.T.Fatalf("Error: Invalid Status Code of %d, needed %d\n", td.Resp.StatusCode, status)
    }

    td.T.Logf("...end TestData.CheckStatus\n")
}

//----------------------------------------------------------------------------
//                              GET Request
//----------------------------------------------------------------------------

// GetReq initializes the http.Request for a GET.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData) GetReq(target string, body string) {


        td.T.Logf("TestData.GetReq()\n")

    if target == "" {
        td.T.Fatalf("Error: Missing Target String\n")
    }

    td.Req = httptest.NewRequest(http.MethodGet, target, strings.NewReader(body))
    td.ServeHttp()          // Perform the test through the mux.


        td.T.Logf("...end TestData.GetReq\n")

}

//----------------------------------------------------------------------------
//                            POST Request
//----------------------------------------------------------------------------

// SetupPostReq initializes the http.Request for a POST.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData) PostReq(target string, body string) {


        td.T.Logf("TestData.PostReq()\n")


    td.Req = httptest.NewRequest(http.MethodPost, target, strings.NewReader(body))
    td.Req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    td.ServeHttp()          // Perform the test through the mux.


        td.T.Logf("...end TestData.PostReq\n")

}

//----------------------------------------------------------------------------
//                            Response Body
//----------------------------------------------------------------------------

// ResponseBody returns the response body converted to a string.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData) ResponseBody() string {
    var str     string


        td.T.Logf("TestData.ResponseBody()\n")

    if td.Resp == nil {
        td.T.Fatalf("Error: Missing HTTP Response\n")
    }

    body, err := ioutil.ReadAll(td.Resp.Body)
    if err != nil {
        td.T.Fatal(err)
    }
    str = string(body)
    td.T.Logf("\tResponse Body: %s\n", body)


        td.T.Logf("...end TestData.ResponseBody\n")

    return str
}

//----------------------------------------------------------------------------
//                             Serve HTTP
//----------------------------------------------------------------------------

// ServeHttp executes the handler through the mux.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData) ServeHttp( ) {


        td.T.Logf("TestData.ServeHttp()\n")


    td.w = httptest.NewRecorder()
    td.Mux.ServeHTTP(td.w, td.Req)
    td.Resp = td.w.Result()


        td.T.Logf("...end TestData.ServeHttp\n")

}

//----------------------------------------------------------------------------
//                             Set up
//----------------------------------------------------------------------------

// Setup initializes the Test Data.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData) Setup(t *testing.T) {

    td.T = t
    td.SetupHandlers()

}

//----------------------------------------------------------------------------
//                             Set up Handlers
//----------------------------------------------------------------------------

// SetupHandlers initializes HTTP Test Handlers.
// If it fails at something, it must issue a t.Fatalf().
func (td *TestData) SetupHandlers( ) {

	// Set up main Handler which parses the templates.
    //FIXME: td.tmpls = hndlrApp01sq.NewTmplsApp01sq("../../tmpl")
    //FIXME: td.tmpls.SetupTmpls()

    // Now set up the Server mux for the test.
    td.Mux = http.NewServeMux()
    if td.Mux == nil {
        td.T.Fatalf("Error: Unable to allocate HTTP mux\n")
    }
    //FIXME: td.H.SetupHandlers(td.Mux)

}

// HndlrHome responds to a URL with no sub-elements.  It defaults to
// providing the default Menu to the browser/caller.
func HndlrHome(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "%s", r.URL.String())

}

//============================================================================
//                              Tests
//============================================================================

//----------------------------------------------------------------------------
//                         Test HTTP/HTTPS Server
//----------------------------------------------------------------------------

func TestServer01(t *testing.T) {
    var err         error
    var td          *TestData

	t.Logf("TestServer01()...\n")
    td = &TestData{}
    td.Setup(t)

	//h.Mux.HandleFunc("/", HndlrHome)

	//h.Serve(true)

	t.Logf("TestServer01() - End\n\n\n")
}

//----------------------------------------------------------------------------
//                             ???
//----------------------------------------------------------------------------

func TestSomething(t *testing.T) {
    var td          *TestData
/*****
    expectedBody    := ""
 *****/

    t.Logf("Test.Something()...\n")
    td = &TestData{}
    td.Setup(t)

    t.Logf("\tSetting up for Something...\n")
    urlStr := fmt.Sprintf("/something")
    td.GetReq(urlStr, "")

    // Now get the Response and check it.
    //FIXME: td.CheckStatus(http.StatusOK)
    td.CheckStatus(http.StatusNotFound)
    t.Logf("\t actualHeader: %q\n", td.Resp.Header)
    actualBody := td.ResponseBody()
    t.Logf("\t actualBody: %s\n", string(actualBody))
/*****
    if expectedBody != string(actualBody) {
        t.Errorf("Expected the message '%s'\n", expectedBody)
    }
 *****/

    t.Logf("Test.Something() - End of Test\n\n\n")
}

