package view

import (
	"fmt"
	"github.com/peterhoward42/frost/server/viewmodels"
	"html/template"
	"io"
)

const templateEntryPointName = "maingui.html"

// The GuiRenderer combines a template and a corresponding view model to generate the html
// that renders the Gui.
type GuiRenderer struct {
	guiTemplate *template.Template
}

// The NewGuiRenderer function is the recommended way to instantiate a GuiRenderer, and provides
// a way to inject the template required.
func NewGuiRenderer(
	guiTemplate *template.Template) *GuiRenderer {
	return &GuiRenderer{guiTemplate: guiTemplate}
}

func (r *GuiRenderer) Render(w io.Writer, guiViewModel *viewmodels.TopLevelViewModel) {
	err := r.guiTemplate.ExecuteTemplate(w, templateEntryPointName, guiViewModel)
	if err != nil {
		panic(fmt.Sprintf("Cannot execute template. Error is: %v", err))
	}
}
