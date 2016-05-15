package view

import (
	"fmt"
	"github.com/peterhoward42/frost/server/pages/view/viewmodels"
	"html/template"
	"io"
)

// The GuiRenderer type combines a template and a corresponding view model to generate the html
// that renders the Gui.
type GuiRenderer struct {
	guiTemplate                *template.Template
	templateEntryPointFilename string
}

// The NewGuiRenderer function is the recommended way to instantiate a GuiRenderer, and provides
// a way to inject the template required.
func NewGuiRenderer(
	guiTemplate *template.Template, templateEntryPointFilename string) *GuiRenderer {
	return &GuiRenderer{
		guiTemplate:                guiTemplate,
		templateEntryPointFilename: templateEntryPointFilename,
	}
}

func (r *GuiRenderer) Render(w io.Writer, guiViewModel *viewmodels.TopLevelViewModel) {
	err := r.guiTemplate.ExecuteTemplate(w, r.templateEntryPointFilename, guiViewModel)
	if err != nil {
		panic(fmt.Sprintf("Cannot execute template. Error is: %v", err))
	}
}
