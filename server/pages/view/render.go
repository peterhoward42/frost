package view

import (
	"fmt"
	"github.com/peterhoward42/frost/server/pages/view/models"
	"net/http"
)

func Render(w http.ResponseWriter, viewModel *models.TopLevel) {
	err := view.GlobalGuiTemplate.ExecuteTemplate(w, "maingui.html", viewModel)
	if err != nil {
		panic(fmt.Sprintf("Cannot execute template. Error is: %v", err))
	}
}
