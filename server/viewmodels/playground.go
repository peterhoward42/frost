package viewmodels

import (
	"fmt"
	"github.com/peterhoward42/frost/filereaders"
	"net/url"
	"strings"
)

// Form element names
const PlaygroundInputTextField = "input-text"

// URL building blocks
const PlaygroundSideBySideUrlFragment = "side-by-side"
const PlaygroundInputTabUrlFragment = "input-tab"
const PlaygroundOutputTabUrlFragment = "output-tab"
const PlaygroundRefreshTabUrlSubStr = "-tab"

// Full URLs
const PlaygroundRefreshSides = "/playground/refresh/" + PlaygroundSideBySideUrlFragment
const PlaygroundRefreshInputTab = "/playground/refresh/" + PlaygroundInputTabUrlFragment
const PlaygroundRefreshOutputTab = "/playground/refresh/" + PlaygroundOutputTabUrlFragment

// Labels for humans
const PlaygroundRefreshSwitchToTabsLabel = "Tabbed view"
const PlaygroundRefreshSideBySideLabel = "Side by side view"

// The PlaygroundViewModel is the model for the playground sub view (in the MVC sense).
// It is recommended that instances are created using the constructor functions.
type PlaygroundViewModel struct {
	// There is partly duplicated data in these fields. Which is regrettable for maintenance,
	// but has been chosen to minimise the logic required downstream in the template that
	// consumes the model. We provide the data in a form that makes the templates as simple as
	// possible, preferring to carry the complexity in Go rather than in the templating
	// language.
	InputTextElementName string // Use this as the name for the input text TextArea.
	InputText            string // Set the text of the input text TextArea to this.
	OutputText           string // Set the text of the output TextArea to this.

	// Use this as the active attribute on the space separated example button. I.e. "active"
	// or empty string.
	SpaceSepActiveMode string

	FormAction       string // Use this as the URL in the form's action attribute.
	SwitchViewAction string // Use this as the form-action attribute on the switch view button.
	SwitchViewLabel  string // Use this text label on the switch view button.

	ShowTabbed     bool // Should the view be in tabbed mode?
	ShowSideBySide bool // should the view be in side by side mode?

	ShowInputTab           bool   // When in tabbed mode, should the input tab be on top?
	ShowOutputTab          bool   // When in tabbed mode, should the output tab be on top?
	FormActionForInputTab  string // URL to post form to when you press the Input Tab
	FormActionForOutputTab string // URL to post form to when you press the Input Tab
}

// The NewPlaygroundViewModelForExample function creates a new PlaygroundViewModel instance that
// is suitable for rendering the playground page pre-populated with the given input text.
func NewPlaygroundViewModelForExample(
	exampleInputText string,
	spaceSeparatedButtonActiveString string) *PlaygroundViewModel {

	mdl := &PlaygroundViewModel{}
	mdl.setConstantFields()

	mdl.InputText = exampleInputText
	mdl.OutputText = mdl.doWhiteSpaceConversionForNow(mdl.InputText)

	// When we are rendering a user's request to see an example, we might want to show
	// the space separate example button as active - we require the caller of this function
	// to inject this decision.
	mdl.SpaceSepActiveMode = spaceSeparatedButtonActiveString

	// Show the example initially in side by side view
	mdl.FormAction = PlaygroundRefreshSides
	mdl.SwitchViewAction = PlaygroundRefreshInputTab
	mdl.SwitchViewLabel = PlaygroundRefreshSwitchToTabsLabel
	mdl.ShowSideBySide = true

	return mdl
}

// The NewPlaygroundViewModelForRefresh function creates a new PlaygroundViewModel instance that
// is suitable for refreshing an existing page that is derived from this model. It takes the user's
// current input text from the form submitted with the request, and renders the converted output
// in the sister pane alongside or in a tabbed view. A family of URLs are routed here, which encode
// the viewing style requested in the URL.
func NewPlaygroundViewModelForRefresh(
	submittedForm url.Values,
	urlPath string) *PlaygroundViewModel {

	pg := &PlaygroundViewModel{}
	pg.setConstantFields()

	// Grab the input text and make the output text assuming a whitespace conversion
	// is required.
	pg.InputText = submittedForm.Get(PlaygroundInputTextField)
	pg.OutputText = pg.doWhiteSpaceConversionForNow(pg.InputText)

	// For a refresh action we reflect back the incoming URL
	pg.FormAction = urlPath

	// For a refresh action we do not want the space-separated example button to show up
	// as being active.
	pg.SpaceSepActiveMode = ""

	// Set up model fields to suit either side by side mode or tabbed mode.
	// (Exploits the zero value of the structure)
	switch {
	// Using side by side view?
	case strings.Contains(urlPath, PlaygroundSideBySideUrlFragment):
		pg.SwitchViewAction = PlaygroundRefreshInputTab
		pg.SwitchViewLabel = PlaygroundRefreshSwitchToTabsLabel
		pg.ShowSideBySide = true

	// Using tabbed view?
	case strings.Contains(urlPath, PlaygroundRefreshTabUrlSubStr):
		pg.SwitchViewAction = PlaygroundRefreshSides
		pg.SwitchViewLabel = PlaygroundRefreshSideBySideLabel
		pg.ShowTabbed = true
		switch {
		case strings.Contains(urlPath, PlaygroundInputTabUrlFragment):
			pg.ShowInputTab = true
		case strings.Contains(urlPath, PlaygroundOutputTabUrlFragment):
			pg.ShowOutputTab = true
		default:
			pg.panicCannotDecodeUrl(urlPath)
		}

	default:
		pg.panicCannotDecodeUrl(urlPath)
	}

	return pg
}

func (pg *PlaygroundViewModel) setConstantFields() {
	pg.InputTextElementName = PlaygroundInputTextField
	pg.FormActionForInputTab = PlaygroundRefreshInputTab
	pg.FormActionForOutputTab = PlaygroundRefreshOutputTab
}

func (pg *PlaygroundViewModel) doWhiteSpaceConversionForNow(inputText string) string {
	return string(filereaders.NewWhitespaceConverter(inputText).Convert())
}

func (pg *PlaygroundViewModel) panicCannotDecodeUrl(urlPath string) {
	panic(fmt.Sprintf("Cannot decode this playground refresh URL: %v", urlPath))
}
