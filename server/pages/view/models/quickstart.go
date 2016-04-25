package models

// The QuickStart struct provides state information for rendering the quick start view.
type QuickStart struct {
}

// The NewQuickStart function is a factory that makes QuickStart view models.
func NewQuickStart() *QuickStart {
	model := &QuickStart{}
	return model
}