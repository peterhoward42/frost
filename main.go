package frost

import (
	"net/http"
	"html/template"
	"server/pages/view"
)

func init() {
	gui_template, _ = template.ParseFiles("server/static/templates/maingui.html")
	// todo what if template read fails during init?

	http.HandleFunc("/playground/", playground_handler)
	http.HandleFunc("/quickstart/", quickstart_handler)

	http.HandleFunc("/", quickstart_handler) // default landing page
}

func quickstart_handler(w http.ResponseWriter, r *http.Request) {
	err := gui_template.Execute(w, view.NewModel("make quickstart active"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func playground_handler(w http.ResponseWriter, r *http.Request) {
	err := gui_template.Execute(w, view.NewModel("make playground active"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var gui_template *template.Template = nil
