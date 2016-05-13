// The package models provides a set of state models that are used when rendering views.
// For example which page has been selected.
package models

import (
	"fmt"
	"strings"
)

// The TopLevel struct provides state information for rendering the top level application view.
// For example, which is the currently selected page.
type TopLevel struct {
	// Just one of these sub-model fields will be set to non-null to signify the currently
	// selected page.
	QuickStart *QuickStart
	Playground *Playground
}

// The NewTopLevel function is a factory that makes TopLevel view models.
// You inject a string fragment to indicate options.
func NewTopLevel(hintsString string, playgroundInputText string,
	playgroundOutputText string) *TopLevel {
	model := &TopLevel{}

	switch {
	case strings.Contains(hintsString, "QUICKSTART"):
		model.QuickStart = NewQuickStart()
	case strings.Contains(hintsString, "PLAYGROUND"):
		model.Playground = NewPlayground(playgroundInputText, playgroundOutputText,
			hintsString)
	default:
		panic(fmt.Sprintf("Hints string not recognized: %v", hintsString))
	}

	return model
}
