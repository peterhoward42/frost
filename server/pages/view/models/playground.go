package models

// The Playground struct provides state information for rendering the playground view.
// For example, which is the currently selected page.
type Playground struct {
	SideBySidePanels         bool
	TabbedPanels             bool
	ShowInputTab             bool
	ShowFrostTab             bool
	TabsOrSideBySideLinkText string
	InputText                string
	OutputText               string
}

// The NewPlayground function is a factory that makes Playground view models.
func NewPlayground(inputTxt string,
	outputTxt string) *Playground {
	model := &Playground{
		SideBySidePanels:         false,
		TabbedPanels:             true,
		ShowInputTab:             false,
		ShowFrostTab:             true,
		TabsOrSideBySideLinkText: "Tabbed view",
		InputText:                inputTxt,
		OutputText:               outputTxt,
	}
	return model
}
