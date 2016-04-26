package frost

import (
	"net/http"
	"html/template"
	"server/pages/view/models"
	"io/ioutil"
)

func init() {
	gui_template, _ = template.ParseGlob("server/static/templates/*.html")
	// todo what if template read fails during init?

	// Todo, not sure yet if these two URLs should continue to use a handler in common.
	http.HandleFunc("/playground", playground)
	http.HandleFunc("/playground-refresh", playground)

	http.HandleFunc("/play-space-separated", playground_space_sep)
	http.HandleFunc("/play-csv", playground_csv)

	http.HandleFunc("/quickstart", quickstart)

	http.HandleFunc("/home", quickstart) // We use quickstart for the landing page.
}

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

	buf, _ := ioutil.ReadFile("server/static/examples/space_delim.txt")
	input_text := string(buf)
	output_text := input_text + "...IS CONVERTED"
	err := gui_template.ExecuteTemplate(w, "maingui.html",
		models.NewTopLevel("make playground active", input_text, output_text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playground_csv(w http.ResponseWriter, r *http.Request) {

	buf, _ := ioutil.ReadFile("server/static/examples/csv.csv")
	input_text := string(buf)
	output_text := input_text + "...IS CONVERTED"
	err := gui_template.ExecuteTemplate(w, "maingui.html",
		models.NewTopLevel("make playground active", input_text, output_text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var gui_template *template.Template = nil
