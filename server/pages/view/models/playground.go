package models

import "strings"

// The Playground struct provides state information for rendering the playground view.
// For example, which is the currently selected page.
type Playground struct {
	SubmitUrl		 string
	SideBySidePanels         bool
	TabbedPanels             bool
	ShowInputTab             bool
	ShowFrostTab             bool
	TabsOrSideBySideLinkText string
	InputText                string
	OutputText               string
}

// The NewPlayground function is a factory that makes Playground view models.
func NewPlayground(inputTxt string, outputTxt string, hintsString string) *Playground {
	if strings.Contains(hintsString, "SIDES_BY_SIDE_VIEW") {
		return &Playground{
			SubmitUrl:		  "/playground-refresh-side-by-side",
			SideBySidePanels:         true,
			TabbedPanels:             false,
			ShowInputTab:             false,
			ShowFrostTab:             true,
			TabsOrSideBySideLinkText: "Tabbed view",
			InputText:                inputTxt,
			OutputText:               outputTxt,
		}
	} else if strings.Contains(hintsString, "TABBED_VIEW") {
		return &Playground{
			SubmitUrl:		  "/playground-refresh-tabbed",
			SideBySidePanels:         false,
			TabbedPanels:             true,
			ShowInputTab:             false,
			ShowFrostTab:             true,
			TabsOrSideBySideLinkText: "Side by side view",
			InputText:                inputTxt,
			OutputText:               outputTxt,
		}
	} else {
		panic("Hints strings not recognized")
	}
}
