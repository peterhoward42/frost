package view

import "strings"

type Model struct {
	QuickStartActive bool
	PlayGroundActive bool
}

func NewModel(activePageHint string) *Model {
	model := &Model{
		QuickStartActive: strings.Contains(activePageHint, "uick"),
		PlayGroundActive: strings.Contains(activePageHint, "ground"),
	}
	return model
}
