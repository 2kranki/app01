// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// This portion of ioApp01pg handles all the
// i/o and manipulation for the vendor table. Any
// table manipulation should be added to this package as
// methods in IO_vendor



// Notes:
//  1. Any Database Query that returns "rows" must have an associated
//      rows.Close(). The best way to handle this is to do the query
//      immediately followed by "defer rows.Close()".  Queries that
//      return a "row" need not be closed.


// 2.   SQL requires OFFSET to follow LIMIT optionally (ie LIMIT n [OFFSET n])
// Generated: Thu Nov 14, 2019 11:17 for postgres Database

package ioApp01pgVendor

import (
    "database/sql"
	"fmt"
    _ "github.com/shopspring/decimal"
    "log"
	_ "strconv"

    
	_ "github.com/lib/pq"
    "app01pg/pkg/App01pgVendor"
    "app01pg/pkg/ioApp01pg"
)

//============================================================================
//                            IO_Vendor
//============================================================================

type IO_App01pgVendor struct {
    io          *ioApp01pg.IO_App01pg
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete deletes the row with keys from the provided record, rcd.
func (io *IO_App01pgVendor) RowDelete(rcd *App01pgVendor.App01pgVendor) error {
    var err         error
    var sqlStmt = "DELETE FROM public.vendor WHERE id = $1;\n"

    

	err = io.io.Exec(sqlStmt, rcd.Id)
	if err != nil {
        
		return fmt.Errorf("500. Internal Server Error")
	}

    
	return nil
}

//----------------------------------------------------------------------------
//                             Row Find
//----------------------------------------------------------------------------

// RowFind searches the Database for a matching row for the keys found in
// the given record and returns the output in that same record.
func (io *IO_App01pgVendor) RowFind(rcd *App01pgVendor.App01pgVendor) error {
    var err         error
    var sqlStmt     = "SELECT * FROM public.vendor WHERE id = $1;\n"

    

	row := io.io.QueryRow(sqlStmt, rcd.Id)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)

    
	return err
}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

// RowFirst returns the first row in the table, Vendor.
// If there are no rows in the table, then a blank/null record is returned
// without error.
func (io *IO_App01pgVendor) RowFirst(rcd *App01pgVendor.App01pgVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM public.vendor ORDER BY id ASC LIMIT 1;\n"

    

    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err == sql.ErrNoRows {
        
	    err = nil
    }

    
    return err
}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

func (io *IO_App01pgVendor) RowInsert(d *App01pgVendor.App01pgVendor) error {
    var err     error
    var sqlStmt = "INSERT INTO public.vendor (name, addr1, addr2, city, state, zip, curbal) VALUES ($1, $2, $3, $4, $5, $6, $7);\n"

    

    // Validate the input record.

    // Add it to the table.
    err = io.io.Exec(sqlStmt, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
    
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

    
	return err
}

//----------------------------------------------------------------------------
//                             Row Last
//----------------------------------------------------------------------------

func (io *IO_App01pgVendor) RowLast(rcd *App01pgVendor.App01pgVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM public.vendor ORDER BY id DESC LIMIT 1;\n"

    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err == sql.ErrNoRows {
        
	    err = nil
    }

    
    return err
}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

// RowNext returns the next row from the row given. If row after the current
// one does not exist, then the first row is returned.
func (io *IO_App01pgVendor) RowNext(rcd *App01pgVendor.App01pgVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM public.vendor WHERE id > $1 ORDER BY id ASC LIMIT 1;\n"

    

    row := io.io.QueryRow(sqlStmt, rcd.Id)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err != nil {
	    err = io.RowFirst(rcd)
	}

    
    return err
}

//----------------------------------------------------------------------------
//                             Row Page
//----------------------------------------------------------------------------

// RowPage returns a page of rows where a page size is the 'limit' parameter and
// 'offset' is the offset into the result set ordered by the main index. Both
// 'limit' and 'offset' are relative to 1. We return an address to the array
// rows (structs) so that we don't have the overhead of copying them everwhere.
func (io *IO_App01pgVendor) RowPage(offset int, limit int) ([]App01pgVendor.App01pgVendor, error) {
    var err         error
    var sqlStmt = "SELECT * FROM public.vendor ORDER BY id ASC LIMIT $1 OFFSET $2;\n"
    data := []App01pgVendor.App01pgVendor{}

    

    err = io.io.Query(
                    sqlStmt,
                    func(r *sql.Rows) {
                        var rcd     App01pgVendor.App01pgVendor
                        err = r.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
                        if err != nil {
                            log.Fatal(err)
                        } else {
                            data = append(data, rcd)
                        }
                    },
    limit,
                    offset)

    
    return data, err
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

func (io *IO_App01pgVendor) RowPrev(rcd *App01pgVendor.App01pgVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM public.vendor WHERE id < $1 ORDER BY id DESC LIMIT 1;\n"

    

    row := io.io.QueryRow(sqlStmt, rcd.Id)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err != nil {
	    err = io.RowLast(rcd)
	}

    
    return err
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

func (io *IO_App01pgVendor) RowUpdate(d *App01pgVendor.App01pgVendor) error {
    var err     error
    var sqlStmt = "INSERT INTO public.vendor (name, addr1, addr2, city, state, zip, curbal) VALUES ($1, $2, $3, $4, $5, $6, $7);\n"

    

    // Validate the input record.

    // Add it to the table.
    err = io.io.Exec(sqlStmt, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
    
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

    
	return err
}


//----------------------------------------------------------------------------
//                             Table Count
//----------------------------------------------------------------------------

func (io *IO_App01pgVendor) TableCount( ) (int, error) {
    var err         error
    var count       int
    var sqlStmt = "SELECT COUNT(*) FROM public.vendor;\n"

    

    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&count)
    if err != nil {
        
        return 0, err
    }

    
    return count, err
}

//----------------------------------------------------------------------------
//                             Table Create
//----------------------------------------------------------------------------

// TableCreate creates the table in the given database deleting the current
// table if present.
func (io *IO_App01pgVendor) TableCreate() error {
    var sqlStmt = "CREATE TABLE IF NOT EXISTS public.vendor (\n\tid\tSERIAL NOT NULL,\n\tname\tVARCHAR(30),\n\taddr1\tVARCHAR(30),\n\taddr2\tVARCHAR(30),\n\tcity\tVARCHAR(20),\n\tstate\tVARCHAR(10),\n\tzip\tVARCHAR(15),\n\tcurbal\tDEC(15,2),\n\tCONSTRAINT PK_vendor PRIMARY KEY(id)\n);\n"
    var err     error

    

    err = io.TableDelete()
    if err != nil {
        
        return err
    }
    err = io.io.Exec(sqlStmt)

    
    return err
}

//----------------------------------------------------------------------------
//                             Table Delete
//----------------------------------------------------------------------------

// TableDelete deletes the table in the given database if present.
func (io *IO_App01pgVendor) TableDelete() error {
    var sqlStmt = "DROP TABLE IF EXISTS public.vendor;\n"
    var err     error

    

    err = io.io.Exec(sqlStmt)

    
    return err
}


//----------------------------------------------------------------------------
//                             Table Scan
//----------------------------------------------------------------------------

// TableScan reads all the rows in the table applying a function to each of
// them.
func (io *IO_App01pgVendor) TableScan(apply func (rcd App01pgVendor.App01pgVendor) error) error {
    var err     error
    var rcd     App01pgVendor.App01pgVendor
    var sqlFirstStmt = "SELECT * FROM public.vendor ORDER BY id ASC LIMIT 1;\n"
    var sqlNextStmt = "SELECT * FROM public.vendor WHERE id > $1 ORDER BY id ASC LIMIT 1;\n"
    var row     *sql.Row

    


    

    row = io.io.QueryRow(sqlFirstStmt)
    for ;; {
        err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
        if err != nil {
            if err == sql.ErrNoRows {
                
                err = nil
            }
            break
        }
        // Warning: Next relies on the current record giving the key(s)
        // to find its position in the table. So, we pass a copy to apply().
        err = apply(rcd)
        if err != nil {
            break
        }
        row = io.io.QueryRow(sqlNextStmt, rcd.Id)
    }

    
    return err
}


//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new io struct.
func NewIoApp01pgVendor(io *ioApp01pg.IO_App01pg) *IO_App01pgVendor {
    db := &IO_App01pgVendor{}
    if io == nil {
        db.io = ioApp01pg.NewIoApp01pg()
    } else {
        db.io = io
    }
    return db
}

