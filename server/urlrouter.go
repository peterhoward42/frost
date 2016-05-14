package server

import "net/http"

func SetUpRouting() {

	http.HandleFunc("/playground", playground)
	http.HandleFunc("/quickstart", quickstart)
	http.HandleFunc("/", quickstart) // landing page and catch-all route is to quickstart.
}
