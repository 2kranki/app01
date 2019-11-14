// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Thu Nov 14, 2019 11:17


package hndlrApp01ma

import (
	"fmt"
	_ "io"
	_ "io/ioutil"
    "html/template"
    
	"net/http"
    _ "os"
    _ "sort"
    _ "strings"

    
	_ "github.com/go-sql-driver/mysql"
)

type TmplsApp01ma  struct {
    tmplsDir        string
    Tmpls           *template.Template
}


func (TmplsApp01ma) Title(i interface{}) string {
    return "Title() - NOT Implemented"
}

func (TmplsApp01ma) Body(i interface{}) string {
    return "Body() - NOT Implemented"
}

func (t *TmplsApp01ma) SetTmplsDir(d string) {
    t.tmplsDir = d
}

//----------------------------------------------------------------------------
//                             Main Display
//----------------------------------------------------------------------------

// Display the main menu with any needed messages.
func (h *TmplsApp01ma) MainDisplay(w http.ResponseWriter, msg string) {
    var err     error
    var name    = "App01ma.main.menu.gohtml"
    

    

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

func NewTmplsApp01ma(dir string) *TmplsApp01ma {
    t := &TmplsApp01ma{}
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
func (t *TmplsApp01ma) SetupTmpls() {
    

    funcs := map[string]interface{}{"Title":t.Title, "Body":t.Body,}
    path := t.tmplsDir + "/*.gohtml"
	t.Tmpls = template.Must(template.New("tmpls").Funcs(funcs).ParseGlob(path))
}

func init() {

}

