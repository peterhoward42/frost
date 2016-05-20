package view

import (
	"bytes"
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/viewmodels"
	"github.com/peterhoward42/frost/testutils"
	"strings"
	"testing"
)

// This is a smoke-test sample of the html produced to render the landing (quickstart) page.
func TestRenderingOfQuickStartPage(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	viewModel.QuickStart = viewmodels.NewQuickStartViewModel()
	renderer := NewGuiRenderer(resources.CompiledTemplates)
	var recorder bytes.Buffer
	renderer.Render(&recorder, viewModel)
	rendered := recorder.String()
	flattened := strings.Replace(rendered, "\n", " ", -1)

	assertions := []*testutils.MatchAssertion{}
	assertions = append(assertions, &testutils.MatchAssertion{
		`Use less code to read`,
		`Sample from standard header`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`active.*href.*quickstart.*Quick Start`,
		`Quickstart button is active`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<li.*href.*playground.*Playground.*</li`,
		`Playground button is not active`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`melting_points`,
		`Fragment from quickstart main body`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`</body>.*</html>`,
		`Fragment from standard footer`})

	testutils.AssertPageContainsSamples(t, flattened, assertions)
}
