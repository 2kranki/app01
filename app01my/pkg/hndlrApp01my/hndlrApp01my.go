// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Notes:
//  *   All static (ie non-changing) files should be served from the 'static'
//      subdirectory.

// Generated: Mon Oct 28, 2019 08:40


package hndlrApp01my

import (
	"fmt"
	"io"
	_ "io/ioutil"
    "html/template"
    "log"
	"net/http"
    _ "os"
    "sort"
    "strings"

    "github.com/2kranki/go_util"
	_ "github.com/go-sql-driver/mysql"
)

type TmplsApp01my  struct {
    tmplsDir        string
    Tmpls           *template.Template
}


func (TmplsApp01my) Title(i interface{}) string {
    return "Title() - NOT Implemented"
}

func (TmplsApp01my) Body(i interface{}) string {
    return "Body() - NOT Implemented"
}

func (t *TmplsApp01my) SetTmplsDir(d string) {
    t.tmplsDir = d
}

//----------------------------------------------------------------------------
//                             Main Display
//----------------------------------------------------------------------------

// Display the main menu with any needed messages.
func (h *TmplsApp01my) MainDisplay(w http.ResponseWriter, msg string) {
    var err     error
    var name    = "App01my.main.menu.gohtml"
    
        var str     strings.Builder
    

    
        log.Printf("App01my.MainDisplay(%s)\n", msg)
        log.Printf("\tname: %s\n", name)
        w2 := io.MultiWriter(w, &str)
    

    data := struct {
                Msg         string
            }{msg}

    
        log.Printf("\tData: %+v\n", data)
    

    log.Printf("\tExecuting template: %s\n", name)
        err = h.Tmpls.ExecuteTemplate(w2, name, data)
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    
        log.Printf("\t output: %s\n", str.String())
        log.Printf("...end App01my.MainDisplay(%s)\n", util.ErrorString(err))
    
}

//----------------------------------------------------------------------------
//                                  N e w
//----------------------------------------------------------------------------

func NewTmplsApp01my(dir string) *TmplsApp01my {
    t := &TmplsApp01my{}
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
func (t *TmplsApp01my) SetupTmpls() {
    
        var templates   []*template.Template
        var tt          *template.Template
        var names       []string
        var name        string
    
        log.Printf("\tSetupTmpls(%s/*.gohtml)\n", t.tmplsDir)

    funcs := map[string]interface{}{"Title":t.Title, "Body":t.Body,}
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

