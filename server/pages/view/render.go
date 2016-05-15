package view

import (
	"fmt"
	"github.com/peterhoward42/frost/server/pages/view/viewmodels"
	"net/http"
)

func Render(w http.ResponseWriter, viewModel *viewmodels.TopLevel) {
	err := view.GlobalGuiTemplate.ExecuteTemplate(w, "maingui.html", viewModel)
	if err != nil {
		panic(fmt.Sprintf("Cannot execute template. Error is: %v", err))
	}
}
