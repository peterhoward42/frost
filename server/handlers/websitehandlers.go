package handlers

import (
	"fmt"
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/render"
	"github.com/peterhoward42/frost/server/viewmodels"
	"net/http"
	"strings"
)

func HandleQuickStart(w http.ResponseWriter, r *http.Request) {
	viewModel := &viewmodels.TopLevelViewModel{}
	viewModel.QuickStart = &viewmodels.QuickStartViewModel{}
	renderer := view.NewGuiRenderer(resources.CompiledTemplates)
	renderer.Render(w, viewModel)
}

func HandlePlayground(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(fmt.Sprintf("Cannot parse form. Error is: %v", err))
	}
	viewModel := &viewmodels.TopLevelViewModel{}
	switch {
	case strings.Contains(r.URL.Path, "refresh"):
		viewModel.Playground = viewmodels.NewPlaygroundViewModelForRefresh(
			r.Form, r.URL.Path)
	case strings.Contains(r.URL.Path, "example"):
		exampleInputText := string(resources.MustAsset(`staticfiles/examples/space_delim
		.txt`))
		viewModel.Playground = viewmodels.NewPlaygroundViewModelForExample(exampleInputText)
	}
	renderer := view.NewGuiRenderer(resources.CompiledTemplates)
	renderer.Render(w, viewModel)
}
