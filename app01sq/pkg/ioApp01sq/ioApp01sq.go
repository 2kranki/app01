// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01sq contains all the methods for manipulating an SQL
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


// Generated: Thu Nov 14, 2019 11:17 for sqlite Database

package ioApp01sq

import (
    "database/sql"
    "fmt"
    
    
     
    "strings"
	// "time" is only needed for Docker support and "sqlite" is the only
	//  database server not using it.
    

    "github.com/2kranki/go_util"
    
	    _ "github.com/mattn/go-sqlite3"
)






//============================================================================
//                            IO_App01sq
//============================================================================

type IO_App01sq struct {
    dbSql       *sql.DB
    dbName      string
    dbPW       	string
    dbPort     	string
    dbServer    string
    dbUser     	string
}

func (io *IO_App01sq) Name() string {
    return io.dbName
}
func (io *IO_App01sq) SetName(str string) {
    io.dbName = str
}

func (io *IO_App01sq) PW() string {
    return io.dbPW
}
func (io *IO_App01sq) SetPW(str string) {
    io.dbPW = str
}

func (io *IO_App01sq) Port() string {
    return io.dbPort
}
func (io *IO_App01sq) SetPort(str string) {
    io.dbPort = str
}

func (io *IO_App01sq) Server() string {
    return io.dbServer
}
func (io *IO_App01sq) SetServer(str string) {
    io.dbServer = str
}

func (io *IO_App01sq) Sql() *sql.DB {
    return io.dbSql
}

func (io *IO_App01sq) User() string {
    return io.dbUser
}
func (io *IO_App01sq) SetUser(str string) {
    io.dbUser = str
}

//============================================================================
//                              Miscellaneous
//============================================================================



// Set up default parameters for the needed SQL Type.
func (io *IO_App01sq) DefaultParms() {
		io.SetPort("")
		io.SetPW("")
		io.SetServer("")
		io.SetUser("")
		//io.SetName("App01sq.db")
}

//============================================================================
//                              Database Methods
//============================================================================

//----------------------------------------------------------------------------
//								Connect - SQLite
//----------------------------------------------------------------------------

// Connect connects the database/sql/driver to the appropriate
// database server using the given parameters.
func (io *IO_App01sq) Connect(dbName string) error {
    var err         error

    
    io.dbSql, err = sql.Open("sqlite3", dbName)
    if err != nil {
        return fmt.Errorf("Error: Cannot Connect: %s\n", err.Error())
    }

    
    err = io.dbSql.Ping()
    if err != nil {
        io.Disconnect( )
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
func (io *IO_App01sq) Disconnect() error {
    var err         error

    
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

func (io *IO_App01sq) IsConnected() bool {
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
func (io *IO_App01sq) DatabaseCreate(dbName string) error {
    var err     error
    

    
    if len(dbName) == 0 {
        return fmt.Errorf("Error: Missing database name for DatabaseCreate()!")
    }

    

    // Reconnect using the newly created database.
    err = io.Connect(dbName)

    
    return err
}

//----------------------------------------------------------------------------
//								Delete
//----------------------------------------------------------------------------

// DatabaseDelete deletes the table in the
// given database if present.
func (io *IO_App01sq) DatabaseDelete(dbName string) error {
    var err     error
	var str		util.StringBuilder

    
    dbName = strings.ToLower(dbName)

	// Build the Create Database SQL Statement.
    	str.WriteStringf("DROP DATABASE IF EXISTS %s;", dbName)


    

    
    return err
}

//----------------------------------------------------------------------------
//						IsDatabaseDefined - sqlite
//----------------------------------------------------------------------------

// IsDatabaseDefined checks to see if the Database is already defined to the SQL server.
// This is not needed in SQLite. So, we just return true.
func (io *IO_App01sq) IsDatabaseDefined() bool {

    

    
    return true
}

//----------------------------------------------------------------------------
//								    Exec
//----------------------------------------------------------------------------

// Exec executes an sql statement which does not return any rows.
func (io *IO_App01sq) Exec(sqlStmt string, args ...interface{}) error {
    var err     error

    

    _, err = io.dbSql.Exec(sqlStmt, args...)
    

    
    return err
}

//----------------------------------------------------------------------------
//								    Query
//----------------------------------------------------------------------------

// Query executes an sql statement which does return row(s).
func (io *IO_App01sq) Query(sqlStmt string, process func(rows *sql.Rows), args ...interface{}) error {
    var err     error
    var rows    *sql.Rows

    

    rows, err = io.dbSql.Query(sqlStmt, args...)
    
    if err == nil {
        defer rows.Close()
        // Process the rows
        for rows.Next() {
            process(rows)
        }
        err = rows.Close()
    }

    
    return err
}

//----------------------------------------------------------------------------
//								    QueryRow
//----------------------------------------------------------------------------

// QueryRow executes an sql statement which does return row(s).
func (io *IO_App01sq) QueryRow(sqlStmt string, args ...interface{}) *sql.Row {
    
    var row     *sql.Row

    

    row = io.dbSql.QueryRow(sqlStmt, args...)

    

    
    return row
}




//----------------------------------------------------------------------------
//                                  NewIoApp01sq
//----------------------------------------------------------------------------

// New creates a new struct.
func NewIoApp01sq() *IO_App01sq {
    db := &IO_App01sq{}
    return db
}

