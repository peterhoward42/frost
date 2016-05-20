package routing

import (
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/handlers"
	"github.com/peterhoward42/frost/server/urls"
	"net/http"
)

func SetUpRouting() {

	http.Handle("/static/", http.FileServer(resources.CompiledFileSystem))

	http.HandleFunc(urls.URLHomePage, handlers.HandleQuickStart)

	http.HandleFunc(urls.URLPlaygroundStub, handlers.HandlePlayground)
	http.HandleFunc(urls.URLQuickstart, handlers.HandleQuickStart)

}
