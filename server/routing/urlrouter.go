package routing

import (
	"net/http"
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/handlers"
)

const URLPlayground = `/playground`
const URLQuickstart = `/quickstart`

func SetUpRouting() {

	http.Handle("/static/", http.FileServer(resources.CompiledFileSystem))

	http.HandleFunc("/", handlers.HandleQuickStart) // landing page and catch-all route is to
	// quickstart.

	http.HandleFunc(URLPlayground, handlers.HandlePlayground)
	http.HandleFunc(URLQuickstart, handlers.HandleQuickStart)

}
