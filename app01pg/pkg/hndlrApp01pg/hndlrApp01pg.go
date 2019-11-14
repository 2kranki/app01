// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Thu Nov 14, 2019 11:17


package hndlrApp01pg

import (
	"fmt"
	_ "io"
	_ "io/ioutil"
    "html/template"
    
	"net/http"
    _ "os"
    _ "sort"
    _ "strings"

    
	_ "github.com/lib/pq"
)

type TmplsApp01pg  struct {
    tmplsDir        string
    Tmpls           *template.Template
}


func (TmplsApp01pg) Title(i interface{}) string {
    return "Title() - NOT Implemented"
}

func (TmplsApp01pg) Body(i interface{}) string {
    return "Body() - NOT Implemented"
}

func (t *TmplsApp01pg) SetTmplsDir(d string) {
    t.tmplsDir = d
}

//----------------------------------------------------------------------------
//                             Main Display
//----------------------------------------------------------------------------

// Display the main menu with any needed messages.
func (h *TmplsApp01pg) MainDisplay(w http.ResponseWriter, msg string) {
    var err     error
    var name    = "App01pg.main.menu.gohtml"
    

    

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

func NewTmplsApp01pg(dir string) *TmplsApp01pg {
    t := &TmplsApp01pg{}
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
func (t *TmplsApp01pg) SetupTmpls() {
    

    funcs := map[string]interface{}{"Title":t.Title, "Body":t.Body,}
    path := t.tmplsDir + "/*.gohtml"
	t.Tmpls = template.Must(template.New("tmpls").Funcs(funcs).ParseGlob(path))
}

func init() {

}

