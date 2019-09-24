// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ma contains all the methods for manipulating an SQL
// Database such as connections, database creation and deletion if needed. It
// does not include the specific Table Maintenance Methods. Those are maintained
// in separate packages.

// References:
//      https://golang.org/src/database/sql/doc.txt
//      https://github.com/golang/go/wiki/SQLDrivers
//      http://go-database-sql.org/

// Remarks:
//  *   A Golang Database driver may or may not support multiple statements per
//      request. So, it is safest to only assume that it will perform 1 statement
//      at a time.
//  *   We recommend certain naming conventions. First all supplied names should
//      be lower-case. You should separate words with an '_' if you use full words
//      in the name.

// Generated: Tue Sep 24, 2019 10:29 for mariadb Database

package main

import (
	"database/sql"
	"fmt"

	"log"

	"strconv"
	"strings"
	// "time" is only needed for Docker support and "sqlite" is the only
	//  database server not using it.
	"time"

	"github.com/2kranki/go_util"
	"github.com/go-sql-driver/mysql"
)

const connect_retries = 100

//============================================================================
//                            IO_App01ma
//============================================================================

type IO_App01ma struct {
	dbSql    *sql.DB
	dbName   string
	dbPW     string
	dbPort   string
	dbServer string
	dbUser   string
}

func (io *IO_App01ma) Name() string {
	return io.dbName
}
func (io *IO_App01ma) SetName(str string) {
	io.dbName = str
}

func (io *IO_App01ma) PW() string {
	return io.dbPW
}
func (io *IO_App01ma) SetPW(str string) {
	io.dbPW = str
}

func (io *IO_App01ma) Port() string {
	return io.dbPort
}
func (io *IO_App01ma) SetPort(str string) {
	io.dbPort = str
}

func (io *IO_App01ma) Server() string {
	return io.dbServer
}
func (io *IO_App01ma) SetServer(str string) {
	io.dbServer = str
}

func (io *IO_App01ma) Sql() *sql.DB {
	return io.dbSql
}

func (io *IO_App01ma) User() string {
	return io.dbUser
}
func (io *IO_App01ma) SetUser(str string) {
	io.dbUser = str
}

//============================================================================
//                              Miscellaneous
//============================================================================

func (io *IO_App01ma) FloatToString(num float64) string {
	s := fmt.Sprintf("%.4f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func (io *IO_App01ma) StringToFloat(str string) float64 {
	var num float64
	num, _ = strconv.ParseFloat(str, 64)
	return num
}

// Set up default parameters for the needed SQL Type.
func (io *IO_App01ma) DefaultParms() {
	io.SetPort("4306")
	io.SetPW("Passw0rd")
	io.SetServer("localhost")
	io.SetUser("root")
	//io.SetName("App01ma")
}

//============================================================================
//                              Database Methods
//============================================================================

//----------------------------------------------------------------------------
//								Connect - MariaDB
//----------------------------------------------------------------------------

// Connect connects the database/sql/driver to the appropriate
// database server using the given parameters.
func (io *IO_App01ma) Connect(dbName string) error {
	var err error

	dbName = strings.ToLower(dbName)

	// Set up connection string, connStr.
	cfg := mysql.NewConfig()
	cfg.User = io.dbUser
	cfg.Passwd = io.dbPW
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", io.dbServer, io.dbPort)
	if len(dbName) > 0 {
		cfg.DBName = dbName
	}
	connStr := cfg.FormatDSN()

	// Allow for the Docker Container to get operational.
	for i := 0; i < connect_retries; i++ {
		log.Printf("\tConnecting %d to mariadb with %s...\n", i, connStr)
		io.dbSql, err = sql.Open("mysql", connStr)
		if err == nil {
			err = io.dbSql.Ping()
			if err == nil {
				break
			}
			io.Disconnect()
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("Error: Cannot Connect: %s\n", err.Error())
	}

	log.Printf("Pinging Server...\n")

	err = io.dbSql.Ping()
	if err != nil {
		io.Disconnect()
		return fmt.Errorf("Ping Error: Cannot Ping: %s\n", err.Error())
	}
	io.SetName(dbName)

	return nil
}

//----------------------------------------------------------------------------
//								Disconnect
//----------------------------------------------------------------------------

// Disconnect() cleans up anything that needs to be
// accomplished before the database is closed
// and then closes the database connection.
func (io *IO_App01ma) Disconnect() error {
	var err error

	log.Printf("\tDisconnecting from Database\n")

	if io.IsConnected() {
		err = io.dbSql.Close()
		io.dbSql = nil
	} else {
		err = fmt.Errorf("Error: Database was not connected!")
	}

	return err
}

//----------------------------------------------------------------------------
//								IsConnected
//----------------------------------------------------------------------------

func (io *IO_App01ma) IsConnected() bool {
	if io.dbSql != nil {
		return true
	}
	return false
}

//============================================================================
//                        Database Maintenance
//============================================================================

//----------------------------------------------------------------------------
//								Create
//----------------------------------------------------------------------------

// DatabaseCreate creates the database within the SQL server if needed and
// opens a connection to it.
func (io *IO_App01ma) DatabaseCreate(dbName string) error {
	var err error
	var str util.StringBuilder

	log.Printf("DatabaseCreate(%s)\n", dbName)
	if len(dbName) == 0 {
		return fmt.Errorf("Error: Missing database name for DatabaseCreate()!")
	}

	dbName = strings.ToLower(dbName)

	// Connect without a database specified if needed.
	if len(io.dbName) > 0 || io.dbSql == nil {
		io.Disconnect()
		err = io.Connect("")
		if err != nil {
			return err
		}
	}

	if io.IsDatabaseDefined(dbName) {
		io.Disconnect()
		err = io.Connect(dbName)
		return err
	}

	// Build the Create Database SQL Statement.
	str.WriteStringf("CREATE DATABASE IF NOT EXISTS %s;", dbName)

	// Create the database.
	err = io.Exec(str.String())
	if err != nil {
		io.Disconnect()
		return err
	}
	time.Sleep(5 * time.Second) // Give it time to get done.
	if !io.IsDatabaseDefined(dbName) {
		io.Disconnect()
		return fmt.Errorf("Error - Could not verify database, %s, exists!", dbName)
	}

	// Now disconnect from the connection without a database.
	if io.IsConnected() {
		io.Disconnect()
	}

	// Reconnect using the newly created database.
	err = io.Connect(dbName)

	log.Printf("...end DatabaseCreate(%s)\n", util.ErrorString(err))

	return err
}

//----------------------------------------------------------------------------
//								Delete
//----------------------------------------------------------------------------

// DatabaseDelete deletes the table in the
// given database if present.
func (io *IO_App01ma) DatabaseDelete(dbName string) error {
	var err error
	var str util.StringBuilder

	log.Printf("DatabaseDelete()\n")

	dbName = strings.ToLower(dbName)

	// Build the Create Database SQL Statement.

	if !io.IsDatabaseDefined(dbName) {
		err = io.Exec(str.String())
	}

	log.Printf("...end DatabaseDelete(%s)\n", util.ErrorString(err))

	return err
}

//----------------------------------------------------------------------------
//						IsDatabaseDefined - mariadb
//----------------------------------------------------------------------------

// IsDatabaseDefined checks to see if the Database is already defined to the SQL server.
// This is not needed in SQLite. So, we just return true.
func (io *IO_App01ma) IsDatabaseDefined(dbName string) bool {
	var str util.StringBuilder
	var err error
	var row *sql.Row
	var Database string

	log.Printf("IsDatabaseDefined(%s)\n", dbName)

	dbName = strings.ToLower(dbName)

	// Build the SQL Statement.
	str.WriteStringf("SELECT schema_name FROM information_schema.schemata WHERE schema_name = '%s';", dbName)

	row = io.dbSql.QueryRow(str.String())
	err = row.Scan(&Database)
	if err == nil {
		if Database == dbName {

			log.Printf("...end IsDatabaseDefined(true)\n")

			return true
		}

	} else {
		log.Printf("\tSELECT schema_name Error: %s  Name: %s\n", err.Error(), Database)

	}

	log.Printf("...end IsDatabaseDefined(false)\n")

	return false
}

//----------------------------------------------------------------------------
//								    Exec
//----------------------------------------------------------------------------

// Exec executes an sql statement which does not return any rows.
func (io *IO_App01ma) Exec(sqlStmt string, args ...interface{}) error {
	var err error

	log.Printf("Exec(%s)\n", sqlStmt)

	_, err = io.dbSql.Exec(sqlStmt, args...)

	log.Printf("...end Exec(%s)\n", util.ErrorString(err))

	return err
}

//----------------------------------------------------------------------------
//								    Query
//----------------------------------------------------------------------------

// Query executes an sql statement which does return row(s).
func (io *IO_App01ma) Query(sqlStmt string, process func(rows *sql.Rows), args ...interface{}) error {
	var err error
	var rows *sql.Rows

	log.Printf("Query(%s)\n", sqlStmt)

	rows, err = io.dbSql.Query(sqlStmt, args...)

	if err == nil {
		defer rows.Close()
		// Process the rows
		for rows.Next() {
			process(rows)
		}
		err = rows.Close()
	}

	log.Printf("...end Query(%s)\n", util.ErrorString(err))

	return err
}

//----------------------------------------------------------------------------
//								    QueryRow
//----------------------------------------------------------------------------

// QueryRow executes an sql statement which does return row(s).
func (io *IO_App01ma) QueryRow(sqlStmt string, args ...interface{}) *sql.Row {
	var err error
	var row *sql.Row

	log.Printf("QueryRow(%s)\n", sqlStmt)

	row = io.dbSql.QueryRow(sqlStmt, args...)

	log.Printf("...end Query(%s)\n", util.ErrorString(err))

	return row
}

//----------------------------------------------------------------------------
//                                  NewIoApp01ma
//----------------------------------------------------------------------------

// New creates a new struct.
func NewIoApp01ma() *IO_App01ma {
	db := &IO_App01ma{}
	return db
}
