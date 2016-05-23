package viewmodels

import (
	"github.com/peterhoward42/frost/server/urls"
)

// The TopLevelViewModel is the model for the top level application view (in the MVC sense).
// It is recommended that instances are created using the constructor functions.
type TopLevelViewModel struct {
	// One of these pointers being set to non-null signifies the viewing mode required as
	// well as addressing the sub model for that mode.
	QuickStart *QuickStartViewModel
	Playground *PlaygroundViewModel
	ComingSoon *ComingSoonViewModel // not really a model, but keeps things symmetrical

	QuickStartURL string
	PlaygroundURL string
}

func NewTopLevelViewModel() *TopLevelViewModel {
	return &TopLevelViewModel{
		QuickStartURL: urls.URLQuickstart,
		PlaygroundURL: urls.URLPlaygroundExampleSpaceDelim,
		ComingSoon: nil,
	}
}
