package server

import (
	"fmt"
	"github.com/peterhoward42/frost/server/pages/view"
	"github.com/peterhoward42/frost/server/pages/view/viewmodels"
	"net/http"
	"strings"
)

func HandleQuickStart(w http.ResponseWriter, r *http.Request) {
	viewModel := &viewmodels.TopLevel{}
	viewModel.QuickStart = viewmodels.QuickStart{}
	view.Render(w, viewModel)
}

func HandlePlayground(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(fmt.Sprintf("Cannot parse form. Error is: %v", err))
	}
	var viewModel *viewmodels.TopLevel
	submittedForm := r.Form
	switch {
	case strings.Contains(r.URL.Path, "refresh"):
		viewModel = viewmodels.NewPlaygroundViewModelForRefresh(
			w, submittedForm, r.URL.Path)
	case strings.Contains(r.URL.Path, "example"):
		viewModel = viewmodels.NewPlaygroundViewModelForExamples(w)
	}
	view.Render(w, viewModel)
}
