// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Generated: Tue Sep 24, 2019 21:07

package main

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

func (t *TmplsApp01ms) SetTmplsDir(d string) {
	t.tmplsDir = d
}

func NewTmplsApp01ms() *TmplsApp01ms {
	t := &TmplsApp01ms{}
	t.tmplsDir = "./tmpl"
	return t
}

//----------------------------------------------------------------------------
//                             Main Display
//----------------------------------------------------------------------------

// Display the main menu with any needed messages.
func (h *TmplsApp01ms) MainDisplay(w http.ResponseWriter, msg string) {
	var err error
	var name = "App01ms.main.menu.gohtml"

	var str strings.Builder

	log.Printf("App01ms.MainDisplay(%s)\n", msg)
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
	log.Printf("...end App01ms.MainDisplay(%s)\n", util.ErrorString(err))

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

func init() {

}
