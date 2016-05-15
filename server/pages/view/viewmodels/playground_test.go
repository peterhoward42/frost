package viewmodels

import (
	"net/url"
	"testing"
	"strings"
)

// First tests for the refresh URL



func TestNewPlaygroundViewModelForRefreshProducesOutputPanelContents(t *testing.T) {
	urlPath := "/playground/refresh/side-by-side"
	form := url.Values{}
	form.Set("input-text", "hello\ngoodbye")
	mdl := NewPlaygroundViewModelForRefresh(form, urlPath)
	outputPanelTxt := mdl.OutputText
	if strings.Contains(outputPanelTxt, "Type") == false {
		t.Errorf("Output panel text does not contain necessary fragment.")
	}
}

func TestNewPlaygroundViewModelForRefreshReflectsUrl(t *testing.T) {
	urlPath := "/playground/refresh/side-by-side"
	form := url.Values{}
	form.Set("input-text", "hello\ngoodbye")
	mdl := NewPlaygroundViewModelForRefresh(form, urlPath)
	if mdl.FormAction != urlPath {
		t.Errorf("Failed to reflect incoming URL in form action.")
	}
}

func TestNewPlaygroundViewModelForRefreshSetsStateRightForSideBySideMode(t *testing.T) {
	urlPath := "anything with side-by-side in it"
	form := url.Values{}
	mdl := NewPlaygroundViewModelForRefresh(form, urlPath)
	if mdl.SwitchViewAction != "/playground/refresh/input-tab" {
		t.Errorf("Unexpected switch view action: %v", mdl.SwitchViewAction)
	}
	if mdl.SwitchViewLabel != "Tabbed view" {
		t.Errorf("Unexpected switch view label: %v", mdl.SwitchViewLabel)
	}
	if mdl.ShowSideBySide != true {
		t.Errorf("Failed to mandate showing side by side")
	}
}

func TestNewPlaygroundViewModelForRefreshSetsStateRightForInputTabbedMode(t *testing.T) {
	urlPath := "anything with input-tab in it"
	form := url.Values{}
	mdl := NewPlaygroundViewModelForRefresh(form, urlPath)
	if mdl.SwitchViewAction != "/playground/refresh/side-by-side" {
		t.Errorf("Unexpected switch view action: %v", mdl.SwitchViewAction)
	}
	if mdl.SwitchViewLabel != "Side by side view" {
		t.Errorf("Unexpected switch view label: %v", mdl.SwitchViewLabel)
	}
	if mdl.ShowInputTab!= true {
		t.Errorf("Failed to mandate showing input tab")
	}
}

func TestNewPlaygroundViewModelForRefreshSetsStateRightForOutputTabbedMode(t *testing.T) {
	urlPath := "anything with output-tab in it"
	form := url.Values{}
	mdl := NewPlaygroundViewModelForRefresh(form, urlPath)
	if mdl.SwitchViewAction != "/playground/refresh/side-by-side" {
		t.Errorf("Unexpected switch view action: %v", mdl.SwitchViewAction)
	}
	if mdl.SwitchViewLabel != "Side by side view" {
		t.Errorf("Unexpected switch view label: %v", mdl.SwitchViewLabel)
	}
	if mdl.ShowOutputTab!= true {
		t.Errorf("Failed to mandate showing output tab")
	}
}


// Now tests for the example URL



func TestNewPlaygroundViewModelForExampleProducesCorrectPanelContents(t *testing.T) {
	mdl := NewPlaygroundViewModelForExamples()
	expected := "fibble"
	if mdl.InputText != expected {
		t.Errorf("Unexpected input panel text: %v", mdl.InputText)
	}
}
