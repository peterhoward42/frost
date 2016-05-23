package handlers

import (
	"fmt"
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/render"
	"github.com/peterhoward42/frost/server/viewmodels"
	"net/http"
	"strings"
	"github.com/peterhoward42/frost/server/urls"
)

func HandleQuickStart(w http.ResponseWriter, r *http.Request) {
	viewModel := viewmodels.NewTopLevelViewModel()
	viewModel.QuickStart = viewmodels.NewQuickStartViewModel()
	renderer := view.NewGuiRenderer(resources.CompiledTemplates)
	renderer.Render(w, viewModel)
}

// The HandlePlayground function is able to handle all the URL requests stimulated by the
// playground page. The playground page houses an HTML form, so that we can harvest the text
// from the input area and show the converted text in the frost area. The page includes two types
// of requests. Firstly the edit-refresh cycle. Secondly direct mandates to display a canned
// example file conversion. Additionally the family of refresh URL's include parameters to that
// state a preference for side by side, or tabbed viewing modes.
func HandlePlayground(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(fmt.Sprintf("Cannot parse form. Error is: %v", err))
	}
	viewModel := viewmodels.NewTopLevelViewModel()
	switch {
	// There are a set of playground refresh URLs, but they all share the root we
	// are looking for below. The tail end of the URL that comes afterwards encodes some
	// viewing modes - like tabbed or side by side.
	case strings.Contains(r.URL.Path, urls.URLPlaygroundRefreshStub):
		viewModel.Playground = viewmodels.NewPlaygroundViewModelForRefresh(
			r.Form, r.URL.Path)
	case strings.Contains(r.URL.Path, urls.URLPlaygroundExampleSpaceDelim):
		exampleInputText :=resources.GetExampleFileContents(resources.SpaceDelimitedExample)
		spaceSeparatedButtonActiveString := "active"
		viewModel.Playground = viewmodels.NewPlaygroundViewModelForExample(
			exampleInputText, spaceSeparatedButtonActiveString)
	default:
		panic(fmt.Sprintf("URL that reached HandlePlayground contains neither refresh"+
			"nor example variant hint: %v", r.URL.Path))
	}
	renderer := view.NewGuiRenderer(resources.CompiledTemplates)
	renderer.Render(w, viewModel)
}

func HandleComingSoon(w http.ResponseWriter, r *http.Request) {
	viewModel := viewmodels.NewTopLevelViewModel()
	viewModel.ComingSoon = &viewmodels.ComingSoonViewModel{}
	renderer := view.NewGuiRenderer(resources.CompiledTemplates)
	renderer.Render(w, viewModel)
}
