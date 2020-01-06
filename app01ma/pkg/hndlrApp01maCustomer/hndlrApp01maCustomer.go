// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Mon Jan  6, 2020 09:54

package hndlrApp01maCustomer

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"strings"

	"sync"

	"app01ma/pkg/App01maCustomer"
	"app01ma/pkg/hndlrApp01ma"
	"app01ma/pkg/ioApp01maCustomer"
	"github.com/2kranki/go_util"
	_ "github.com/go-sql-driver/mysql"
)

//============================================================================
//                              Miscellaneous
//============================================================================

//============================================================================
//                        Handlers for App01ma.Customer
//============================================================================

type HandlersApp01maCustomer struct {
	mu          sync.Mutex
	db          *ioApp01maCustomer.IO_App01maCustomer
	rowsPerPage int
	Tmpls       *hndlrApp01ma.TmplsApp01ma
}

//----------------------------------------------------------------------------
//                             Accessors
//----------------------------------------------------------------------------

func (h *HandlersApp01maCustomer) DB() *ioApp01maCustomer.IO_App01maCustomer {
	return h.db
}

func (h *HandlersApp01maCustomer) SetDB(db *ioApp01maCustomer.IO_App01maCustomer) {
	h.db = db
}

func (h *HandlersApp01maCustomer) RowsPerPage() int {
	return h.rowsPerPage
}

func (h *HandlersApp01maCustomer) SetRowsPerPage(r int) {
	h.rowsPerPage = r
}

//----------------------------------------------------------------------------
//                           Setup Handlers
//----------------------------------------------------------------------------

// SetupHandlers creates a Handler object and sets up each of the handlers
// with it given a mux.
func (h *HandlersApp01maCustomer) SetupHandlers(mux *http.ServeMux) {

	log.Printf("\thndlrCustomer.SetupHandlers()\n")

	mux.HandleFunc("/Customer/list/first", h.ListFirst)
	mux.HandleFunc("/Customer", h.ListFirst)
	mux.HandleFunc("/Customer/list/last", h.ListLast)
	mux.HandleFunc("/Customer/list/next", h.ListNext)
	mux.HandleFunc("/Customer/list/prev", h.ListPrev)
	mux.HandleFunc("/Customer/delete", h.RowDelete)
	mux.HandleFunc("/Customer/empty", h.RowEmpty)
	mux.HandleFunc("/Customer/find", h.RowFind)
	mux.HandleFunc("/Customer/first", h.RowFirst)
	mux.HandleFunc("/Customer/form", h.RowForm)
	mux.HandleFunc("/Customer/insert", h.RowInsert)
	mux.HandleFunc("/Customer/last", h.RowLast)
	mux.HandleFunc("/Customer/next", h.RowNext)
	mux.HandleFunc("/Customer/prev", h.RowPrev)
	mux.HandleFunc("/Customer/show", h.RowShow)
	mux.HandleFunc("/Customer/update", h.RowUpdate)
	mux.HandleFunc("/Customer/table/create", h.TableCreate)
	mux.HandleFunc("/Customer/table/load/csv", h.TableLoadCSV)
	mux.HandleFunc("/Customer/table/load/test", h.TableLoadTestData)
	mux.HandleFunc("/Customer/table/save/csv", h.TableSaveCSV)

	log.Printf("\tend of hndlrCustomer.SetupHandlers()\n")

}

//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new Handlers object given the parameters needed by the handlers
// and returns it to the caller if successful.  If it fails to properly create
// the handlers then it must fail rather than return an error indicator.
func NewHandlersApp01maCustomer(db *ioApp01maCustomer.IO_App01maCustomer, rowsPerPage int, mux *http.ServeMux) *HandlersApp01maCustomer {
	var h *HandlersApp01maCustomer

	h = &HandlersApp01maCustomer{db: db, rowsPerPage: rowsPerPage}
	if h == nil {
		log.Fatalf("Error: Unable to allocate Handlers for hndlrApp01maCustomer!\n")
	}
	h.SetupHandlers(mux)

	return h
}

//============================================================================
//                              List Form Handlers
//============================================================================

//----------------------------------------------------------------------------
//                             List First
//----------------------------------------------------------------------------

// ListFirst displays the first page of rows.
func (h *HandlersApp01maCustomer) ListFirst(w http.ResponseWriter, r *http.Request) {

	log.Printf("hndlrCustomer.ListFirst(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.ListFirst(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Display the row in the form.
	h.ListShow(w, 0, "")

	log.Printf("...end hndlrCustomer.ListFirst()\n")

}

//----------------------------------------------------------------------------
//                             List Last
//----------------------------------------------------------------------------

// ListLast displays the last page of rows.
func (h *HandlersApp01maCustomer) ListLast(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset int

	log.Printf("hndlrCustomer.ListLast(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.ListLast(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Calculate the offset.
	offset, err = h.db.TableCount()
	if err != nil {

		log.Printf("...end hndlrCustomer.ListLast(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	offset -= h.rowsPerPage
	if offset < 0 {
		offset = 0
	}

	// Display the row in the form.
	h.ListShow(w, offset, "")

	log.Printf("...end hndlrCustomer.ListLast()\n")

}

//----------------------------------------------------------------------------
//                             List Next
//----------------------------------------------------------------------------

// ListNext displays the next page of rows.
func (h *HandlersApp01maCustomer) ListNext(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset int
	var cTable int

	log.Printf("hndlrCustomer.ListNext(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.ListNext(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Calculate the offset.
	cTable, err = h.db.TableCount()
	if err != nil {

		log.Printf("...end hndlrCustomer.ListLast(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	offset, _ = strconv.Atoi(r.FormValue("offset"))
	offset += h.rowsPerPage
	if offset < 0 || offset > cTable {
		offset = 0
	}

	// Display the row in the form.
	h.ListShow(w, offset, "")

	log.Printf("...end hndlrCustomer.ListNext()\n")

}

//----------------------------------------------------------------------------
//                             List Prev
//----------------------------------------------------------------------------

// ListPrev displays the next page of rows.
func (h *HandlersApp01maCustomer) ListPrev(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset int
	var begin int
	var cTable int

	log.Printf("hndlrCustomer.ListPrev(%s, %s)\n", r.Method, r.FormValue("offset"))

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.ListPrev(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Calculate the offset.
	cTable, err = h.db.TableCount()
	if err != nil {

		log.Printf("...end hndlrCustomer.ListLast(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	begin, _ = strconv.Atoi(r.FormValue("offset"))
	offset = begin - h.rowsPerPage
	if offset < 0 {
		if begin > 0 {
			offset = 0
		} else {
			offset = cTable - h.rowsPerPage
			if offset < 0 {
				offset = 0
			}
		}
	}

	// Display the row in the form.
	h.ListShow(w, offset, "")

	log.Printf("...end hndlrCustomer.ListPrev()\n")

}

//----------------------------------------------------------------------------
//                             List Show
//----------------------------------------------------------------------------

// ListShow displays a list page given a starting offset.
func (h *HandlersApp01maCustomer) ListShow(w http.ResponseWriter, offset int, msg string) {
	var err error
	var rcds []App01maCustomer.App01maCustomer
	var name = "App01ma.Customer.list.gohtml"

	var str strings.Builder

	log.Printf("hndlrCustomer.ListShow(%d)\n", offset)
	log.Printf("\tname: %s\n", name)
	w2 := io.MultiWriter(w, &str)

	// Get the records to display
	rcds, err = h.db.RowPage(offset, h.rowsPerPage)
	if err != nil {

		log.Printf("...end hndlrCustomer.ListShow(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	data := struct {
		Rcds   []App01maCustomer.App01maCustomer
		Offset int
		Msg    string
	}{rcds, offset, msg}

	log.Printf("\tData: %+v\n", data)

	log.Printf("\tExecuting template: %s\n", name)
	err = h.Tmpls.Tmpls.ExecuteTemplate(w2, name, data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	log.Printf("\t output: %s\n", str.String())
	log.Printf("...end hndlrCustomer.ListShow(%s)\n", util.ErrorString(err))

}

//============================================================================
//                             Row Form Handlers
//============================================================================

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete handles an delete request which comes from the row display form.
func (h *HandlersApp01maCustomer) RowDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maCustomer.App01maCustomer
	var i int
	var key string

	log.Printf("hndlrCustomer.RowDelete(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.RowDelete(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Num, _ = strconv.ParseInt(key, 0, 64)

	i++

	log.Printf("\t rcd: %+v\n", rcd)

	// Delete the row with data given.
	err = h.db.RowDelete(&rcd)
	if err != nil {

		log.Printf("...end hndlrCustomer.RowDelete(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Get the next row in the form with status message and display it.
	err = h.db.RowNext(&rcd)
	if err != nil {

		log.Printf("...end hndlrCustomer.RowDelete(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "Row deleted!")

	log.Printf("...end hndlrCustomer.RowDelete(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                                Row Display
//----------------------------------------------------------------------------

// RowDisplay displays the given record.
func (h *HandlersApp01maCustomer) RowDisplay(w http.ResponseWriter, rcd *App01maCustomer.App01maCustomer, msg string) {
	var err error

	log.Printf("hndlrCustomer.RowDisplay(%+v, %s)\n", rcd, msg)

	if h.Tmpls != nil {
		data := struct {
			Rcd *App01maCustomer.App01maCustomer
			Msg string
		}{rcd, msg}
		name := "App01ma.Customer.form.gohtml"

		log.Printf("\tRcd: %+v\n", data.Rcd)
		log.Printf("\tMsg: %s\n", data.Msg)
		log.Printf("\tname: %s\n", name)

		err = h.Tmpls.Tmpls.ExecuteTemplate(w, name, data)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	}

	log.Printf("...end hndlrCustomer.RowDisplay(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                             Row Empty
//----------------------------------------------------------------------------

// RowEmpty displays the table row form with an empty row.
func (h *HandlersApp01maCustomer) RowEmpty(w http.ResponseWriter, r *http.Request) {
	var rcd App01maCustomer.App01maCustomer

	log.Printf("hndlrCustomer.RowEmpty(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.RowEmpty(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the row to display and display it.
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrCustomer.RowEmpty()\n")

}

//----------------------------------------------------------------------------
//                             Row Find
//----------------------------------------------------------------------------

// RowFind handles displaying of the table row form display.
func (h *HandlersApp01maCustomer) RowFind(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maCustomer.App01maCustomer
	var msg string
	var i int
	var key string

	log.Printf("hndlrCustomer.RowFind(%s, %s)\n", r.Method, r.FormValue("key"))

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.RowFind(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Num, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Get the row and display it.
	err = h.db.RowFind(&rcd)
	if err != nil {
		msg = "Row NOT Found!"
		err = h.db.RowFirst(&rcd)
	}
	if err != nil {

		log.Printf("...end hndlrCustomer.RowFind(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, msg)

	log.Printf("...end hndlrCustomer.RowFind()\n")

}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

// RowFirst displays the first row.
func (h *HandlersApp01maCustomer) RowFirst(w http.ResponseWriter, r *http.Request) {
	var rcd App01maCustomer.App01maCustomer
	var err error

	log.Printf("hndlrCustomer.RowFirst(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.RowFirst(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the next row and display it.
	err = h.db.RowFirst(&rcd)
	if err != nil {

		log.Printf("...end hndlrCustomer.RowFirst(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrCustomer.RowFirst()\n")

}

//----------------------------------------------------------------------------
//                             Row Form
//----------------------------------------------------------------------------

// RowForm displays the raw table row form without data.
func (h *HandlersApp01maCustomer) RowForm(w http.ResponseWriter, r *http.Request) {

	log.Printf("hndlrCustomer.RowForm(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.RowForm(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Verify any fields that need it.

	// Get the row to display.

	// Display the row in the form.
	http.ServeFile(w, r, "./tmpl/App01ma.Customer.form.gohtml")

	log.Printf("...end hndlrCustomer.RowForm()\n")

}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

// RowInsert handles an add row request which comes from the row display form.
func (h *HandlersApp01maCustomer) RowInsert(w http.ResponseWriter, r *http.Request) {
	var rcd App01maCustomer.App01maCustomer
	var err error

	log.Printf("hndlrCustomer.RowInsert(%s)\n", r.Method)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Create a record from the data given.
	err = rcd.Request2Struct(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Verify any fields that need it.

	// Add the row.
	err = h.db.RowInsert(&rcd)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Get the last row as a guess of where the inserted row went and display it.
	_ = h.db.RowLast(&rcd)
	h.RowDisplay(w, &rcd, "Row added!")

	log.Printf("...end hndlrCustomer.RowInsert(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                             Row Last
//----------------------------------------------------------------------------

// RowLast displays the first row.
func (h *HandlersApp01maCustomer) RowLast(w http.ResponseWriter, r *http.Request) {
	var rcd App01maCustomer.App01maCustomer
	var err error

	log.Printf("hndlrCustomer.RowLast(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomer.RowLast(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the next row to display.
	err = h.db.RowLast(&rcd)
	if err != nil {

		log.Printf("...end hndlrCustomer.RowLast(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Display the row in the form.
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrCustomer.RowLast()\n")

}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

// RowNext handles an next request which comes from the row display form and
// should display the next row from the current one.
func (h *HandlersApp01maCustomer) RowNext(w http.ResponseWriter, r *http.Request) {
	var rcd App01maCustomer.App01maCustomer
	var err error
	var i int
	var key string

	log.Printf("hndlrCustomer.RowNext(%s)\n", r.Method)
	log.Printf("\tURL: %q\n", r.URL)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the prior key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Num, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Get the next row and display it.
	err = h.db.RowNext(&rcd)
	if err != nil {

		log.Printf("...end hndlrCustomer.RowNext(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrCustomer.RowNext()\n")

}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

// RowPrev handles an previous request which comes from the row display form
// and should display the previous row from the current one.
func (h *HandlersApp01maCustomer) RowPrev(w http.ResponseWriter, r *http.Request) {
	var rcd App01maCustomer.App01maCustomer
	var err error
	var i int
	var key string

	log.Printf("hndlrCustomer.RowPrev(%s)\n", r.Method)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the prior key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Num, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Get the next row and display it.
	err = h.db.RowPrev(&rcd)
	if err != nil {

		log.Printf("...end Customer.RowNext(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrCustomer.RowPrev()\n")

}

//----------------------------------------------------------------------------
//                             Row Show
//----------------------------------------------------------------------------

// RowShow handles displaying of the table row form display.
func (h *HandlersApp01maCustomer) RowShow(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	var rcd App01maCustomer.App01maCustomer

	log.Printf("hndlrCustomer.RowShow(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrCustomerHndlrShow(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Verify any fields that need it.
	//TODO: key = r.FormValue("[.Table.PrimaryKey.Name]")
	//TODO: if key is not present, assume first record.

	//TODO: log.Printf("\tkey: %s\n", key)

	// Get the row to display.
	if key == "" {
		err = h.db.RowFirst(&rcd)
	} else {
		err = h.db.RowFind(&rcd)
	}
	if err != nil {

		log.Printf("...end hndlrCustomer.RowShow(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Display the row in the form.
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrCustomer.RowShow()\n")

}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

// RowUpdate handles an update request which comes from the row display form.
func (h *HandlersApp01maCustomer) RowUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	var rcd App01maCustomer.App01maCustomer
	var i int

	log.Printf("hndlrCustomer.RowUpdate(%s)\n", r.Method)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the prior key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Num, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Delete the row.
	err = h.db.RowDelete(&rcd)
	if err != nil {

		log.Printf("...end hndlrCustomer.RowNext(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Create a record from the data given.
	err = rcd.Request2Struct(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Add the row.
	err = h.db.RowInsert(&rcd)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Display the next row in the form.
	h.RowDisplay(w, &rcd, "Record updated")

	log.Printf("...end hndlrCustomer.RowUpdate()\n")

}

//============================================================================
//                             Table Form Handlers
//============================================================================

//----------------------------------------------------------------------------
//                             Table Create
//----------------------------------------------------------------------------

// TableCreate creates the table deleting any current ones.
func (h *HandlersApp01maCustomer) TableCreate(w http.ResponseWriter, r *http.Request) {
	var err error

	log.Printf("hndlrCustomer.TableCreate(%s)\n", r.Method)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Create the table.
	err = h.db.TableCreate()
	if err == nil {
		//h.ListShow(w, 0, "Table was created")
		w.Write([]byte("Table was created"))
	} else {
		w.Write([]byte("Table creation had an error of:" + err.Error()))
	}

	log.Printf("...end hndlrCustomer.TableCreate(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                            Table Load CSV
//----------------------------------------------------------------------------

// TableLoadCSV creates the table deleting any current ones and loads in
// data from a CSV file.
func (h *HandlersApp01maCustomer) TableLoadCSV(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maCustomer.App01maCustomer
	var fileIn multipart.File
	var cnt int
	var maxMem int64
	var handler *multipart.FileHeader

	log.Printf("hndlrCustomer.TableLoadCSV(%s)\n", r.Method)
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// ParseMultipartForm parses a request body as multipart/form-data.
	// The whole request body is parsed and up to a total of maxMemory
	// bytes of its file parts are stored in memory, with the remainder
	// stored on disk in temporary files. ParseMultipartForm calls ParseForm
	// if necessary. After one call to ParseMultipartForm, subsequent
	// calls have no effect.
	name := "csvFile" // Must match Name parameter of Form's "<input type=file name=???>"
	maxMem = 64 << 20 // 64mb
	r.ParseMultipartForm(maxMem)

	// FormFile returns the first file for the given key which was
	// specified on the Form Input Type=file Name parameter.
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	fileIn, handler, err = r.FormFile(name)
	if err != nil {
		log.Printf("...end hndlrCustomer.TableLoadCSV(Error:500) - %s\n", util.ErrorString(err))
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer fileIn.Close() //close the file when we finish
	log.Printf("\tUploaded File: %+v\n", handler.Filename)
	log.Printf("\tFile Size: %+v\n", handler.Size)
	log.Printf("\tMIME Header: %+v\n", handler.Header)
	rdr := csv.NewReader(fileIn)

	// Create the table.
	err = h.db.TableCreate()
	if err != nil {
		w.Write([]byte("Table creation had an error of:" + util.ErrorString(err)))
	}

	log.Printf("\tLoading data...\n")
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			str := fmt.Sprintf("ERROR: Reading row %d from csv - %s\n", cnt, util.ErrorString(err))
			w.Write([]byte(str))
			return
		}

		rcd.Num, _ = strconv.ParseInt(record[0], 0, 64)

		rcd.Name = record[1]

		rcd.Addr1 = record[2]

		rcd.Addr2 = record[3]

		rcd.City = record[4]

		rcd.State = record[5]

		rcd.Zip = record[6]

		rcd.Curbal, _ = strconv.ParseFloat(record[7], 64)

		err = h.db.RowInsert(&rcd)
		if err != nil {
			str := fmt.Sprintf("ERROR: Table creation had an error of: %s\n", util.ErrorString(err))
			w.Write([]byte(str))
			return
		}
		cnt++
		log.Printf("\t...Added row %d\n", cnt)
	}
	for i := 1; i > 0; i-- {
		str := fmt.Sprintf("Added %d rows\n", cnt)
		w.Write([]byte(str))
	}

	log.Printf("...end hndlrCustomer.TableLoadCSV(ok) - %d\n", cnt)

}

//----------------------------------------------------------------------------
//                             Table Load Test Data
//----------------------------------------------------------------------------

// TableLoadTestData creates the table deleting any current ones and loads
// in some test rows.
func (h *HandlersApp01maCustomer) TableLoadTestData(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maCustomer.App01maCustomer

	log.Printf("hndlrCustomer.TableLoadTestData(%s)\n", r.Method)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Create the table.
	err = h.db.TableCreate()
	if err == nil {
		w.Write([]byte("Table was created\n"))
	} else {
		w.Write([]byte("Table creation had an error of:" + util.ErrorString(err)))
	}

	// Load the test rows.
	// Now add some records.
	for i := 0; i < 26; i++ {
		chr := 'A' + i
		rcd.TestData(i)
		err = h.db.RowInsert(&rcd)
		if err == nil {
			str := fmt.Sprintf("Added row: %c\n", chr)
			w.Write([]byte(str))
		} else {
			str := fmt.Sprintf("Table creation had an error of: %c\n", chr)
			w.Write([]byte(str))
		}
	}

	log.Printf("...end hndlrCustomer.TableLoadTestData(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                            Table Save CSV
//----------------------------------------------------------------------------

// TableSaveCSV creates the table deleting any current ones and loads in
// data from a CSV file.
func (h *HandlersApp01maCustomer) TableSaveCSV(w http.ResponseWriter, r *http.Request) {
	var err error
	var cntGood int
	var cntTotal int

	log.Printf("hndlrCustomer.TableSaveCSV(%s)\n", r.Method)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Set up to write the CSV file.
	fileName := "Customer.csv"
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	wtr := csv.NewWriter(w)

	// Write the CSV file.
	if wtr != nil {
		apply := func(rcd App01maCustomer.App01maCustomer) error {
			log.Printf("\tRow: %v\n", rcd.ToStrings())
			err2 := wtr.Write(rcd.ToStrings())
			cntTotal++
			if err2 == nil {
				cntGood++
			}
			return err2
		}
		err = h.db.TableScan(apply)
		wtr.Flush()
	} else {
		err = fmt.Errorf("Error: Could not create CSV Writer\n")
		log.Printf("\t%s - for App01maCustomer table!\n", util.ErrorString(err))
	}

	log.Printf("...end hndlrCustomer.TableSaveCSV(%s)\n", util.ErrorString(err))
}
