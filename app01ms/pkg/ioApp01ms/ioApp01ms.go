// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ms contains all the methods for manipulating an SQL
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


// Generated: Fri Oct 25, 2019 11:40 for mssql Database

package ioApp01ms

import (
    "database/sql"
    "fmt"
    
        "log"
    
    "net/url"
    
     "strconv"
    "strings"
	// "time" is only needed for Docker support and "sqlite" is the only
	//  database server not using it.
    "time"
    

    "github.com/2kranki/go_util"
    
	    _ "github.com/denisenkom/go-mssqldb"
)

const connect_retries=100


type ErrorWithExtraInfo interface {
	SQLErrorLineNo() int32
	SQLErrorNumber() int32
}


//============================================================================
//                            IO_App01ms
//============================================================================

type IO_App01ms struct {
    dbSql       *sql.DB
    dbName      string
    dbPW       	string
    dbPort     	string
    dbServer    string
    dbUser     	string
}

func (io *IO_App01ms) Name() string {
    return io.dbName
}
func (io *IO_App01ms) SetName(str string) {
    io.dbName = str
}

func (io *IO_App01ms) PW() string {
    return io.dbPW
}
func (io *IO_App01ms) SetPW(str string) {
    io.dbPW = str
}

func (io *IO_App01ms) Port() string {
    return io.dbPort
}
func (io *IO_App01ms) SetPort(str string) {
    io.dbPort = str
}

func (io *IO_App01ms) Server() string {
    return io.dbServer
}
func (io *IO_App01ms) SetServer(str string) {
    io.dbServer = str
}

func (io *IO_App01ms) Sql() *sql.DB {
    return io.dbSql
}

func (io *IO_App01ms) User() string {
    return io.dbUser
}
func (io *IO_App01ms) SetUser(str string) {
    io.dbUser = str
}

//============================================================================
//                              Miscellaneous
//============================================================================



    func (io *IO_App01ms) FloatToString(num float64) string {
        s := fmt.Sprintf("%.4f", num)
        return strings.TrimRight(strings.TrimRight(s, "0"), ".")
    }

    func (io *IO_App01ms) StringToFloat(str string) float64 {
        var num float64
        num, _ = strconv.ParseFloat(str, 64)
        return num
    }



// Set up default parameters for the needed SQL Type.
func (io *IO_App01ms) DefaultParms() {
		io.SetPort("1401")
		io.SetPW("Passw0rd")
		io.SetServer("localhost")
		io.SetUser("sa")
		//io.SetName("App01ms")
}

//============================================================================
//                              Database Methods
//============================================================================

//----------------------------------------------------------------------------
//								Connect - MS SQL
//----------------------------------------------------------------------------

// Connect connects the database/sql/driver to the appropriate
// database server using the given parameters.
func (io *IO_App01ms) Connect(dbName string) error {
    var err         error

    dbName = strings.ToLower(dbName)

    // Set up connection string, connStr.
	query := url.Values{}
	query.Add("database", dbName)
	query.Add("connection+timeout", "30")
	u := &url.URL{
		Scheme:		"sqlserver",
		User:		url.UserPassword(io.dbUser, io.dbPW),
		Host:		fmt.Sprintf("%s:%s", io.dbServer, io.dbPort),
		RawQuery:	query.Encode(),
	}
	connStr := u.String()

    // Allow for the Docker Container to get operational.
    for i:=0; i<connect_retries; i++ {
        log.Printf("\tConnecting %d to mssql with %s...\n", i, connStr)
        io.dbSql, err = sql.Open("mssql", connStr)
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
func (io *IO_App01ms) Disconnect() error {
    var err         error

    
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

func (io *IO_App01ms) IsConnected() bool {
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
func (io *IO_App01ms) DatabaseCreate(dbName string) error {
    var err     error
    var str		util.StringBuilder

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
        	str.WriteStringf("CREATE DATABASE %s;", dbName)


        // Create the database.
        err = io.Exec(str.String())
        if err != nil {
            io.Disconnect()
            return err
        }
        time.Sleep(5 * time.Second)         // Give it time to get done.
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
func (io *IO_App01ms) DatabaseDelete(dbName string) error {
    var err     error
	var str		util.StringBuilder

    
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
//						IsDatabaseDefined - mssql
//----------------------------------------------------------------------------

// IsDatabaseDefined checks to see if the Database is already defined to the
// SQL server.
func (io *IO_App01ms) IsDatabaseDefined(dbName string) bool {
	var str	    util.StringBuilder
    var err     error
    var row     *sql.Row
    var name    string

    
        log.Printf("IsDatabaseDefined(%s)\n", dbName)
    
    dbName = strings.ToLower(dbName)

    // Build the SQL Statement.
    str.WriteStringf("SELECT name FROM sys.databases WHERE name = N'%s';", dbName)

    row = io.dbSql.QueryRow(str.String())
    err = row.Scan(&name)
	if err == nil {
        if name == dbName {
        
            log.Printf("...end IsDatabaseDefined(true)\n")
        
            return true;
        }
    
	} else {
	        log.Printf("\tSELECT Error: %s  Name: %s\n", err.Error(), name)
    
	}

    
        log.Printf("...end IsDatabaseDefined(false)\n")
    
    return false
}

//----------------------------------------------------------------------------
//								ErrChk - mssql
//----------------------------------------------------------------------------

// ErrChk updates errors from mssql with other information provided.
func (io *IO_App01ms) ErrChk(err error) error {

    
        log.Printf("ErrChk(%s)\n", util.ErrorString(err))
    

    if err != nil {
		extra, ok := err.(ErrorWithExtraInfo)
		if ok {
			errNo  := int(extra.SQLErrorNumber())
			lineNo := int(extra.SQLErrorLineNo())
			err = fmt.Errorf("Error: %d  Line: %d - %s\n", errNo, lineNo, err.Error())
		}
    }

    
        log.Printf("...end ErrChk(%s)\n", util.ErrorString(err))
    
    return err
}



//----------------------------------------------------------------------------
//								    Exec
//----------------------------------------------------------------------------

// Exec executes an sql statement which does not return any rows.
func (io *IO_App01ms) Exec(sqlStmt string, args ...interface{}) error {
    var err     error

    
        log.Printf("Exec(%s)\n", sqlStmt)
    

    _, err = io.dbSql.Exec(sqlStmt, args...)
    err = io.ErrChk(err)

    
        log.Printf("...end Exec(%s)\n", util.ErrorString(err))
    
    return err
}

//----------------------------------------------------------------------------
//								    Query
//----------------------------------------------------------------------------

// Query executes an sql statement which does return row(s).
func (io *IO_App01ms) Query(sqlStmt string, process func(rows *sql.Rows), args ...interface{}) error {
    var err     error
    var rows    *sql.Rows

    
        log.Printf("Query(%s)\n", sqlStmt)
    

    rows, err = io.dbSql.Query(sqlStmt, args...)
    err = io.ErrChk(err)
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
func (io *IO_App01ms) QueryRow(sqlStmt string, args ...interface{}) *sql.Row {
    var err     error
    var row     *sql.Row

    
        log.Printf("QueryRow(%s)\n", sqlStmt)
    

    row = io.dbSql.QueryRow(sqlStmt, args...)

    err = io.ErrChk(err)

    
        log.Printf("...end Query(%s)\n", util.ErrorString(err))
    
    return row
}




//----------------------------------------------------------------------------
//                                  NewIoApp01ms
//----------------------------------------------------------------------------

// New creates a new struct.
func NewIoApp01ms() *IO_App01ms {
    db := &IO_App01ms{}
    return db
}

