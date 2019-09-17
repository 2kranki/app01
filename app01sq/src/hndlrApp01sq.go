// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Generated: Tue Sep 17, 2019 10:59

package main

import (
	_ "fmt"
	"html/template"
	_ "io"
	_ "io/ioutil"
	_ "net/http"

	"log"

	_ "os"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

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

func init() {

}

func (t *TmplsApp01sq) SetTmplsDir(d string) {
	t.tmplsDir = d
}

func NewTmplsApp01sq() *TmplsApp01sq {
	t := &TmplsApp01sq{}
	t.tmplsDir = "./tmpl"
	return t
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
