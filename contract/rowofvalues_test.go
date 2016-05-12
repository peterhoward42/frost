package contract

import (
	"testing"
)

// First a bunch of special cases - to ensure these are working and not confusing the
// mainstream tests that follow.

func TestHasSameSignatureAs(t *testing.T) {
	rowA := NewRowOfValues([]string{"foo", "42", "3.14", "false"})
	rowB := NewRowOfValues([]string{"hello", "999", "0.001", "true"})
	if rowA.HasSameSignatureAs(rowB) == false {
		t.Errorf("Incorrectly judge row signatures to be different.")
	}
	// Row C differs from A only by the second field being float rather than int.
	rowC := NewRowOfValues([]string{"foo", "42.0", "3.14", "false"})
	if rowA.HasSameSignatureAs(rowC) == true {
		t.Errorf("Incorrectly judge row signatures to be the same.")
	}
}
