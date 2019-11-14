// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Thu Nov 14, 2019 11:17


package hndlrApp01sq

import (
	"fmt"
	_ "io"
	_ "io/ioutil"
    "html/template"
    
	"net/http"
    _ "os"
    _ "sort"
    _ "strings"

    
	_ "github.com/mattn/go-sqlite3"
)

type TmplsApp01sq  struct {
    tmplsDir        string
    Tmpls           *template.Template
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
    var err     error
    var name    = "App01sq.main.menu.gohtml"
    

    

    data := struct {
                Msg         string
            }{msg}

    

    
        err = h.Tmpls.ExecuteTemplate(w, name, data)
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    
}

//----------------------------------------------------------------------------
//                                  N e w
//----------------------------------------------------------------------------

func NewTmplsApp01sq(dir string) *TmplsApp01sq {
    t := &TmplsApp01sq{}
    if "" == dir {
        t.tmplsDir = "./tmpl"
    }
    return t
}

//----------------------------------------------------------------------------
//                           Setup Templates
//----------------------------------------------------------------------------

// SetupTmpls initializes the functions used in the templates
// and loads them.
func (t *TmplsApp01sq) SetupTmpls() {
    

    funcs := map[string]interface{}{"Title":t.Title, "Body":t.Body,}
    path := t.tmplsDir + "/*.gohtml"
	t.Tmpls = template.Must(template.New("tmpls").Funcs(funcs).ParseGlob(path))
}

func init() {

}

