// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// This portion of ioApp01ma handles all the
// i/o and manipulation for the customer table. Any
// table manipulation should be added to this package as
// methods in IO_customer



// Notes:
//  1. Any Database Query that returns "rows" must have an associated
//      rows.Close(). The best way to handle this is to do the query
//      immediately followed by "defer rows.Close()".  Queries that
//      return a "row" need not be closed.


// 2.   SQL requires OFFSET to follow LIMIT optionally (ie LIMIT n [OFFSET n])
// Generated: Thu Nov 14, 2019 11:17 for mariadb Database

package ioApp01maCustomer

import (
    "database/sql"
	"fmt"
    _ "github.com/shopspring/decimal"
    "log"
	_ "strconv"

    
	_ "github.com/go-sql-driver/mysql"
    "app01ma/pkg/App01maCustomer"
    "app01ma/pkg/ioApp01ma"
)

//============================================================================
//                            IO_Customer
//============================================================================

type IO_App01maCustomer struct {
    io          *ioApp01ma.IO_App01ma
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete deletes the row with keys from the provided record, rcd.
func (io *IO_App01maCustomer) RowDelete(rcd *App01maCustomer.App01maCustomer) error {
    var err         error
    var sqlStmt = "DELETE FROM customer WHERE num = ?;\n"

    

	err = io.io.Exec(sqlStmt, rcd.Num)
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
func (io *IO_App01maCustomer) RowFind(rcd *App01maCustomer.App01maCustomer) error {
    var err         error
    var sqlStmt     = "SELECT * FROM customer WHERE num = ?;\n"

    

	row := io.io.QueryRow(sqlStmt, rcd.Num)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)

    
	return err
}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

// RowFirst returns the first row in the table, Customer.
// If there are no rows in the table, then a blank/null record is returned
// without error.
func (io *IO_App01maCustomer) RowFirst(rcd *App01maCustomer.App01maCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM customer ORDER BY num ASC LIMIT 1;\n"

    

    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err == sql.ErrNoRows {
        
	    err = nil
    }

    
    return err
}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

func (io *IO_App01maCustomer) RowInsert(d *App01maCustomer.App01maCustomer) error {
    var err     error
    var sqlStmt = "INSERT INTO customer (num, name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?, ?);\n"

    

    // Validate the input record.

    // Add it to the table.
    err = io.io.Exec(sqlStmt, d.Num, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
    
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

    
	return err
}

//----------------------------------------------------------------------------
//                             Row Last
//----------------------------------------------------------------------------

func (io *IO_App01maCustomer) RowLast(rcd *App01maCustomer.App01maCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM customer ORDER BY num DESC LIMIT 1;\n"

    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
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
func (io *IO_App01maCustomer) RowNext(rcd *App01maCustomer.App01maCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM customer WHERE num > ? ORDER BY num ASC LIMIT 1;\n"

    

    row := io.io.QueryRow(sqlStmt, rcd.Num)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
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
func (io *IO_App01maCustomer) RowPage(offset int, limit int) ([]App01maCustomer.App01maCustomer, error) {
    var err         error
    var sqlStmt = "SELECT * FROM customer ORDER BY num ASC LIMIT ? OFFSET ?;\n"
    data := []App01maCustomer.App01maCustomer{}

    

    err = io.io.Query(
                    sqlStmt,
                    func(r *sql.Rows) {
                        var rcd     App01maCustomer.App01maCustomer
                        err = r.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
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

func (io *IO_App01maCustomer) RowPrev(rcd *App01maCustomer.App01maCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM customer WHERE num < ? ORDER BY num DESC LIMIT 1;\n"

    

    row := io.io.QueryRow(sqlStmt, rcd.Num)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err != nil {
	    err = io.RowLast(rcd)
	}

    
    return err
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

func (io *IO_App01maCustomer) RowUpdate(d *App01maCustomer.App01maCustomer) error {
    var err     error
    var sqlStmt = "INSERT INTO customer (num, name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?, ?);\n"

    

    // Validate the input record.

    // Add it to the table.
    err = io.io.Exec(sqlStmt, d.Num, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
    
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

    
	return err
}


//----------------------------------------------------------------------------
//                             Table Count
//----------------------------------------------------------------------------

func (io *IO_App01maCustomer) TableCount( ) (int, error) {
    var err         error
    var count       int
    var sqlStmt = "SELECT COUNT(*) FROM customer;\n"

    

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
func (io *IO_App01maCustomer) TableCreate() error {
    var sqlStmt = "CREATE TABLE IF NOT EXISTS customer (\n\tnum\tINT NOT NULL,\n\tname\tVARCHAR(30),\n\taddr1\tVARCHAR(30),\n\taddr2\tVARCHAR(30),\n\tcity\tVARCHAR(20),\n\tstate\tVARCHAR(10),\n\tzip\tVARCHAR(15),\n\tcurbal\tDEC(15,2),\n\tCONSTRAINT PK_customer PRIMARY KEY(num)\n);\n"
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
func (io *IO_App01maCustomer) TableDelete() error {
    var sqlStmt = "DROP TABLE IF EXISTS customer;\n"
    var err     error

    

    err = io.io.Exec(sqlStmt)

    
    return err
}


//----------------------------------------------------------------------------
//                             Table Scan
//----------------------------------------------------------------------------

// TableScan reads all the rows in the table applying a function to each of
// them.
func (io *IO_App01maCustomer) TableScan(apply func (rcd App01maCustomer.App01maCustomer) error) error {
    var err     error
    var rcd     App01maCustomer.App01maCustomer
    var sqlFirstStmt = "SELECT * FROM customer ORDER BY num ASC LIMIT 1;\n"
    var sqlNextStmt = "SELECT * FROM customer WHERE num > ? ORDER BY num ASC LIMIT 1;\n"
    var row     *sql.Row

    


    

    row = io.io.QueryRow(sqlFirstStmt)
    for ;; {
        err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
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
        row = io.io.QueryRow(sqlNextStmt, rcd.Num)
    }

    
    return err
}


//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new io struct.
func NewIoApp01maCustomer(io *ioApp01ma.IO_App01ma) *IO_App01maCustomer {
    db := &IO_App01maCustomer{}
    if io == nil {
        db.io = ioApp01ma.NewIoApp01ma()
    } else {
        db.io = io
    }
    return db
}

