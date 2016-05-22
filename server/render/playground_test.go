package view

import (
	"bytes"
	"github.com/peterhoward42/frost/resources"
	"github.com/peterhoward42/frost/server/viewmodels"
	"github.com/peterhoward42/frost/testutils"
	"net/url"
	"strings"
	"testing"
)

// Simple smoke test for playground rendering for a given input text content.
func TestRenderingOfCannedExample(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	exampleInputText := resources.GetExampleFileContents(resources.SpaceDelimitedExample)
	spaceSeparatedButtonActiveString := "active"
	viewModel.Playground = viewmodels.NewPlaygroundViewModelForExample(
		exampleInputText, spaceSeparatedButtonActiveString)
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
		`<a class="btn btn-default active".*href=.*Space Separated`,
		`The space separated example button is active.`})
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

func TestRenderingOfPlayGroundRefreshSideBySide(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	formValues := url.Values{}
	// Provide some form input so we can make sure the refresh response echos it back, in the
	// input pane, and shows the converted value in the output.
	formValues.Set(viewmodels.PlaygroundInputTextField, "fibble")

	// Use arbitrary URL that  includes the side by side trigger tag
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
		`<div class="btn-group">.*<a class="btn btn-default[\s]*"[\s]*href="#">Space`,
		`None of the playground buttons should be active`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`name="input-text.*fibble`,
		`The input text area has picked up the incoming contents.`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<textarea.*&#34;fibble&#34`,
		`The output text area has been rendered with the conversion`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`Switch to.*<button.*type="submit".*formaction=/playground/refresh/input-tab>.*Tabbed view`,
		`Switch to button offers tabbed view and has tabbed URL as the action`})

	testutils.AssertPageContainsSamples(t, flattened, assertions)

	// fmt.Printf("The page under test is:\n%v", rendered)
}

func TestRenderingPlayGroundRefreshTabbedMode(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	formValues := url.Values{}
	formValues.Set(viewmodels.PlaygroundInputTextField, "fibble")

	// Use arbitrary URL that includes the tag that stimulates tabbed mode with the
	// input tab in front.
	incumbentURL := "xxx-input-tab-xxx"
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
		`Switch to.*formaction=/playground/refresh/side-by-side>.*Side by side view`,
		`Switch to button offers side by side view and has side by side URL as the action`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<button[\s]*class="btn btn-default active".*input-tab">[\s]*Input`,
		`Input tab is active`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`class="btn btn-default ".*resh/output-tab">`,
		`Output tab is not active`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<textarea.*name="input-text".*<div class="hidden" >.*<textarea`,
		`Input tab is visible`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`type="submit"[\s]*formaction="/playground/refresh/input-tab">[\s]*Input`,
		`Href is correct on input tab`})

	testutils.AssertPageContainsSamples(t, flattened, assertions)

	// fmt.Printf("The page under test is:\n%v", rendered)

}
func TestRenderingPlayGroundRefreshTabbedModeVariantsWhenOutputTabSelected(t *testing.T) {
	viewModel := viewmodels.NewTopLevelViewModel()
	formValues := url.Values{}
	formValues.Set(viewmodels.PlaygroundInputTextField, "fibble")

	// Use arbitrary URL that includes the tag that stimulates tabbed mode with the
	// input tab in front.
	incumbentURL := "xxx-output-tab-xxx"
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
		`Switch to.*formaction=/playground/refresh/side-by-side>.*Side by side view`,
		`Switch to button offers side by side view and has side by side URL as the action`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`<button[\s]*class="btn btn-default active".*output-tab">[\s]*Frost`,
		`Output tab is active`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`class="btn btn-default ".*resh/output-tab">`,
		`Input tab is not active`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`class="hidden".*input-text".*<div >.*\[`,
		`Output tab is visible`})
	assertions = append(assertions, &testutils.MatchAssertion{
		`type="submit"[\s]*formaction="/playground/refresh/output-tab">[\s]*Frost`,
		`Href is correct on output tab`})

	testutils.AssertPageContainsSamples(t, flattened, assertions)

	// fmt.Printf("The page under test is:\n%v", rendered)



}
