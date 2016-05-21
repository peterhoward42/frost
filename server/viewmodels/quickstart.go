package viewmodels

// The QuickStartViewModel type provides state information for rendering the quick start view.
type QuickStartViewModel struct {
}

// The NewPlaygroundViewModelForExample function creates a new QuickStartViewModel instance that
// is suitable for rendering the quickstart  page.
func NewQuickStartViewModel() *QuickStartViewModel {

	mdl := &QuickStartViewModel{}
	return mdl
}
