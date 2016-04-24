// The package models provides a set of state models that are used when rendering views.
// For example which page has been selected.
package models

import (
	"strings"
	"fmt"
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
// You inject a string fragment to indicate the currently selected page required.
func NewTopLevel(activePageHint string) *TopLevel {
	model := &TopLevel{}

	switch {
	case strings.Contains(activePageHint, "uick"):
		model.QuickStart = &QuickStart{}
	case strings.Contains(activePageHint, "ground"):
		model.Playground = &Playground{}
	default:
		panic(fmt.Sprintf("Active page hint not recognized: %v", activePageHint))
	}

	return model
}