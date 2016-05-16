package main

import (
	"github.com/peterhoward42/frost/resources"
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(resources.VirtualFileSystem))
	http.HandleFunc("/thegui", gui_home_page_handler)

	http.ListenAndServe(":47066", nil)
}
