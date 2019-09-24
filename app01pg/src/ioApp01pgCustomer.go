// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// This portion of ioApp01pg handles all the
// i/o and manipulation for the customer table. Any
// table manipulation should be added to this package as
// methods in IO_customer

// Notes:
//  1. Any Database Query that returns "rows" must have an associated
//      rows.Close(). The best way to handle this is to do the query
//      immediately followed by "defer rows.Close()".  Queries that
//      return a "row" need not be closed.

// 2.   SQL requires OFFSET to follow LIMIT optionally (ie LIMIT n [OFFSET n])
// Generated: Tue Sep 24, 2019 10:29 for postgres Database

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/shopspring/decimal"
	"log"
	_ "strconv"

	"github.com/2kranki/go_util"
	_ "github.com/lib/pq"
)

//============================================================================
//                            IO_Customer
//============================================================================

type IO_App01pgCustomer struct {
	io *IO_App01pg
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete deletes the row with keys from the provided record, rcd.
func (io *IO_App01pgCustomer) RowDelete(rcd *App01pgCustomer) error {
	var err error
	var sqlStmt = "DELETE FROM public.customer WHERE num = $1;\n"

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
func (io *IO_App01pgCustomer) RowFind(rcd *App01pgCustomer) error {
	var err error
	var sqlStmt = "SELECT * FROM public.customer WHERE num = $1;\n"

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
func (io *IO_App01pgCustomer) RowFirst(rcd *App01pgCustomer) error {
	var err error
	var sqlStmt = "SELECT * FROM public.customer ORDER BY num ASC LIMIT 1;\n"

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

func (io *IO_App01pgCustomer) RowInsert(d *App01pgCustomer) error {
	var err error
	var sqlStmt = "INSERT INTO public.customer (num, name, addr1, addr2, city, state, zip, curbal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);\n"

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

func (io *IO_App01pgCustomer) RowLast(rcd *App01pgCustomer) error {
	var err error
	var sqlStmt = "SELECT * FROM public.customer ORDER BY num DESC LIMIT 1;\n"

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
func (io *IO_App01pgCustomer) RowNext(rcd *App01pgCustomer) error {
	var err error
	var sqlStmt = "SELECT * FROM public.customer WHERE num > $1 ORDER BY num ASC LIMIT 1;\n"

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
func (io *IO_App01pgCustomer) RowPage(offset int, limit int) ([]App01pgCustomer, error) {
	var err error
	var sqlStmt = "SELECT * FROM public.customer ORDER BY num ASC LIMIT $1 OFFSET $2;\n"
	data := []App01pgCustomer{}

	log.Printf("ioCustomer.RowPage(%d,%d)\n", offset, limit)

	err = io.io.Query(
		sqlStmt,
		func(r *sql.Rows) {
			var rcd App01pgCustomer
			err = r.Scan(&rcd.Num, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
			if err != nil {
				log.Fatal(err)
			} else {
				data = append(data, rcd)
			}
		},
		limit,
		offset)

	log.Printf("...end ioCustomer.RowPage(%s)\n", util.ErrorString(err))
	return data, err
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

func (io *IO_App01pgCustomer) RowPrev(rcd *App01pgCustomer) error {
	var err error
	var sqlStmt = "SELECT * FROM public.customer WHERE num < $1 ORDER BY num DESC LIMIT 1;\n"

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

func (io *IO_App01pgCustomer) RowUpdate(d *App01pgCustomer) error {
	var err error
	var sqlStmt = "INSERT INTO public.customer (num, name, addr1, addr2, city, state, zip, curbal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);\n"

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

func (io *IO_App01pgCustomer) TableCount() (int, error) {
	var err error
	var count int
	var sqlStmt = "SELECT COUNT(*) FROM public.customer;\n"

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
func (io *IO_App01pgCustomer) TableCreate() error {
	var sqlStmt = "CREATE TABLE IF NOT EXISTS public.customer (\n\tnum\tINT NOT NULL,\n\tname\tVARCHAR(30),\n\taddr1\tVARCHAR(30),\n\taddr2\tVARCHAR(30),\n\tcity\tVARCHAR(20),\n\tstate\tVARCHAR(10),\n\tzip\tVARCHAR(15),\n\tcurbal\tDEC(15,2),\n\tCONSTRAINT PK_customer PRIMARY KEY(num)\n);\n"
	var err error

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
func (io *IO_App01pgCustomer) TableDelete() error {
	var sqlStmt = "DROP TABLE IF EXISTS public.customer;\n"
	var err error

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
func (io *IO_App01pgCustomer) TableScan(apply func(rcd App01pgCustomer) error) error {
	var sqlStmt = "DROP TABLE IF EXISTS public.customer;\n"
	var err error
	var rcd App01pgCustomer
	var sqlFirstStmt = "SELECT * FROM public.customer ORDER BY num ASC LIMIT 1;\n"
	var sqlNextStmt = "SELECT * FROM public.customer WHERE num > $1 ORDER BY num ASC LIMIT 1;\n"
	var row *sql.Row

	log.Printf("ioCustomer.TableScanner()\n")
	log.Printf("\tSQL:\n%s\n", sqlStmt)

	log.Printf("ioCustomer.RowFirst()\n")

	row = io.io.QueryRow(sqlFirstStmt)
	for {
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
func NewIoApp01pgCustomer(io *IO_App01pg) *IO_App01pgCustomer {
	db := &IO_App01pgCustomer{}
	if io == nil {
		db.io = NewIoApp01pg()
	} else {
		db.io = io
	}
	return db
}
