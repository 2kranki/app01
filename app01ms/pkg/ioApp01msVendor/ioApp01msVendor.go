// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// This portion of ioApp01ms handles all the
// i/o and manipulation for the vendor table. Any
// table manipulation should be added to this package as
// methods in IO_vendor



// Notes:
//  1. Any Database Query that returns "rows" must have an associated
//      rows.Close(). The best way to handle this is to do the query
//      immediately followed by "defer rows.Close()".  Queries that
//      return a "row" need not be closed.

// 2.   T-SQL does not seem to support LIMIT or OFFSET in SQL Server 2017. So, you
//      have to use an ORDER BY clause followed by an OFFSET clause optionally
//      followed by the FETCH clause (ie ORDER BY xx [OFFSET n ROWS [FETCH NEXT n ROWS ONLY]]).
// Generated: Thu Nov 14, 2019 11:17 for mssql Database

package ioApp01msVendor

import (
    "database/sql"
	"fmt"
    _ "github.com/shopspring/decimal"
    "log"
	_ "strconv"

    
	_ "github.com/denisenkom/go-mssqldb"
    "app01ms/pkg/App01msVendor"
    "app01ms/pkg/ioApp01ms"
)

//============================================================================
//                            IO_Vendor
//============================================================================

type IO_App01msVendor struct {
    io          *ioApp01ms.IO_App01ms
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete deletes the row with keys from the provided record, rcd.
func (io *IO_App01msVendor) RowDelete(rcd *App01msVendor.App01msVendor) error {
    var err         error
    var sqlStmt = "DELETE FROM dbo.vendor WHERE id = ?;\n"

    

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
func (io *IO_App01msVendor) RowFind(rcd *App01msVendor.App01msVendor) error {
    var err         error
    var sqlStmt     = "SELECT * FROM dbo.vendor WHERE id = ?;\n"

    

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
func (io *IO_App01msVendor) RowFirst(rcd *App01msVendor.App01msVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Vendor ORDER BY id ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"

    

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

func (io *IO_App01msVendor) RowInsert(d *App01msVendor.App01msVendor) error {
    var err     error
    var sqlStmt = "INSERT INTO dbo.vendor (name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?);\n"

    

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

func (io *IO_App01msVendor) RowLast(rcd *App01msVendor.App01msVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Vendor ORDER BY id DESC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"

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
func (io *IO_App01msVendor) RowNext(rcd *App01msVendor.App01msVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Vendor WHERE id > ? ORDER BY id ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"

    

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
func (io *IO_App01msVendor) RowPage(offset int, limit int) ([]App01msVendor.App01msVendor, error) {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Vendor ORDER BY id ASC OFFSET ? ROWS FETCH NEXT ? ROWS ONLY;\n"
    data := []App01msVendor.App01msVendor{}

    

    err = io.io.Query(
                    sqlStmt,
                    func(r *sql.Rows) {
                        var rcd     App01msVendor.App01msVendor
                        err = r.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
                        if err != nil {
                            log.Fatal(err)
                        } else {
                            data = append(data, rcd)
                        }
                    },
    offset,
                    limit)
    

    
    return data, err
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

func (io *IO_App01msVendor) RowPrev(rcd *App01msVendor.App01msVendor) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Vendor WHERE id < ? ORDER BY id DESC OFFSET 0 ROWS;\n"

    

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

func (io *IO_App01msVendor) RowUpdate(d *App01msVendor.App01msVendor) error {
    var err     error
    var sqlStmt = "INSERT INTO dbo.vendor (name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?);\n"

    

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

func (io *IO_App01msVendor) TableCount( ) (int, error) {
    var err         error
    var count       int
    var sqlStmt = "SELECT COUNT(*) FROM dbo.vendor;\n"

    

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
func (io *IO_App01msVendor) TableCreate() error {
    var sqlStmt = "CREATE TABLE dbo.Vendor (\n\tid\tINT NOT NULL IDENTITY(1,1),\n\tname\tNVARCHAR(30),\n\taddr1\tNVARCHAR(30),\n\taddr2\tNVARCHAR(30),\n\tcity\tNVARCHAR(20),\n\tstate\tNVARCHAR(10),\n\tzip\tNVARCHAR(15),\n\tcurbal\tDEC(15,2),\n\tCONSTRAINT PK_Vendor PRIMARY KEY(id)\n);\n"
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
func (io *IO_App01msVendor) TableDelete() error {
    var sqlStmt = "DROP TABLE IF EXISTS dbo.vendor;\n"
    var err     error

    

    err = io.io.Exec(sqlStmt)

    
    return err
}


//----------------------------------------------------------------------------
//                             Table Scan
//----------------------------------------------------------------------------

// TableScan reads all the rows in the table applying a function to each of
// them.
func (io *IO_App01msVendor) TableScan(apply func (rcd App01msVendor.App01msVendor) error) error {
    var err     error
    var rcd     App01msVendor.App01msVendor
    var sqlFirstStmt = "SELECT * FROM dbo.Vendor ORDER BY id ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"
    var sqlNextStmt = "SELECT * FROM dbo.Vendor WHERE id > ? ORDER BY id ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"
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
func NewIoApp01msVendor(io *ioApp01ms.IO_App01ms) *IO_App01msVendor {
    db := &IO_App01msVendor{}
    if io == nil {
        db.io = ioApp01ms.NewIoApp01ms()
    } else {
        db.io = io
    }
    return db
}

