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
const PlaygroundRefreshInputTab = "/playground/refresh/input-tab"
const PlaygroundRefreshOutputTab = "/playground/refresh/output-tab"

// File system paths
const PlaygroundExamplePath = "static/examples/space_delim.txt"

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
	InputTextElementName string
	InputText  string
	OutputText string

	FormAction       string
	SwitchViewAction string
	SwitchViewLabel  string

	ShowTabbed     bool
	ShowSideBySide bool

	ShowInputTab  bool
	ShowOutputTab bool
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

	// Set constant fields
	pg.InputTextElementName = PlaygroundInputTextField

	// Grab the input text and make the output text assuming a whitespace conversion
	// is required.
	pg.InputText = submittedForm.Get(PlaygroundInputTextField)
	pg.OutputText = pg.doWhiteSpaceConversionForNow(pg.InputText)

	// For a refresh action we reflect back the incoming URL
	pg.FormAction = urlPath

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

// The NewPlaygroundViewModelForExample function creates a new PlaygroundViewModel instance that
// is suitable for rendering the playground page pre-populated with the given input text.
func NewPlaygroundViewModelForExample(exampleInputText string) *PlaygroundViewModel {

	mdl := &PlaygroundViewModel{}

	// Set constant fields
	mdl.InputTextElementName = PlaygroundInputTextField

	mdl.InputText = exampleInputText
	mdl.OutputText = mdl.doWhiteSpaceConversionForNow(mdl.InputText)

	// Show the example initially in side by side view
	mdl.FormAction = PlaygroundRefreshSides
	mdl.SwitchViewAction = PlaygroundRefreshInputTab
	mdl.SwitchViewLabel = PlaygroundRefreshSwitchToTabsLabel
	mdl.ShowSideBySide = true

	return mdl
}

func (pg *PlaygroundViewModel) doWhiteSpaceConversionForNow(inputText string) string {
	return string(filereaders.NewWhitespaceConverter(inputText).Convert())
}

func (pg *PlaygroundViewModel) panicCannotDecodeUrl(urlPath string) {
	panic(fmt.Sprintf("Cannot decode this playground refresh URL: %v", urlPath))
}
