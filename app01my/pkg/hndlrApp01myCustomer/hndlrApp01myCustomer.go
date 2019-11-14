// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Thu Nov 14, 2019 11:17


package hndlrApp01myCustomer

import (
	"encoding/csv"
    "fmt"
    "io"
    "log"
    "mime/multipart"
	"net/http"
	"strconv"
    
	"sync"
    

	"github.com/2kranki/go_util"
	_ "github.com/go-sql-driver/mysql"
	"app01my/pkg/App01myCustomer"
	"app01my/pkg/hndlrApp01my"
	"app01my/pkg/ioApp01myCustomer"
)



//============================================================================
//                              Miscellaneous
//============================================================================

//============================================================================
//                        Handlers for App01my.Customer
//============================================================================

type HandlersApp01myCustomer struct {
    mu          sync.Mutex
    db          *ioApp01myCustomer.IO_App01myCustomer
    rowsPerPage int
    Tmpls       *hndlrApp01my.TmplsApp01my
}

//----------------------------------------------------------------------------
//                             Accessors
//----------------------------------------------------------------------------

func (h *HandlersApp01myCustomer) DB() *ioApp01myCustomer.IO_App01myCustomer {
    return h.db
}

func (h *HandlersApp01myCustomer) SetDB(db *ioApp01myCustomer.IO_App01myCustomer) {
    h.db = db
}

func (h *HandlersApp01myCustomer) RowsPerPage() int {
    return h.rowsPerPage
}

func (h *HandlersApp01myCustomer) SetRowsPerPage(r int) {
    h.rowsPerPage = r
}

//----------------------------------------------------------------------------
//                           Setup Handlers
//----------------------------------------------------------------------------

// SetupHandlers creates a Handler object and sets up each of the handlers
// with it given a mux.
func (h *HandlersApp01myCustomer) SetupHandlers(mux *http.ServeMux) {

    

	    mux.HandleFunc("/Customer/list/first",         h.ListFirst)
	    mux.HandleFunc("/Customer",                    h.ListFirst)
	    mux.HandleFunc("/Customer/list/last",          h.ListLast)
	    mux.HandleFunc("/Customer/list/next",          h.ListNext)
	    mux.HandleFunc("/Customer/list/prev",          h.ListPrev)
	    mux.HandleFunc("/Customer/delete",             h.RowDelete)
	    mux.HandleFunc("/Customer/empty",              h.RowEmpty)
	    mux.HandleFunc("/Customer/find",               h.RowFind)
	    mux.HandleFunc("/Customer/first",              h.RowFirst)
	    mux.HandleFunc("/Customer/form",               h.RowForm)
	    mux.HandleFunc("/Customer/insert",             h.RowInsert)
	    mux.HandleFunc("/Customer/last",               h.RowLast)
	    mux.HandleFunc("/Customer/next",               h.RowNext)
	    mux.HandleFunc("/Customer/prev",               h.RowPrev)
	    mux.HandleFunc("/Customer/show",               h.RowShow)
	    mux.HandleFunc("/Customer/update",             h.RowUpdate)
	    mux.HandleFunc("/Customer/table/create",       h.TableCreate)
	    mux.HandleFunc("/Customer/table/load/csv",     h.TableLoadCSV)
	    mux.HandleFunc("/Customer/table/load/test",    h.TableLoadTestData)
	    mux.HandleFunc("/Customer/table/save/csv",     h.TableSaveCSV)

    
}

//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new Handlers object given the parameters needed by the handlers
// and returns it to the caller if successful.  If it fails to properly create
// the handlers then it must fail rather than return an error indicator.
func NewHandlersApp01myCustomer(db *ioApp01myCustomer.IO_App01myCustomer, rowsPerPage int, mux *http.ServeMux) *HandlersApp01myCustomer {
    var h       *HandlersApp01myCustomer

 	h = &HandlersApp01myCustomer{db:db, rowsPerPage:rowsPerPage}
    if h == nil {
        log.Fatalf("Error: Unable to allocate Handlers for hndlrApp01myCustomer!\n")
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
func (h *HandlersApp01myCustomer) ListFirst(w http.ResponseWriter, r *http.Request) {

    

    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Display the row in the form.
    h.ListShow(w, 0, "")

    
}

//----------------------------------------------------------------------------
//                             List Last
//----------------------------------------------------------------------------

// ListLast displays the last page of rows.
func (h *HandlersApp01myCustomer) ListLast(w http.ResponseWriter, r *http.Request) {
    var err     error
    var offset  int

    

    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Calculate the offset.
    offset, err = h.db.TableCount()
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
    }
    offset -= h.rowsPerPage
    if offset < 0 {
        offset = 0
    }

    // Display the row in the form.
    h.ListShow(w, offset, "")

    
}

//----------------------------------------------------------------------------
//                             List Next
//----------------------------------------------------------------------------

// ListNext displays the next page of rows.
func (h *HandlersApp01myCustomer) ListNext(w http.ResponseWriter, r *http.Request) {
    var err     error
    var offset  int
    var cTable  int

    

    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Calculate the offset.
    cTable, err = h.db.TableCount()
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
    }
    offset, _ = strconv.Atoi(r.FormValue("offset"))
    offset += h.rowsPerPage
    if offset < 0 || offset > cTable {
        offset = 0
    }

    // Display the row in the form.
    h.ListShow(w, offset, "")

    
}

//----------------------------------------------------------------------------
//                             List Prev
//----------------------------------------------------------------------------

// ListPrev displays the next page of rows.
func (h *HandlersApp01myCustomer) ListPrev(w http.ResponseWriter, r *http.Request) {
    var err     error
    var offset  int
    var begin   int
    var cTable  int

    

    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Calculate the offset.
    cTable, err = h.db.TableCount()
    if err != nil {
        
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

    
}

//----------------------------------------------------------------------------
//                             List Show
//----------------------------------------------------------------------------

// ListShow displays a list page given a starting offset.
func (h *HandlersApp01myCustomer) ListShow(w http.ResponseWriter, offset int, msg string) {
    var err     error
    var rcds    []App01myCustomer.App01myCustomer
    var name    = "App01my.Customer.list.gohtml"
    

    

    // Get the records to display
    rcds, err = h.db.RowPage(offset, h.rowsPerPage)
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }

    data := struct {
                Rcds        []App01myCustomer.App01myCustomer
                Offset      int
                Msg         string
            }{rcds, offset, msg}

    

    
        err = h.Tmpls.Tmpls.ExecuteTemplate(w, name, data)
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    
}

//============================================================================
//                             Row Form Handlers
//============================================================================

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete handles an delete request which comes from the row display form.
func (h *HandlersApp01myCustomer) RowDelete(w http.ResponseWriter, r *http.Request) {
    var err     error
    var rcd     App01myCustomer.App01myCustomer
    var i       int
    var key     string

    
    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Get the key(s).
    i = 0
    key = r.FormValue(fmt.Sprintf("key%d", i))
        	rcd.Num, _ = strconv.ParseInt(key,0,64)

        i++
    

    // Delete the row with data given.
    err = h.db.RowDelete(&rcd)
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }

    // Get the next row in the form with status message and display it.
    err = h.db.RowNext(&rcd)
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }
    h.RowDisplay(w, &rcd, "Row deleted!")

    
}

//----------------------------------------------------------------------------
//                                Row Display
//----------------------------------------------------------------------------

// RowDisplay displays the given record.
func (h *HandlersApp01myCustomer) RowDisplay(w http.ResponseWriter, rcd  *App01myCustomer.App01myCustomer, msg string) {
    var err     error

    

    if h.Tmpls != nil {
        data := struct {
                    Rcd         *App01myCustomer.App01myCustomer
                    Msg         string
                }{rcd, msg}
        name := "App01my.Customer.form.gohtml"
        
        err = h.Tmpls.Tmpls.ExecuteTemplate(w, name, data)
        if err != nil {
            fmt.Fprintf(w, err.Error())
        }
    }

    
}

//----------------------------------------------------------------------------
//                             Row Empty
//----------------------------------------------------------------------------

// RowEmpty displays the table row form with an empty row.
func (h *HandlersApp01myCustomer) RowEmpty(w http.ResponseWriter, r *http.Request) {
    var rcd     App01myCustomer.App01myCustomer

    
    if r.Method != "GET" {
    
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Get the row to display and display it.
    h.RowDisplay(w, &rcd, "")

    
}

//----------------------------------------------------------------------------
//                             Row Find
//----------------------------------------------------------------------------

// RowFind handles displaying of the table row form display.
func (h *HandlersApp01myCustomer) RowFind(w http.ResponseWriter, r *http.Request) {
    var err     error
    var rcd     App01myCustomer.App01myCustomer
    var msg     string
    var i       int
    var key     string

    
    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Get the key(s).
    i = 0
    key = r.FormValue(fmt.Sprintf("key%d", i))
        	rcd.Num, _ = strconv.ParseInt(key,0,64)

        i++

    // Get the row and display it.
    err = h.db.RowFind(&rcd)
    if err != nil {
        msg = "Row NOT Found!"
        err = h.db.RowFirst(&rcd)
    }
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }
    h.RowDisplay(w, &rcd, msg)

    
}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

// RowFirst displays the first row.
func (h *HandlersApp01myCustomer) RowFirst(w http.ResponseWriter, r *http.Request) {
    var rcd     App01myCustomer.App01myCustomer
    var err     error

    

    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Get the next row and display it.
    err = h.db.RowFirst(&rcd)
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }
    h.RowDisplay(w, &rcd, "")


    
}

//----------------------------------------------------------------------------
//                             Row Form
//----------------------------------------------------------------------------

// RowForm displays the raw table row form without data.
func (h *HandlersApp01myCustomer) RowForm(w http.ResponseWriter, r *http.Request) {

    
    if r.Method != "GET" {
    
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Verify any fields that need it.

    // Get the row to display.

    // Display the row in the form.
    http.ServeFile(w, r, "./tmpl/App01my.Customer.form.gohtml")

    
}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

// RowInsert handles an add row request which comes from the row display form.
func (h *HandlersApp01myCustomer) RowInsert(w http.ResponseWriter, r *http.Request) {
    var rcd         App01myCustomer.App01myCustomer
    var err         error

    
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

    
}

//----------------------------------------------------------------------------
//                             Row Last
//----------------------------------------------------------------------------

// RowLast displays the first row.
func (h *HandlersApp01myCustomer) RowLast(w http.ResponseWriter, r *http.Request) {
    var rcd         App01myCustomer.App01myCustomer
    var err         error

    
    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Get the next row to display.
    err = h.db.RowLast(&rcd)
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }

    // Display the row in the form.
    h.RowDisplay(w, &rcd, "")

    
}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

// RowNext handles an next request which comes from the row display form and
// should display the next row from the current one.
func (h *HandlersApp01myCustomer) RowNext(w http.ResponseWriter, r *http.Request) {
    var rcd         App01myCustomer.App01myCustomer
    var err         error
    var i           int
    var key         string

    
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Get the prior key(s).
    i = 0
    key = r.FormValue(fmt.Sprintf("key%d", i))
        	rcd.Num, _ = strconv.ParseInt(key,0,64)

        i++

    // Get the next row and display it.
    err = h.db.RowNext(&rcd)
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }
    h.RowDisplay(w, &rcd, "")

    
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

// RowPrev handles an previous request which comes from the row display form
// and should display the previous row from the current one.
func (h *HandlersApp01myCustomer) RowPrev(w http.ResponseWriter, r *http.Request) {
    var rcd         App01myCustomer.App01myCustomer
    var err         error
    var i           int
    var key         string

    
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Get the prior key(s).
    i = 0
    key = r.FormValue(fmt.Sprintf("key%d", i))
        	rcd.Num, _ = strconv.ParseInt(key,0,64)

        i++

    // Get the next row and display it.
    err = h.db.RowPrev(&rcd)
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }
    h.RowDisplay(w, &rcd, "")

    
}

//----------------------------------------------------------------------------
//                             Row Show
//----------------------------------------------------------------------------

// RowShow handles displaying of the table row form display.
func (h *HandlersApp01myCustomer) RowShow(w http.ResponseWriter, r *http.Request) {
    var err         error
    var key         string
    var rcd         App01myCustomer.App01myCustomer

    
    if r.Method != "GET" {
        
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    // Verify any fields that need it.
    //TODO: key = r.FormValue("[.Table.PrimaryKey.Name]")
    //TODO: if key is not present, assume first record.
    

    // Get the row to display.
    if key == "" {
        err = h.db.RowFirst(&rcd)
    } else {
        err = h.db.RowFind(&rcd)
    }
    if err != nil {
        
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }

    // Display the row in the form.
    h.RowDisplay(w, &rcd, "")

    
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

// RowUpdate handles an update request which comes from the row display form.
func (h *HandlersApp01myCustomer) RowUpdate(w http.ResponseWriter, r *http.Request) {
    var err         error
    var key         string
    var rcd         App01myCustomer.App01myCustomer
    var i           int

    
    if r.Method != "POST" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    
    
    

    // Get the prior key(s).
    i = 0
    key = r.FormValue(fmt.Sprintf("key%d", i))
        	rcd.Num, _ = strconv.ParseInt(key,0,64)

        i++

    // Delete the row.
    err = h.db.RowDelete(&rcd)
    if err != nil {
        
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

    
}

//============================================================================
//                             Table Form Handlers
//============================================================================

//----------------------------------------------------------------------------
//                             Table Create
//----------------------------------------------------------------------------

// TableCreate creates the table deleting any current ones.
func (h *HandlersApp01myCustomer) TableCreate(w http.ResponseWriter, r *http.Request) {
    var err         error

    
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

    
}

//----------------------------------------------------------------------------
//                            Table Load CSV
//----------------------------------------------------------------------------

// TableLoadCSV creates the table deleting any current ones and loads in
// data from a CSV file.
func (h *HandlersApp01myCustomer) TableLoadCSV(w http.ResponseWriter, r *http.Request) {
    var err         error
    var rcd         App01myCustomer.App01myCustomer
    var fileIn      multipart.File
    var cnt         int
    var maxMem      int64
    

    
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
    name := "csvFile"           // Must match Name parameter of Form's "<input type=file name=???>"
    maxMem = 64 << 20           // 64mb
    r.ParseMultipartForm(maxMem)

    // FormFile returns the first file for the given key which was
    // specified on the Form Input Type=file Name parameter.
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    
        fileIn, _, err = r.FormFile(name)
    if err != nil {
    
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        return
    }
    defer fileIn.Close() //close the file when we finish
    
    rdr := csv.NewReader(fileIn)

    // Create the table.
    err = h.db.TableCreate()
    if err != nil {
        w.Write([]byte("Table creation had an error of:" + util.ErrorString(err)))
    }

    
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

        
            	rcd.Num, _ = strconv.ParseInt(record[0],0,64)

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
        
    }
    for i := 1; i > 0; i-- {
        str := fmt.Sprintf("Added %d rows\n", cnt)
        w.Write([]byte(str))
    }

    
}

//----------------------------------------------------------------------------
//                             Table Load Test Data
//----------------------------------------------------------------------------

// TableLoadTestData creates the table deleting any current ones and loads
// in some test rows.
func (h *HandlersApp01myCustomer) TableLoadTestData(w http.ResponseWriter, r *http.Request) {
    var err         error
    var rcd         App01myCustomer.App01myCustomer

    
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

    
}

//----------------------------------------------------------------------------
//                            Table Save CSV
//----------------------------------------------------------------------------

// TableSaveCSV creates the table deleting any current ones and loads in
// data from a CSV file.
func (h *HandlersApp01myCustomer) TableSaveCSV(w http.ResponseWriter, r *http.Request) {
    var err         error
    var cntGood     int
    var cntTotal    int

    
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
        apply := func (rcd App01myCustomer.App01myCustomer) error {
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
        log.Printf("\t%s - for App01myCustomer table!\n", util.ErrorString(err))
    }

    
}

