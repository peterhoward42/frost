package view

import (
	"strings"
	"testing"
	"github.com/peterhoward42/frost/server/pages/view/viewmodels"
	"bytes"
)

func TestRenderingOfQuickStartPage(t *testing.T) {

	template, _ := how get template????
	renderer := NewGuiRenderer(template, "fred")
	ioWriterThatRecords := &bytes.Buffer{}

	viewModel := viewmodels.TopLevelViewModel{QuickStart: &viewmodels.QuickStartViewModel{}}
	renderer.Render(ioWriterThatRecords, viewModel)
	rendered := string(ioWriterThatRecords.Bytes())
	if strings.Contains(rendered, "shouldincludethis") == false {
		t.Errorf("Rendered content (follows) does not include what is expected: %v",
			renderedContentAsString)
	}
}
