// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Mon Jan  6, 2020 11:09

package hndlrApp01sq

import (
	"fmt"
	"html/template"
	"io"
	_ "io/ioutil"
	"log"
	"net/http"
	_ "os"
	"sort"
	"strings"

	"github.com/2kranki/go_util"
	_ "github.com/mattn/go-sqlite3"
)

//----------------------------------------------------------------------------
//                     App01sq Templates
//----------------------------------------------------------------------------

type TmplsApp01sq struct {
	tmplsDir string
	Tmpls    *template.Template
}

func (TmplsApp01sq) Title(i interface{}) string {
	return "Title() - NOT Implemented"
}

func (TmplsApp01sq) Body(i interface{}) string {
	return "Body() - NOT Implemented"
}

func (t *TmplsApp01sq) SetTmplsDir(d string) {
	t.tmplsDir = d
}

//----------------------------------------------------------------------------
//                             Main Display
//----------------------------------------------------------------------------

// Display the main menu with any needed messages.
func (h *TmplsApp01sq) MainDisplay(w http.ResponseWriter, msg string) {
	var err error
	var name = "App01sq.main.menu.gohtml"

	var str strings.Builder

	log.Printf("App01sq.MainDisplay(%s)\n", msg)
	log.Printf("\tname: %s\n", name)
	w2 := io.MultiWriter(w, &str)

	data := struct {
		Msg string
	}{msg}

	log.Printf("\tData: %+v\n", data)

	log.Printf("\tExecuting template: %s\n", name)
	err = h.Tmpls.ExecuteTemplate(w2, name, data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	log.Printf("\t output: %s\n", str.String())
	log.Printf("...end App01sq.MainDisplay(%s)\n", util.ErrorString(err))

}

//----------------------------------------------------------------------------
//                           Setup Templates
//----------------------------------------------------------------------------

// SetupTmpls initializes the functions used in the templates
// and loads them.
func (t *TmplsApp01sq) SetupTmpls() {

	var templates []*template.Template
	var tt *template.Template
	var names []string
	var name string

	log.Printf("\tSetupTmpls(%s/*.gohtml)\n", t.tmplsDir)

	funcs := map[string]interface{}{"Title": t.Title, "Body": t.Body}
	path := t.tmplsDir + "/*.gohtml"
	t.Tmpls = template.Must(template.New("tmpls").Funcs(funcs).ParseGlob(path))
	templates = t.Tmpls.Templates()
	for _, tt = range templates {
		names = append(names, tt.Name())
	}
	sort.Strings(names)
	for _, name = range names {
		log.Printf("\t\t template: %s\n", name)
	}
	log.Printf("\tend of SetupTmpls()\n")
}

//----------------------------------------------------------------------------
//                                  N e w
//----------------------------------------------------------------------------

func NewTmplsApp01sq(dir string) *TmplsApp01sq {
	t := &TmplsApp01sq{}
	if dir == "" {
		t.tmplsDir = "./tmpl"
	} else {
		t.tmplsDir = dir
	}
	return t
}

func init() {

}
