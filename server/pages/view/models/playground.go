package models

// The Playground struct provides state information for rendering the playground view.
// For example, which is the currently selected page.
type Playground struct {
	InputText  string
	OutputText string
}

// The NewPlayground function is a factory that makes Playground view models.
func NewPlayground(input_txt string,
	output_txt string) *Playground {
	model := &Playground{input_txt, output_txt}
	return model
}
