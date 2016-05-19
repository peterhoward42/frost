package resources

import (
	"testing"
)

func TestCompiledFileSystem(t *testing.T) {
	fs := CompiledFileSystem
	_, err := fs.Open(`/static/examples/space_delim.txt`)
	if err != nil {
		t.Errorf("Error from fs.Open(): %v", err)
	}
}
