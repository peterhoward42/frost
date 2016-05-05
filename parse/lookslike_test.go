package parse

import (
	"fmt"
	"testing"
)

func TestLooksLikeAnIntegerWithValidInputs(t *testing.T) {
	legitInputs := []string{"42", "+42", "-42", "4"}
	expected := []int{42, 42, -42, 4}
	for idx, s := range legitInputs {
		matched, v := LooksLikeAnInteger(s)
		if matched != true {
			t.Errorf("Failed to recognize: %v", s)
		}
		if v != expected[idx] {
			t.Errorf("Failed to extract value from: %v", s)
		}
	}
}

func TestLooksLikeAnIntegerRejectsMalformedInputs(t *testing.T) {
	malformed := []string{" 42", "42 ", "a42", "42a42", "42a", "42.0", "42."}
	for _, s := range malformed {
		matched, _ := LooksLikeAnInteger(s)
		if matched == true {
			t.Errorf("Spuriously recognized: %v", s)
		}
	}
}

func TestLooksLikeAFloatWithValidInputs(t *testing.T) {
	legitInputs := []string{"1.0", "+1.0", "-1.0", ".1", "0.11"}
	// Casting the floats received back to formatted strings makes comparison easier.
	expected := []string{"1.00", "1.00", "-1.00", "0.10", "0.11"}
	for idx, s := range legitInputs {
		matched, v := LooksLikeAFloat(s)
		if matched != true {
			t.Errorf("Failed to recognize: %v", s)
		}
		got := fmt.Sprintf("%.2f", v)
		if got != expected[idx] {
			t.Errorf("Wrong value extracted (%v) from: %v", got, s)
		}
	}
}

func TestLooksLikeAFloatRejectsMalformedInputs(t *testing.T) {
	malformed := []string{" 4.2", "4.2 ", "a4.2", "4.2a42", "4.2a", "42"}
	for _, s := range malformed {
		matched, _ := LooksLikeAFloat(s)
		if matched == true {
			t.Errorf("Spuriously recognized: %v", s)
		}
	}
}

func TestLooksLikeABoolWithValidInputs(t *testing.T) {
	legitInputs := []string{"true", "tRuE", "false", "fAlSe"}
	expected := []bool{true, true, false, false}
	for idx, s := range legitInputs {
		matched, v := LooksLikeABool(s)
		if matched != true {
			t.Errorf("Failed to recognize: %v", s)
		}
		if v != expected[idx] {
			t.Errorf("Failed to extract value from: %v", s)
		}
	}
}

func TestLooksLikeABoolRejectsMalformedInputs(t *testing.T) {
	malformed := []string{" true", "false "}
	for _, s := range malformed {
		matched, _ := LooksLikeABool(s)
		if matched == true {
			t.Errorf("Spuriously recognized: %v", s)
		}
	}
}
