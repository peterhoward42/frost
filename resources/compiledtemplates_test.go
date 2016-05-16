package resources

import (
	"testing"
	"strings"
)

func TestCompiledTemplates(t *testing.T) {
	definedTemplates := CompiledTemplates.DefinedTemplates()
	if strings.Contains(definedTemplates, "commonfooter.html") == false {
		t.Errorf("Did not find and parse the right templates.")
	}
}
