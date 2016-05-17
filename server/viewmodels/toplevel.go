package viewmodels

import (
	"github.com/peterhoward42/frost/server/routing"
)

type TopLevelViewModel struct {
	// One of these pointers being set to non-null signifies the viewing mode required as
	// well as addressing the sub model for that mode.
	QuickStart *QuickStartViewModel
	Playground *PlaygroundViewModel

	QuickStartURL string
	PlaygroundURL string
}

func NewTopLevelViewModel() *TopLevelViewModel {
	return &TopLevelViewModel{
		QuickStartURL: routing.URLQuickstart,
		PlaygroundURL: routing.URLPlayground,
		}
}