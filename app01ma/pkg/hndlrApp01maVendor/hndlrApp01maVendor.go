// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Mon Jan  6, 2020 09:54

package hndlrApp01maVendor

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

	"app01ma/pkg/App01maVendor"
	"app01ma/pkg/hndlrApp01ma"
	"app01ma/pkg/ioApp01maVendor"
	"github.com/2kranki/go_util"
	_ "github.com/go-sql-driver/mysql"
)

//============================================================================
//                              Miscellaneous
//============================================================================

//============================================================================
//                        Handlers for App01ma.Vendor
//============================================================================

type HandlersApp01maVendor struct {
	mu          sync.Mutex
	db          *ioApp01maVendor.IO_App01maVendor
	rowsPerPage int
	Tmpls       *hndlrApp01ma.TmplsApp01ma
}

//----------------------------------------------------------------------------
//                             Accessors
//----------------------------------------------------------------------------

func (h *HandlersApp01maVendor) DB() *ioApp01maVendor.IO_App01maVendor {
	return h.db
}

func (h *HandlersApp01maVendor) SetDB(db *ioApp01maVendor.IO_App01maVendor) {
	h.db = db
}

func (h *HandlersApp01maVendor) RowsPerPage() int {
	return h.rowsPerPage
}

func (h *HandlersApp01maVendor) SetRowsPerPage(r int) {
	h.rowsPerPage = r
}

//----------------------------------------------------------------------------
//                           Setup Handlers
//----------------------------------------------------------------------------

// SetupHandlers creates a Handler object and sets up each of the handlers
// with it given a mux.
func (h *HandlersApp01maVendor) SetupHandlers(mux *http.ServeMux) {

	log.Printf("\thndlrVendor.SetupHandlers()\n")

	mux.HandleFunc("/Vendor/list/first", h.ListFirst)
	mux.HandleFunc("/Vendor", h.ListFirst)
	mux.HandleFunc("/Vendor/list/last", h.ListLast)
	mux.HandleFunc("/Vendor/list/next", h.ListNext)
	mux.HandleFunc("/Vendor/list/prev", h.ListPrev)
	mux.HandleFunc("/Vendor/delete", h.RowDelete)
	mux.HandleFunc("/Vendor/empty", h.RowEmpty)
	mux.HandleFunc("/Vendor/find", h.RowFind)
	mux.HandleFunc("/Vendor/first", h.RowFirst)
	mux.HandleFunc("/Vendor/form", h.RowForm)
	mux.HandleFunc("/Vendor/insert", h.RowInsert)
	mux.HandleFunc("/Vendor/last", h.RowLast)
	mux.HandleFunc("/Vendor/next", h.RowNext)
	mux.HandleFunc("/Vendor/prev", h.RowPrev)
	mux.HandleFunc("/Vendor/show", h.RowShow)
	mux.HandleFunc("/Vendor/update", h.RowUpdate)
	mux.HandleFunc("/Vendor/table/create", h.TableCreate)
	mux.HandleFunc("/Vendor/table/load/csv", h.TableLoadCSV)
	mux.HandleFunc("/Vendor/table/load/test", h.TableLoadTestData)
	mux.HandleFunc("/Vendor/table/save/csv", h.TableSaveCSV)

	log.Printf("\tend of hndlrVendor.SetupHandlers()\n")

}

//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new Handlers object given the parameters needed by the handlers
// and returns it to the caller if successful.  If it fails to properly create
// the handlers then it must fail rather than return an error indicator.
func NewHandlersApp01maVendor(db *ioApp01maVendor.IO_App01maVendor, rowsPerPage int, mux *http.ServeMux) *HandlersApp01maVendor {
	var h *HandlersApp01maVendor

	h = &HandlersApp01maVendor{db: db, rowsPerPage: rowsPerPage}
	if h == nil {
		log.Fatalf("Error: Unable to allocate Handlers for hndlrApp01maVendor!\n")
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
func (h *HandlersApp01maVendor) ListFirst(w http.ResponseWriter, r *http.Request) {

	log.Printf("hndlrVendor.ListFirst(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.ListFirst(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Display the row in the form.
	h.ListShow(w, 0, "")

	log.Printf("...end hndlrVendor.ListFirst()\n")

}

//----------------------------------------------------------------------------
//                             List Last
//----------------------------------------------------------------------------

// ListLast displays the last page of rows.
func (h *HandlersApp01maVendor) ListLast(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset int

	log.Printf("hndlrVendor.ListLast(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.ListLast(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Calculate the offset.
	offset, err = h.db.TableCount()
	if err != nil {

		log.Printf("...end hndlrVendor.ListLast(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	offset -= h.rowsPerPage
	if offset < 0 {
		offset = 0
	}

	// Display the row in the form.
	h.ListShow(w, offset, "")

	log.Printf("...end hndlrVendor.ListLast()\n")

}

//----------------------------------------------------------------------------
//                             List Next
//----------------------------------------------------------------------------

// ListNext displays the next page of rows.
func (h *HandlersApp01maVendor) ListNext(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset int
	var cTable int

	log.Printf("hndlrVendor.ListNext(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.ListNext(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Calculate the offset.
	cTable, err = h.db.TableCount()
	if err != nil {

		log.Printf("...end hndlrVendor.ListLast(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	offset, _ = strconv.Atoi(r.FormValue("offset"))
	offset += h.rowsPerPage
	if offset < 0 || offset > cTable {
		offset = 0
	}

	// Display the row in the form.
	h.ListShow(w, offset, "")

	log.Printf("...end hndlrVendor.ListNext()\n")

}

//----------------------------------------------------------------------------
//                             List Prev
//----------------------------------------------------------------------------

// ListPrev displays the next page of rows.
func (h *HandlersApp01maVendor) ListPrev(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset int
	var begin int
	var cTable int

	log.Printf("hndlrVendor.ListPrev(%s, %s)\n", r.Method, r.FormValue("offset"))

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.ListPrev(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Calculate the offset.
	cTable, err = h.db.TableCount()
	if err != nil {

		log.Printf("...end hndlrVendor.ListLast(Error:400) - %s\n", util.ErrorString(err))

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

	log.Printf("...end hndlrVendor.ListPrev()\n")

}

//----------------------------------------------------------------------------
//                             List Show
//----------------------------------------------------------------------------

// ListShow displays a list page given a starting offset.
func (h *HandlersApp01maVendor) ListShow(w http.ResponseWriter, offset int, msg string) {
	var err error
	var rcds []App01maVendor.App01maVendor
	var name = "App01ma.Vendor.list.gohtml"

	var str strings.Builder

	log.Printf("hndlrVendor.ListShow(%d)\n", offset)
	log.Printf("\tname: %s\n", name)
	w2 := io.MultiWriter(w, &str)

	// Get the records to display
	rcds, err = h.db.RowPage(offset, h.rowsPerPage)
	if err != nil {

		log.Printf("...end hndlrVendor.ListShow(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	data := struct {
		Rcds   []App01maVendor.App01maVendor
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
	log.Printf("...end hndlrVendor.ListShow(%s)\n", util.ErrorString(err))

}

//============================================================================
//                             Row Form Handlers
//============================================================================

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete handles an delete request which comes from the row display form.
func (h *HandlersApp01maVendor) RowDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maVendor.App01maVendor
	var i int
	var key string

	log.Printf("hndlrVendor.RowDelete(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.RowDelete(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Id, _ = strconv.ParseInt(key, 0, 64)

	i++

	log.Printf("\t rcd: %+v\n", rcd)

	// Delete the row with data given.
	err = h.db.RowDelete(&rcd)
	if err != nil {

		log.Printf("...end hndlrVendor.RowDelete(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Get the next row in the form with status message and display it.
	err = h.db.RowNext(&rcd)
	if err != nil {

		log.Printf("...end hndlrVendor.RowDelete(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "Row deleted!")

	log.Printf("...end hndlrVendor.RowDelete(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                                Row Display
//----------------------------------------------------------------------------

// RowDisplay displays the given record.
func (h *HandlersApp01maVendor) RowDisplay(w http.ResponseWriter, rcd *App01maVendor.App01maVendor, msg string) {
	var err error

	log.Printf("hndlrVendor.RowDisplay(%+v, %s)\n", rcd, msg)

	if h.Tmpls != nil {
		data := struct {
			Rcd *App01maVendor.App01maVendor
			Msg string
		}{rcd, msg}
		name := "App01ma.Vendor.form.gohtml"

		log.Printf("\tRcd: %+v\n", data.Rcd)
		log.Printf("\tMsg: %s\n", data.Msg)
		log.Printf("\tname: %s\n", name)

		err = h.Tmpls.Tmpls.ExecuteTemplate(w, name, data)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	}

	log.Printf("...end hndlrVendor.RowDisplay(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                             Row Empty
//----------------------------------------------------------------------------

// RowEmpty displays the table row form with an empty row.
func (h *HandlersApp01maVendor) RowEmpty(w http.ResponseWriter, r *http.Request) {
	var rcd App01maVendor.App01maVendor

	log.Printf("hndlrVendor.RowEmpty(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.RowEmpty(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the row to display and display it.
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrVendor.RowEmpty()\n")

}

//----------------------------------------------------------------------------
//                             Row Find
//----------------------------------------------------------------------------

// RowFind handles displaying of the table row form display.
func (h *HandlersApp01maVendor) RowFind(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maVendor.App01maVendor
	var msg string
	var i int
	var key string

	log.Printf("hndlrVendor.RowFind(%s, %s)\n", r.Method, r.FormValue("key"))

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.RowFind(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Id, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Get the row and display it.
	err = h.db.RowFind(&rcd)
	if err != nil {
		msg = "Row NOT Found!"
		err = h.db.RowFirst(&rcd)
	}
	if err != nil {

		log.Printf("...end hndlrVendor.RowFind(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, msg)

	log.Printf("...end hndlrVendor.RowFind()\n")

}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

// RowFirst displays the first row.
func (h *HandlersApp01maVendor) RowFirst(w http.ResponseWriter, r *http.Request) {
	var rcd App01maVendor.App01maVendor
	var err error

	log.Printf("hndlrVendor.RowFirst(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.RowFirst(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the next row and display it.
	err = h.db.RowFirst(&rcd)
	if err != nil {

		log.Printf("...end hndlrVendor.RowFirst(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrVendor.RowFirst()\n")

}

//----------------------------------------------------------------------------
//                             Row Form
//----------------------------------------------------------------------------

// RowForm displays the raw table row form without data.
func (h *HandlersApp01maVendor) RowForm(w http.ResponseWriter, r *http.Request) {

	log.Printf("hndlrVendor.RowForm(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.RowForm(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Verify any fields that need it.

	// Get the row to display.

	// Display the row in the form.
	http.ServeFile(w, r, "./tmpl/App01ma.Vendor.form.gohtml")

	log.Printf("...end hndlrVendor.RowForm()\n")

}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

// RowInsert handles an add row request which comes from the row display form.
func (h *HandlersApp01maVendor) RowInsert(w http.ResponseWriter, r *http.Request) {
	var rcd App01maVendor.App01maVendor
	var err error

	log.Printf("hndlrVendor.RowInsert(%s)\n", r.Method)

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

	log.Printf("...end hndlrVendor.RowInsert(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                             Row Last
//----------------------------------------------------------------------------

// RowLast displays the first row.
func (h *HandlersApp01maVendor) RowLast(w http.ResponseWriter, r *http.Request) {
	var rcd App01maVendor.App01maVendor
	var err error

	log.Printf("hndlrVendor.RowLast(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendor.RowLast(Error:405) - Not GET\n")

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the next row to display.
	err = h.db.RowLast(&rcd)
	if err != nil {

		log.Printf("...end hndlrVendor.RowLast(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Display the row in the form.
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrVendor.RowLast()\n")

}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

// RowNext handles an next request which comes from the row display form and
// should display the next row from the current one.
func (h *HandlersApp01maVendor) RowNext(w http.ResponseWriter, r *http.Request) {
	var rcd App01maVendor.App01maVendor
	var err error
	var i int
	var key string

	log.Printf("hndlrVendor.RowNext(%s)\n", r.Method)
	log.Printf("\tURL: %q\n", r.URL)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the prior key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Id, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Get the next row and display it.
	err = h.db.RowNext(&rcd)
	if err != nil {

		log.Printf("...end hndlrVendor.RowNext(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrVendor.RowNext()\n")

}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

// RowPrev handles an previous request which comes from the row display form
// and should display the previous row from the current one.
func (h *HandlersApp01maVendor) RowPrev(w http.ResponseWriter, r *http.Request) {
	var rcd App01maVendor.App01maVendor
	var err error
	var i int
	var key string

	log.Printf("hndlrVendor.RowPrev(%s)\n", r.Method)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the prior key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Id, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Get the next row and display it.
	err = h.db.RowPrev(&rcd)
	if err != nil {

		log.Printf("...end Vendor.RowNext(Error:400) - No Key\n")

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrVendor.RowPrev()\n")

}

//----------------------------------------------------------------------------
//                             Row Show
//----------------------------------------------------------------------------

// RowShow handles displaying of the table row form display.
func (h *HandlersApp01maVendor) RowShow(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	var rcd App01maVendor.App01maVendor

	log.Printf("hndlrVendor.RowShow(%s)\n", r.Method)

	if r.Method != "GET" {

		log.Printf("...end hndlrVendorHndlrShow(Error:405) - Not GET\n")

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

		log.Printf("...end hndlrVendor.RowShow(Error:400) - %s\n", util.ErrorString(err))

		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Display the row in the form.
	h.RowDisplay(w, &rcd, "")

	log.Printf("...end hndlrVendor.RowShow()\n")

}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

// RowUpdate handles an update request which comes from the row display form.
func (h *HandlersApp01maVendor) RowUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	var rcd App01maVendor.App01maVendor
	var i int

	log.Printf("hndlrVendor.RowUpdate(%s)\n", r.Method)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Get the prior key(s).
	i = 0
	key = r.FormValue(fmt.Sprintf("key%d", i))
	rcd.Id, _ = strconv.ParseInt(key, 0, 64)

	i++

	// Delete the row.
	err = h.db.RowDelete(&rcd)
	if err != nil {

		log.Printf("...end hndlrVendor.RowNext(Error:400) - %s\n", util.ErrorString(err))

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

	log.Printf("...end hndlrVendor.RowUpdate()\n")

}

//============================================================================
//                             Table Form Handlers
//============================================================================

//----------------------------------------------------------------------------
//                             Table Create
//----------------------------------------------------------------------------

// TableCreate creates the table deleting any current ones.
func (h *HandlersApp01maVendor) TableCreate(w http.ResponseWriter, r *http.Request) {
	var err error

	log.Printf("hndlrVendor.TableCreate(%s)\n", r.Method)

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

	log.Printf("...end hndlrVendor.TableCreate(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                            Table Load CSV
//----------------------------------------------------------------------------

// TableLoadCSV creates the table deleting any current ones and loads in
// data from a CSV file.
func (h *HandlersApp01maVendor) TableLoadCSV(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maVendor.App01maVendor
	var fileIn multipart.File
	var cnt int
	var maxMem int64
	var handler *multipart.FileHeader

	log.Printf("hndlrVendor.TableLoadCSV(%s)\n", r.Method)
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
		log.Printf("...end hndlrVendor.TableLoadCSV(Error:500) - %s\n", util.ErrorString(err))
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

		rcd.Id, _ = strconv.ParseInt(record[0], 0, 64)

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

	log.Printf("...end hndlrVendor.TableLoadCSV(ok) - %d\n", cnt)

}

//----------------------------------------------------------------------------
//                             Table Load Test Data
//----------------------------------------------------------------------------

// TableLoadTestData creates the table deleting any current ones and loads
// in some test rows.
func (h *HandlersApp01maVendor) TableLoadTestData(w http.ResponseWriter, r *http.Request) {
	var err error
	var rcd App01maVendor.App01maVendor

	log.Printf("hndlrVendor.TableLoadTestData(%s)\n", r.Method)

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

	log.Printf("...end hndlrVendor.TableLoadTestData(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                            Table Save CSV
//----------------------------------------------------------------------------

// TableSaveCSV creates the table deleting any current ones and loads in
// data from a CSV file.
func (h *HandlersApp01maVendor) TableSaveCSV(w http.ResponseWriter, r *http.Request) {
	var err error
	var cntGood int
	var cntTotal int

	log.Printf("hndlrVendor.TableSaveCSV(%s)\n", r.Method)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Set up to write the CSV file.
	fileName := "Vendor.csv"
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	wtr := csv.NewWriter(w)

	// Write the CSV file.
	if wtr != nil {
		apply := func(rcd App01maVendor.App01maVendor) error {
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
		log.Printf("\t%s - for App01maVendor table!\n", util.ErrorString(err))
	}

	log.Printf("...end hndlrVendor.TableSaveCSV(%s)\n", util.ErrorString(err))
}
