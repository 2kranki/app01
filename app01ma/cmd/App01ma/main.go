// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//              Application main program
// This module handles the CLI parameters, displaying help
// if needed. It then passes control to mainExec for the
// primary application processing.

// Notes:
//  1.  When working with package main, please keep in mind that the
//      more functionality that you can move into functions, the easier
//      testing will be. This allows you to test the functionality in
//      small portions. Moving common functionality to packages that are
//      easily tested is even better.
//  2.  If HTTPS is specified, we will default to looking for key.pem and
//      cert.pem in the certDir ("/tmp/cert" default). If the directory
//      or the required files are not present, we will generate temporary
//      versions of them in the specified directory using openssl with
//      default parameters.

// Generated: Mon Jan  6, 2020 09:54 for mariadb Database

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	debug     bool
	force     bool
	noop      bool
	quiet     bool
	db_name   string
	db_pw     string
	db_port   string
	db_srvr   string
	db_user   string
	http_srvr string
	http_port string
	baseDir   string
	execPath  string // exec json path (optional)

	certDir    string
	https_port string
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nNotes:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "'baseDir' is assumed to point to a directory where the application\n")
	fmt.Fprintf(flag.CommandLine.Output(), " can find 'html', 'css' and 'tmpl' sub-directories.\n\n")

	fmt.Fprintf(flag.CommandLine.Output(), "'exec json' is a file that defines the command line parameters \n")
	fmt.Fprintf(flag.CommandLine.Output(), "so that you can set them and then execute gen with -x or -exec\n")
	fmt.Fprintf(flag.CommandLine.Output(), "option.\n\n")

}

// parseFlags parses the command line flags. If there are any errors,
// it displays the usage help and exits.
func parseFlags() {

	// Set up flag variables
	log.Printf("\tSetting up the flag variables...\n")

	flag.Usage = usage
	flag.BoolVar(&debug, "debug", true, "enable debugging")
	flag.BoolVar(&force, "force", true, "enable over-writes and deletions")
	flag.BoolVar(&force, "f", true, "enable over-writes and deletions")
	flag.BoolVar(&noop, "noop", true, "execute program, but do not make real changes")
	flag.BoolVar(&quiet, "quiet", true, "enable quiet mode")
	flag.BoolVar(&quiet, "q", true, "enable quiet mode")
	flag.StringVar(&execPath, "exec", "", "exec json path (optional)")

	flag.StringVar(&db_pw, "dbPW", "Passw0rd", "the database password")
	flag.StringVar(&db_port, "dbPort", "4306", "the database port")
	flag.StringVar(&db_srvr, "dbServer", "localhost", "the database server")
	flag.StringVar(&db_user, "dbUser", "root", "the database user")
	flag.StringVar(&db_name, "dbName", "App01ma", "the database name")

	flag.StringVar(&http_port, "httpPort", "8090", "server port")
	flag.StringVar(&http_srvr, "httpServer", "localhost", "server site")
	flag.StringVar(&baseDir, "basedir", ".", "Base Directory for Templates, HTML and CSS")
	flag.StringVar(&certDir, "certdir", "/tmp/certs", "Base Directory for HTTPS Certificates")
	flag.StringVar(&https_port, "httpsPort", "8095", "HTTPS server port")

	// Parse the flags and check them
	log.Printf("\tParsing the flags...\n")
	flag.Parse()
	if debug {
		log.Println("\tIn Debug Mode...")
	}

}

// envOverride looks for certain environment variables and if found
// overrides the flags that they speciffy.
func envOverride() {
	var wrk string

	// Collect variables from Environment and override value if present.
	log.Printf("\tCollecting the variables from Environment and override value if present...\n")
	wrk = os.Getenv("APP01MA_HTTP_PORT")
	if len(wrk) > 0 {
		http_port = wrk
	}
	wrk = os.Getenv("APP01MA_HTTP_SERVER")
	if len(wrk) > 0 {
		http_srvr = wrk
	}
	wrk = os.Getenv("APP01MA_BASEDIR")
	if len(wrk) > 0 {
		baseDir = wrk
	}
	wrk = os.Getenv("APP01MA_EXEC")
	if len(wrk) > 0 {
		execPath = wrk
	}

	wrk = os.Getenv("APP01MA_DB_PW")
	if len(wrk) > 0 {
		db_pw = wrk
	}
	wrk = os.Getenv("APP01MA_DB_PORT")
	if len(wrk) > 0 {
		db_port = wrk
	}
	wrk = os.Getenv("APP01MA_DB_SERVER")
	if len(wrk) > 0 {
		db_srvr = wrk
	}
	wrk = os.Getenv("APP01MA_DB_USER")
	if len(wrk) > 0 {
		db_user = wrk
	}
	wrk = os.Getenv("APP01MA_DB_NAME")
	if len(wrk) > 0 {
		db_name = wrk
	}

}

// main is the main entry point of the application. It parses the
// CLI flags, overrides any flags specified by Environment variables
// and executes the the main application logic.
func main() {

	parseFlags()
	envOverride()

	// Execute the main process.
	log.Printf("\tExecuting the main process...\n")
	mainExec()
}
