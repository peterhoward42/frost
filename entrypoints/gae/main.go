package frost

import (
	"appengine"
	"github.com/peterhoward42/frost/filereaders"
	"github.com/peterhoward42/frost/server/pages/view/models"
	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
)

func init() {
	var err error
	gui_template, err = template.ParseGlob("../../server/static/templates/*.html")
    	if err != nil {
        	panic(fmt.Sprintf("template parsing failed with %v", err.Error()))
	}

	// Todo, not sure yet if these two URLs should continue to use a handler in common.
	http.HandleFunc("/playground", playground)
	http.HandleFunc("/playground-refresh", playground)

	http.HandleFunc("/play-space-separated", playground_space_sep)
	http.HandleFunc("/play-csv", playground_csv)

	http.HandleFunc("/quickstart", quickstart)

	http.HandleFunc("/", quickstart) // landing page and catch-all route is to quickstart.
}

// todo, plenty of opportunity for consolidation and reuse once the full set of urls is
// clearer.
func quickstart(w http.ResponseWriter, r *http.Request) {
	// todo or panic?
	err := gui_template.ExecuteTemplate(w, "maingui.html",
		models.NewTopLevel("make quickstart active", "", ""))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playground(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	input_text := r.PostFormValue("input-text")
	output_text := input_text + "plus some extra stuff"
	err = gui_template.ExecuteTemplate(w, "maingui.html",
		models.NewTopLevel("make playground active", input_text, output_text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playground_space_sep(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("../../server/static/examples/space_delim.txt")
	input_text := string(buf)
	output_text := string(filereaders.NewWhitespace(
		input_text, appengine.NewContext(r)).Convert())
	err := gui_template.ExecuteTemplate(w, "maingui.html",
		models.NewTopLevel("make playground active", input_text, output_text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playground_csv(w http.ResponseWriter, r *http.Request) {

	buf, _ := ioutil.ReadFile("../../server/static/examples/csv.csv")
	input_text := string(buf)
	output_text := "CONVERTED TO JSON of this" + input_text
	err := gui_template.ExecuteTemplate(w, "maingui.html",
		models.NewTopLevel("make playground active", input_text, output_text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var gui_template *template.Template = nil
