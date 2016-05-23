package routing

import (
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/handlers"
	"github.com/peterhoward42/frost/server/urls"
	"net/http"
)

func SetUpRouting() {

	// Special cases first...

	// Static links like css, fonts and .js go to a psuedo static file server
	http.Handle("/static/", http.FileServer(resources.CompiledFileSystem))

	// Home page is a special case
	http.HandleFunc(urls.URLHomePage, handlers.HandleQuickStart)

	// The general case MUX
	http.HandleFunc(urls.URLPlaygroundStub, handlers.HandlePlayground)
	http.HandleFunc(urls.URLQuickstart, handlers.HandleQuickStart)

	// The catch-all case means we haven't (yet) put anything on the end of the incoming
	// URL
	http.HandleFunc("/", handlers.HandleComingSoon)

}
