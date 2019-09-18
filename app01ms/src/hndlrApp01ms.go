// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Generated: Wed Sep 18, 2019 11:02

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

	_ "github.com/denisenkom/go-mssqldb"
)

type TmplsApp01ms struct {
	tmplsDir string
	Tmpls    *template.Template
}

func (TmplsApp01ms) Title(i interface{}) string {
	return "Title() - NOT Implemented"
}

func (TmplsApp01ms) Body(i interface{}) string {
	return "Body() - NOT Implemented"
}

func init() {

}

func (t *TmplsApp01ms) SetTmplsDir(d string) {
	t.tmplsDir = d
}

func NewTmplsApp01ms() *TmplsApp01ms {
	t := &TmplsApp01ms{}
	t.tmplsDir = "./tmpl"
	return t
}

//----------------------------------------------------------------------------
//                           Setup Templates
//----------------------------------------------------------------------------

// SetupTmpls initializes the functions used in the templates
// and loads them.
func (t *TmplsApp01ms) SetupTmpls() {

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
