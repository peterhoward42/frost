package viewmodels

import (
	"net/url"
	"testing"
	"strings"
)

// First a series of tests for the refresh URL


func TestRefreshProducesOutputPanelContents(t *testing.T) {
	urlPath := "/playground/refresh/side-by-side"
	form := url.Values{}
	form.Set("input-text", "hello\ngoodbye")
	mdl := NewPlaygroundViewModelForRefresh(form, urlPath)
	outputPanelTxt := mdl.OutputText
	if strings.Contains(outputPanelTxt, "Type") == false {
		t.Errorf("Output panel text does not contain necessary fragment.")
	}
}

func TestRefreshReflectsUrl(t *testing.T) {
	urlPath := "/playground/refresh/side-by-side"
	form := url.Values{}
	form.Set("input-text", "hello\ngoodbye")
	mdl := NewPlaygroundViewModelForRefresh(form, urlPath)
	if mdl.FormAction != urlPath {
		t.Errorf("Failed to reflect incoming URL in form action.")
	}
}

func TestRefreshSetsStateRightForSideBySideMode(t *testing.T) {
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

func TestRefreshSetsStateRightForInputTabbedMode(t *testing.T) {
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

func TestRefreshSetsStateRightForOutputTabbedMode(t *testing.T) {
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


// Now a series of tests for the set of example button action URLs

func TestExamplePrePopulatesInputPanel(t *testing.T) {
	inputText := "foo\nbar"
	mdl := NewPlaygroundViewModelForExample(inputText)
	if mdl.InputText != inputText {
		t.Errorf("Unexpected input panel text: %v", mdl.InputText)
	}
}

func TestExamplePopulatesOutputPanel(t *testing.T) {
	inputText := "foo\nbar"
	mdl := NewPlaygroundViewModelForExample(inputText)
	expected := `"Value": "foo"`
	if strings.Contains(mdl.OutputText, expected) == false {
		t.Errorf("Unexpected output panel text: %v", mdl.OutputText)
	}
}

func TestExampleSetsUpSideBySideViewOptions(t *testing.T) {
	inputText := "irrelevant"
	mdl := NewPlaygroundViewModelForExample(inputText)
	if mdl.FormAction != "/playground/refresh/side-by-side" {
		t.Errorf("Unexpected form action: %v", mdl.FormAction)
	}
	if mdl.SwitchViewAction != "/playground/refresh/input-tab" {
		t.Errorf("Unexpected switch view action: %v", mdl.SwitchViewAction)
	}
	if mdl.SwitchViewLabel != "Tabbed view" {
		t.Errorf("Unexpected switch view label: %v", mdl.SwitchViewLabel)
	}
	if mdl.ShowSideBySide != true{
		t.Errorf("Unexpected show side by side: %v", mdl.ShowSideBySide)
	}
}
