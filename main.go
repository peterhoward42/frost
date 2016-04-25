package frost

import (
	"net/http"
	"html/template"
	"server/pages/view/models"
)

func init() {
	gui_template, _ = template.ParseGlob("server/static/templates/*.html")
	// todo what if template read fails during init?

	// Todo, not sure yet if these two URLs should continue to use a handler in common.
	http.HandleFunc("/playground", playground)
	http.HandleFunc("/playground-refresh", playground)
	http.HandleFunc("/play-space-separated", playground_space_sep)

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

	input_text := "this will be the space separated input content"
	output_text := input_text + "...IS CONVERTED"
	err := gui_template.ExecuteTemplate(w, "maingui.html",
		models.NewTopLevel("make playground active", input_text, output_text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var gui_template *template.Template = nil
