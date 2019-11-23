// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// This portion of ioApp01my handles all the
// i/o and manipulation for the vendor table. Any
// table manipulation should be added to this package as
// methods in IO_vendor

// Notes:
//  1. Any Database Query that returns "rows" must have an associated
//      rows.Close(). The best way to handle this is to do the query
//      immediately followed by "defer rows.Close()".  Queries that
//      return a "row" need not be closed.

// 2.   SQL requires OFFSET to follow LIMIT optionally (ie LIMIT n [OFFSET n])
// Generated: Sat Nov 23, 2019 00:27 for mysql Database

package ioApp01myVendor

import (
	"database/sql"
	"fmt"
	_ "github.com/shopspring/decimal"
	"log"
	_ "strconv"

	"app01my/pkg/App01myVendor"
	"app01my/pkg/ioApp01my"
	"github.com/2kranki/go_util"
	_ "github.com/go-sql-driver/mysql"
)

//============================================================================
//                            IO_Vendor
//============================================================================

type IO_App01myVendor struct {
	io *ioApp01my.IO_App01my
}

//----------------------------------------------------------------------------
//                             Row Delete
//----------------------------------------------------------------------------

// RowDelete deletes the row with keys from the provided record, rcd.
func (io *IO_App01myVendor) RowDelete(rcd *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "DELETE FROM vendor WHERE id = ?;\n"

	log.Printf("ioVendor.RowDelete()\n")

	err = io.io.Exec(sqlStmt, rcd.Id)
	if err != nil {
		log.Printf("...end ioVendor.RowDelete(Error:500) - Internal Error\n")
		return fmt.Errorf("500. Internal Server Error")
	}

	log.Printf("...end ioVendor.RowDelete()\n")
	return nil
}

//----------------------------------------------------------------------------
//                             Row Find
//----------------------------------------------------------------------------

// RowFind searches the Database for a matching row for the keys found in
// the given record and returns the output in that same record.
func (io *IO_App01myVendor) RowFind(rcd *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "SELECT * FROM vendor WHERE id = ?;\n"

	log.Printf("ioVendor.RowFind(%+v)\n", rcd)

	row := io.io.QueryRow(sqlStmt, rcd.Id)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)

	log.Printf("...end ioVendor.RowFind(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row First
//----------------------------------------------------------------------------

// RowFirst returns the first row in the table, Vendor.
// If there are no rows in the table, then a blank/null record is returned
// without error.
func (io *IO_App01myVendor) RowFirst(rcd *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "SELECT * FROM vendor ORDER BY id ASC LIMIT 1;\n"

	log.Printf("ioVendor.RowFirst()\n")

	row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err == sql.ErrNoRows {
		log.Printf("\tNo Rows found!\n")
		err = nil
	}

	log.Printf("...end ioVendor.RowFirst(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row Insert
//----------------------------------------------------------------------------

func (io *IO_App01myVendor) RowInsert(d *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "INSERT INTO vendor (name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?);\n"

	log.Printf("ioVendor.RowInsert(%+v)\n", d)
	log.Printf("\tSQL:\n%s\n", sqlStmt)

	// Validate the input record.

	// Add it to the table.
	err = io.io.Exec(sqlStmt, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
		log.Printf("...end ioVendor.RowInsert(Error:500) - Internal Error\n")
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

	log.Printf("...end ioVendor.RowInsert(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row Last
//----------------------------------------------------------------------------

func (io *IO_App01myVendor) RowLast(rcd *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "SELECT * FROM vendor ORDER BY id DESC LIMIT 1;\n"

	log.Printf("ioVendor.RowLast()\n")
	row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err == sql.ErrNoRows {
		log.Printf("\tNo Rows found!\n")
		err = nil
	}

	log.Printf("...end ioVendor.RowLast(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row Next
//----------------------------------------------------------------------------

// RowNext returns the next row from the row given. If row after the current
// one does not exist, then the first row is returned.
func (io *IO_App01myVendor) RowNext(rcd *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "SELECT * FROM vendor WHERE id > ? ORDER BY id ASC LIMIT 1;\n"

	log.Printf("ioVendor.RowNext(%+v)\n", rcd)

	row := io.io.QueryRow(sqlStmt, rcd.Id)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err != nil {
		err = io.RowFirst(rcd)
	}

	log.Printf("...end ioVendor.RowNext(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row Page
//----------------------------------------------------------------------------

// RowPage returns a page of rows where a page size is the 'limit' parameter and
// 'offset' is the offset into the result set ordered by the main index. Both
// 'limit' and 'offset' are relative to 1. We return an address to the array
// rows (structs) so that we don't have the overhead of copying them everwhere.
func (io *IO_App01myVendor) RowPage(offset int, limit int) ([]App01myVendor.App01myVendor, error) {
	var err error
	var sqlStmt = "SELECT * FROM vendor ORDER BY id ASC LIMIT ? OFFSET ?;\n"
	data := []App01myVendor.App01myVendor{}

	log.Printf("ioVendor.RowPage(%d,%d)\n", offset, limit)

	err = io.io.Query(
		sqlStmt,
		func(r *sql.Rows) {
			var rcd App01myVendor.App01myVendor
			err = r.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
			if err != nil {
				log.Fatal(err)
			} else {
				data = append(data, rcd)
			}
		},
		limit,
		offset)

	log.Printf("...end ioVendor.RowPage(%s)\n", util.ErrorString(err))
	return data, err
}

//----------------------------------------------------------------------------
//                             Row Prev
//----------------------------------------------------------------------------

func (io *IO_App01myVendor) RowPrev(rcd *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "SELECT * FROM vendor WHERE id < ? ORDER BY id DESC LIMIT 1;\n"

	log.Printf("ioVendor.RowPrev(%+v)\n", rcd)

	row := io.io.QueryRow(sqlStmt, rcd.Id)

	err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
	if err != nil {
		err = io.RowLast(rcd)
	}

	log.Printf("...end ioVendor.RowPrev(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Row Update
//----------------------------------------------------------------------------

func (io *IO_App01myVendor) RowUpdate(d *App01myVendor.App01myVendor) error {
	var err error
	var sqlStmt = "INSERT INTO vendor (name, addr1, addr2, city, state, zip, curbal) VALUES (?, ?, ?, ?, ?, ?, ?);\n"

	log.Printf("ioVendor.RowUpdate(%+v)\n", d)

	// Validate the input record.

	// Add it to the table.
	err = io.io.Exec(sqlStmt, d.Name, d.Addr1, d.Addr2, d.City, d.State, d.Zip, d.Curbal)
	if err != nil {
		log.Printf("...end ioVendor.RowUpdate(Error:500) - Internal Error\n")
		err = fmt.Errorf("500. Internal Server Error. %s\n", err.Error())
	}

	log.Printf("...end ioVendor.RowUpdate(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Table Count
//----------------------------------------------------------------------------

func (io *IO_App01myVendor) TableCount() (int, error) {
	var err error
	var count int
	var sqlStmt = "SELECT COUNT(*) FROM vendor;\n"

	log.Printf("ioVendor.TableCount()\n")

	row := io.io.QueryRow(sqlStmt)

	err = row.Scan(&count)
	if err != nil {

		log.Printf("...end ioVendor.TableCount(%s) %d\n", util.ErrorString(err), count)
		return 0, err
	}

	log.Printf("...end ioVendor.TableCount(%s) %d\n", util.ErrorString(err), count)
	return count, err
}

//----------------------------------------------------------------------------
//                             Table Create
//----------------------------------------------------------------------------

// TableCreate creates the table in the given database deleting the current
// table if present.
func (io *IO_App01myVendor) TableCreate() error {
	var sqlStmt = "CREATE TABLE IF NOT EXISTS vendor (\n\tid\tINT NOT NULL AUTO_INCREMENT,\n\tname\tNVARCHAR(30),\n\taddr1\tNVARCHAR(30),\n\taddr2\tNVARCHAR(30),\n\tcity\tNVARCHAR(20),\n\tstate\tNVARCHAR(10),\n\tzip\tNVARCHAR(15),\n\tcurbal\tDEC(15,2),\n\tCONSTRAINT PK_vendor PRIMARY KEY(id)\n);\n"
	var err error

	log.Printf("ioVendor.TableCreate()\n")
	log.Printf("\tSQL:\n%s\n", sqlStmt)

	err = io.TableDelete()
	if err != nil {
		log.Printf("...end ioVendor.TableCreate(Error:%s)\n", err.Error())
		return err
	}
	err = io.io.Exec(sqlStmt)

	log.Printf("...end ioVendor.TableCreate(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Table Delete
//----------------------------------------------------------------------------

// TableDelete deletes the table in the given database if present.
func (io *IO_App01myVendor) TableDelete() error {
	var sqlStmt = "DROP TABLE IF EXISTS vendor;\n"
	var err error

	log.Printf("ioVendor.TableDelete()\n")
	log.Printf("\tSQL:\n%s\n", sqlStmt)

	err = io.io.Exec(sqlStmt)

	log.Printf("...end ioVendor.TableDelete(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                             Table Scan
//----------------------------------------------------------------------------

// TableScan reads all the rows in the table applying a function to each of
// them.
func (io *IO_App01myVendor) TableScan(apply func(rcd App01myVendor.App01myVendor) error) error {
	var err error
	var rcd App01myVendor.App01myVendor
	var sqlFirstStmt = "SELECT * FROM vendor ORDER BY id ASC LIMIT 1;\n"
	var sqlNextStmt = "SELECT * FROM vendor WHERE id > ? ORDER BY id ASC LIMIT 1;\n"
	var row *sql.Row

	log.Printf("ioVendor.TableScanner()\n")
	log.Printf("\tSQL:\n%s\n", sqlFirstStmt)

	log.Printf("ioVendor.RowFirst()\n")

	row = io.io.QueryRow(sqlFirstStmt)
	for {
		err = row.Scan(&rcd.Id, &rcd.Name, &rcd.Addr1, &rcd.Addr2, &rcd.City, &rcd.State, &rcd.Zip, &rcd.Curbal)
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
		row = io.io.QueryRow(sqlNextStmt, rcd.Id)
	}

	log.Printf("...end ioVendor.TableDelete(%s)\n", util.ErrorString(err))
	return err
}

//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

// New creates a new io struct.
func NewIoApp01myVendor(io *ioApp01my.IO_App01my) *IO_App01myVendor {
	db := &IO_App01myVendor{}
	if io == nil {
		db.io = ioApp01my.NewIoApp01my()
	} else {
		db.io = io
	}
	return db
}
