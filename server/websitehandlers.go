package server

import (
	"fmt"
	"github.com/peterhoward42/frost/server/pages/view"
	"github.com/peterhoward42/frost/server/pages/view/models"
	"net/http"
	"strings"
)

func HandleQuickStart(w http.ResponseWriter, r *http.Request) {
	viewModel := &models.TopLevel{}
	viewModel.QuickStart = models.QuickStart{}
	view.Render(w, viewModel)
}

func HandlePlayground(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(fmt.Sprintf("Cannot parse form. Error is: %v", err))
	}
	var viewModel *models.TopLevel
	submittedForm := r.Form
	switch {
	case strings.Contains(r.URL.Path, "refresh"):
		viewModel = models.NewPlaygroundViewModelForRefresh(
			w, submittedForm, r.URL.Path)
	case strings.Contains(r.URL.Path, "example"):
		viewModel = models.NewPlaygroundViewModelForExamples(w)
	}
	view.Render(w, viewModel)
}
