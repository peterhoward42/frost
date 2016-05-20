package view

import (
	"bytes"
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/viewmodels"
	"github.com/peterhoward42/frost/testutils"
	"strings"
	"testing"
	"net/url"
	"fmt"
)

// Simple smoke test for playground rendering for a given input text content.
func TestRenderingOfPlayGroundLanding(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	exampleInputText := string(resources.MustAsset(`static/examples/space_delim.txt`))
	viewModel.Playground = viewmodels.NewPlaygroundViewModelForExample(exampleInputText)
	renderer := NewGuiRenderer(resources.CompiledTemplates)
	var recorder bytes.Buffer
	renderer.Render(&recorder, viewModel)
	rendered := recorder.String()
	flattened := strings.Replace(rendered, "\n", " ", -1)

	assertions := []*testutils.MatchAssertion{}
	assertions = append(assertions, &testutils.MatchAssertion{
		`Use less code to read`,
		`Smoke test of standard header.`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`active.*href.*/playground/example.*Playground`,
		`Playground button is active and has correct URL`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<li.*href.*/quickstart.*Quick Start.*</li`,
		`Quickstart button is not active, and has correct url`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`Switch to.*<button.*type="submit".*formaction=/playground/refresh/input-tab>.*Tabbed view`,
		`Switch to button offers tabbed view and has tabbed URL as the action`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<form action="/playground/refresh/side-by-side" method="post">`,
		`Main form action asks for refresh side by side`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<div class="row">.*<div class="col-lg-6">.*<textarea.*name="input-text"`,
		`Left hand side contains correctly named text area over 6 columns`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`# name      gamma   idn .*01_B_INN    0.00    42  `,
		`Input text area is pre-populated with text from example file`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`</div>.* <div class="col-lg-6">.*<textarea`,
		`Right hand side contains converted text in 6 columns`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`</textarea>.*</div>.*</div>.*</form>`,
		`Form is properly terminated`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`</div class="container">.*</body>.*</html>`,
		`Smoke test of standard footer.`})

	testutils.AssertPageContainsSamples(t, flattened, assertions)
}

// Simple smoke test for playground refresh in side by side mode.
func TestRenderingOfPlayGroundRefreshSideBySide(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	formValues := url.Values{}
	// Use arbitrary URL that in includes the trigger tags
	incumbentURL := "xxx-side-by-side-xxx"
	viewModel.Playground = viewmodels.NewPlaygroundViewModelForRefresh(formValues, incumbentURL)
	renderer := NewGuiRenderer(resources.CompiledTemplates)
	var recorder bytes.Buffer
	renderer.Render(&recorder, viewModel)
	rendered := recorder.String()
	flattened := strings.Replace(rendered, "\n", " ", -1)

	assertions := []*testutils.MatchAssertion{}

	assertions = append(assertions, &testutils.MatchAssertion{
		`Use less code to read`,
		`Smoke test of standard header.`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<form action="xxx-side-by-side-xxx" method="post">`,
		`The form action has captured and repeated the incoming URL`})
	// We test that none of the playground buttons is active by sampling the
	// one that is active for the landing page - ie the space delim text button.
	assertions = append(assertions, &testutils.MatchAssertion{
		`<a class="btn btn-default[\s]*href="#">Space Separated</a>`,
		`None of the playground buttons should be active`})

	testutils.AssertPageContainsSamples(t, flattened, assertions)

	fmt.Printf("The page under test is:\n%v", rendered)
}
