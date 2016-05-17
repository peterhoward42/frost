package view

import (
	"strings"
	"testing"
	"bytes"
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/viewmodels"
)

func TestRenderingOfQuickStartPage(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	viewModel.QuickStart = &viewmodels.QuickStartViewModel{}
	renderer := NewGuiRenderer(resources.CompiledTemplates, "maingui.html")
	var recorder bytes.Buffer
	renderer.Render(&recorder, viewModel)
	outputAsString := recorder.String()
	if strings.Contains(outputAsString, "shouldincludethis") == false {
		t.Errorf("Rendered content (follows) does not include what is expected: %v",
			outputAsString)
	}
}
