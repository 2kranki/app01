// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// This portion of ioApp01ms handles all the
// i/o and manipulation for the customer table. Any
// table manipulation should be added to this package as
// methods in IO_customer



// Notes:
//  1. Any Database Query that returns "rows" must have an associated
//      rows.Close(). The best way to handle this is to do the query
//      immediately followed by "defer rows.Close()".  Queries that
//      return a "row" need not be closed.

// 2.   T-SQL does not seem to support LIMIT or OFFSET in SQL Server 2017. So, you
//      have to use an ORDER BY clause followed by an OFFSET clause optionally
//      followed by the FETCH clause (ie ORDER BY xx [OFFSET n ROWS [FETCH NEXT n ROWS ONLY]]).
// Generated: Fri Oct 18, 2019 14:50 for mssql Database

package ioApp01msCustomer

import (
    "database/sql"
	"fmt"
    _ "github.com/shopspring/decimal"
    "log"
    _ "strconv"

    "github.com/2kranki/go_util"
	_ "github.com/denisenkom/go-mssqldb"
    "app01ms/pkg/App01msCustomer"
    "app01ms/pkg/ioApp01ms"
)

//============================================================================
//                            IO_Customer
//============================================================================

type IO_App01msCustomer struct {
    io          *ioApp01ms.IO_App01ms
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete deletes the row with keys from the provided record, rcd.
func (io *IO_App01msCustomer) RowDelete(rcd *App01msCustomer.App01msCustomer) error {
    var err         error
    var sqlStmt = "DELETE FROM dbo.customer WHERE num = ?;\n"

    log.Printf("ioCustomer.RowDelete()\n")

	err = io.io.Exec(sqlStmt, rcd.Num)
	if err != nil {
        log.Printf("...end ioCustomer.RowDelete(Error:500) - Internal Error\n")
		return fmt.Errorf("500. Internal Server Error")
	}

    log.Printf("...end ioCustomer.RowDelete()\n")
	return nil
}

//----------------------------------------------------------------------------
//                             Row Find
//----------------------------------------------------------------------------

// RowFind searches the Database for a matching row for the keys found in
// the given record and returns the output in that same record.
func (io *IO_App01msCustomer) RowFind(rcd *App01msCustomer.App01msCustomer) error {
    var err         error
    var sqlStmt     = "SELECT * FROM dbo.customer WHERE num = ?;\n"

    log.Printf("ioCustomer.RowFind(%+v)\n", rcd)

	row := io.io.QueryRow(sqlStmt, rcd.Num)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)

    log.Printf("...end ioCustomer.RowFind(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

// RowFirst returns the first row in the table, Customer.
// If there are no rows in the table, then a blank/null record is returned
// without error.
func (io *IO_App01msCustomer) RowFirst(rcd *App01msCustomer.App01msCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Customer ORDER BY num ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"

    log.Printf("ioCustomer.RowFirst()\n")

    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err == sql.ErrNoRows {
        log.Printf("\tNo Rows found!\n")
	    err = nil
    }

    log.Printf("...end ioCustomer.RowFirst(%s)\n", util.ErrorString(err))
    return err
}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

func (io *IO_App01msCustomer) RowInsert(d *App01msCustomer.App01msCustomer) error {
    var err     error
    var sqlStmt = "INSERT INTO dbo.customer (num, name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?, ?);\n"

    log.Printf("ioCustomer.RowInsert(%+v)\n", d)
        log.Printf("\tSQL:\n%s\n", sqlStmt)

    // Validate the input record.

    // Add it to the table.
    err = io.io.Exec(sqlStmt, d.Num, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
    log.Printf("...end ioCustomer.RowInsert(Error:500) - Internal Error\n")
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

    log.Printf("...end ioCustomer.RowInsert(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row Last
//----------------------------------------------------------------------------

func (io *IO_App01msCustomer) RowLast(rcd *App01msCustomer.App01msCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Customer ORDER BY num DESC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"

    log.Printf("ioCustomer.RowLast()\n")
    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err == sql.ErrNoRows {
        log.Printf("\tNo Rows found!\n")
	    err = nil
    }

    log.Printf("...end ioCustomer.RowLast(%s)\n", util.ErrorString(err))
    return err
}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

// RowNext returns the next row from the row given. If row after the current
// one does not exist, then the first row is returned.
func (io *IO_App01msCustomer) RowNext(rcd *App01msCustomer.App01msCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Customer WHERE num > ? ORDER BY num ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"

    log.Printf("ioCustomer.RowNext(%+v)\n", rcd)

    row := io.io.QueryRow(sqlStmt, rcd.Num)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err != nil {
	    err = io.RowFirst(rcd)
	}

    log.Printf("...end ioCustomer.RowNext(%s)\n", util.ErrorString(err))
    return err
}

//----------------------------------------------------------------------------
//                             Row Page
//----------------------------------------------------------------------------

// RowPage returns a page of rows where a page size is the 'limit' parameter and
// 'offset' is the offset into the result set ordered by the main index. Both
// 'limit' and 'offset' are relative to 1. We return an address to the array
// rows (structs) so that we don't have the overhead of copying them everwhere.
func (io *IO_App01msCustomer) RowPage(offset int, limit int) ([]App01msCustomer.App01msCustomer, error) {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Customer ORDER BY num ASC OFFSET ? ROWS FETCH NEXT ? ROWS ONLY;\n"
    data := []App01msCustomer.App01msCustomer{}

    log.Printf("ioCustomer.RowPage(%d,%d)\n",offset,limit)

    err = io.io.Query(
                    sqlStmt,
                    func(r *sql.Rows) {
                        var rcd     App01msCustomer.App01msCustomer
                        err = r.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
                        if err != nil {
                            log.Fatal(err)
                        } else {
                            data = append(data, rcd)
                        }
                    },
    offset,
                    limit)
    

    log.Printf("...end ioCustomer.RowPage(%s)\n", util.ErrorString(err))
    return data, err
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

func (io *IO_App01msCustomer) RowPrev(rcd *App01msCustomer.App01msCustomer) error {
    var err         error
    var sqlStmt = "SELECT * FROM dbo.Customer WHERE num < ? ORDER BY num DESC OFFSET 0 ROWS;\n"

    log.Printf("ioCustomer.RowPrev(%+v)\n", rcd)

    row := io.io.QueryRow(sqlStmt, rcd.Num)

	err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err != nil {
	    err = io.RowLast(rcd)
	}

    log.Printf("...end ioCustomer.RowPrev(%s)\n", util.ErrorString(err))
    return err
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

func (io *IO_App01msCustomer) RowUpdate(d *App01msCustomer.App01msCustomer) error {
    var err     error
    var sqlStmt = "INSERT INTO dbo.customer (num, name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?, ?);\n"

    log.Printf("ioCustomer.RowUpdate(%+v)\n", d)

    // Validate the input record.

    // Add it to the table.
    err = io.io.Exec(sqlStmt, d.Num, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
    log.Printf("...end ioCustomer.RowUpdate(Error:500) - Internal Error\n")
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

    log.Printf("...end ioCustomer.RowUpdate(%s)\n", util.ErrorString(err))
	return err
}


//----------------------------------------------------------------------------
//                             Table Count
//----------------------------------------------------------------------------

func (io *IO_App01msCustomer) TableCount( ) (int, error) {
    var err         error
    var count       int
    var sqlStmt = "SELECT COUNT(*) FROM dbo.customer;\n"

    log.Printf("ioCustomer.TableCount()\n")

    row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&count)
    if err != nil {
        
            log.Printf("...end ioCustomer.TableCount(%s) %d\n", util.ErrorString(err), count)
        return 0, err
    }

    log.Printf("...end ioCustomer.TableCount(%s) %d\n", util.ErrorString(err), count)
    return count, err
}

//----------------------------------------------------------------------------
//                             Table Create
//----------------------------------------------------------------------------

// TableCreate creates the table in the given database deleting the current
// table if present.
func (io *IO_App01msCustomer) TableCreate() error {
    var sqlStmt = "CREATE TABLE dbo.Customer (\n\tnum\tINT NOT NULL,\n\tname\tNVARCHAR(30),\n\taddr1\tNVARCHAR(30),\n\taddr2\tNVARCHAR(30),\n\tcity\tNVARCHAR(20),\n\tstate\tNVARCHAR(10),\n\tzip\tNVARCHAR(15),\n\tcurbal\tDEC(15,2),\n\tCONSTRAINT PK_Customer PRIMARY KEY(num)\n);\n"
    var err     error

    log.Printf("ioCustomer.TableCreate()\n")
        log.Printf("\tSQL:\n%s\n", sqlStmt)

    err = io.TableDelete()
    if err != nil {
        log.Printf("...end ioCustomer.TableCreate(Error:%s)\n", err.Error())
        return err
    }
    err = io.io.Exec(sqlStmt)

    log.Printf("...end ioCustomer.TableCreate(%s)\n", util.ErrorString(err))
    return err
}

//----------------------------------------------------------------------------
//                             Table Delete
//----------------------------------------------------------------------------

// TableDelete deletes the table in the given database if present.
func (io *IO_App01msCustomer) TableDelete() error {
    var sqlStmt = "DROP TABLE IF EXISTS dbo.customer;\n"
    var err     error

    log.Printf("ioCustomer.TableDelete()\n")
        log.Printf("\tSQL:\n%s\n", sqlStmt)

    err = io.io.Exec(sqlStmt)

    log.Printf("...end ioCustomer.TableDelete(%s)\n", util.ErrorString(err))
    return err
}


//----------------------------------------------------------------------------
//                             Table Scan
//----------------------------------------------------------------------------

// TableScan reads all the rows in the table applying a function to each of
// them.
func (io *IO_App01msCustomer) TableScan(apply func (rcd App01msCustomer.App01msCustomer) error) error {
    var sqlStmt = "DROP TABLE IF EXISTS dbo.customer;\n"
    var err     error
    var rcd     App01msCustomer.App01msCustomer
    var sqlFirstStmt = "SELECT * FROM dbo.Customer ORDER BY num ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"
    var sqlNextStmt = "SELECT * FROM dbo.Customer WHERE num > ? ORDER BY num ASC OFFSET 0 ROWS FETCH NEXT 1 ROW ONLY;\n"
    var row     *sql.Row

    log.Printf("ioCustomer.TableScanner()\n")
        log.Printf("\tSQL:\n%s\n", sqlStmt)


    log.Printf("ioCustomer.RowFirst()\n")

    row = io.io.QueryRow(sqlFirstStmt)
    for ;; {
        err = row.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
        if err != nil {
            if err == sql.ErrNoRows {
                log.Printf("\tNo Rows found!\n")
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

    log.Printf("...end ioCustomer.TableDelete(%s)\n", util.ErrorString(err))
    return err
}


//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new io struct.
func NewIoApp01msCustomer(io *ioApp01ms.IO_App01ms) *IO_App01msCustomer {
    db := &IO_App01msCustomer{}
    if io == nil {
        db.io = ioApp01ms.NewIoApp01ms()
    } else {
        db.io = io
    }
    return db
}

