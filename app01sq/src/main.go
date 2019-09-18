// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Generated: Wed Sep 18, 2019 11:02 for sqlite Database

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

func main() {
	var wrk string

	// Set up flag variables

	flag.Usage = usage
	flag.BoolVar(&debug, "debug", true, "enable debugging")
	flag.BoolVar(&force, "force", true, "enable over-writes and deletions")
	flag.BoolVar(&force, "f", true, "enable over-writes and deletions")
	flag.BoolVar(&noop, "noop", true, "execute program, but do not make real changes")
	flag.BoolVar(&quiet, "quiet", true, "enable quiet mode")
	flag.BoolVar(&quiet, "q", true, "enable quiet mode")
	flag.StringVar(&execPath, "exec", "", "exec json path (optional)")

	flag.StringVar(&db_name, "dbName", "App01sq.db", "the database path")

	flag.StringVar(&http_port, "httpPort", "8090", "server port")
	flag.StringVar(&http_srvr, "httpServer", "localhost", "server site")
	flag.StringVar(&baseDir, "basedir", ".", "Base Directory for Templates, HTML and CSS")

	// Parse the flags and check them
	flag.Parse()
	if debug {
		log.Println("\tIn Debug Mode...")
	}

	// Collect variables from Environment and override value if present.
	wrk = os.Getenv("APP01SQ_HTTP_PORT")
	if len(wrk) > 0 {
		http_port = wrk
	}
	wrk = os.Getenv("APP01SQ_HTTP_SERVER")
	if len(wrk) > 0 {
		http_srvr = wrk
	}
	wrk = os.Getenv("APP01SQ_BASEDIR")
	if len(wrk) > 0 {
		baseDir = wrk
	}
	wrk = os.Getenv("APP01SQ_EXEC")
	if len(wrk) > 0 {
		execPath = wrk
	}

	wrk = os.Getenv("APP01SQ_DB_NAME")
	if len(wrk) > 0 {
		db_name = wrk
	}

	// Execute the main process.
	exec()
}
