package parse

import (
	"testing"
)

// Set of tests for disguising operation

func TestDisguisingNoOp(t *testing.T) {
	input := `hello`
	masked := DisguiseDoubleQuotedSegments(input)
	expected := `hello`
	if masked != expected {
		t.Errorf("Mask wrong. Got: <%v>, expected: <%v>", masked, expected)
	}
}

func TestDisguisingSingleQuotedSection(t *testing.T) {
	input := `ppp "aaa aaa aaa" qqq`
	masked := DisguiseDoubleQuotedSegments(input)
	expected := `ppp aaa*($&x12#a~aaa*($&x12#a~aaa qqq`
	if masked != expected {
		t.Errorf("Mask wrong. Got: <%v>, expected: <%v>", masked, expected)
	}
}

func TestDisguisingEmptyQuotedSection(t *testing.T) {
	input := `p""q`
	masked := DisguiseDoubleQuotedSegments(input)
	expected := `pq`
	if masked != expected {
		t.Errorf("Mask wrong. Got: <%v>, expected: <%v>", masked, expected)
	}
}

func TestDisguisingMultipleQuotedSections(t *testing.T) {
	input := `p"a a"q"b b"r`
	masked := DisguiseDoubleQuotedSegments(input)
	expected := `pa*($&x12#a~aqb*($&x12#a~br`
	if masked != expected {
		t.Errorf("Mask wrong. Got: <%v>, expected: <%v>", masked, expected)
	}
}

func TestDisguisingWhenOddNumberOfQuoteMarks(t *testing.T) {
	input := `p"a"q"r`
	masked := DisguiseDoubleQuotedSegments(input)
	expected := `paq"r`
	if masked != expected {
		t.Errorf("Mask wrong. Got: <%v>, expected: <%v>", masked, expected)
	}
}

// set of tests for reinstating

func TestUnDisguisingNoOp(t *testing.T) {
	input := `foo`
	unmasked := UnDisguise(input)
	expected := `foo`
	if unmasked != expected {
		t.Errorf("Un masing went wrong. Got: <%v>, expected: <%v>", unmasked, expected)
	}
}

func TestUnDisguisingActive(t *testing.T) {
	input := `foo*($&x12#a~bar*($&x12#a~baz`
	unmasked := UnDisguise(input)
	expected := `foo bar baz`
	if unmasked != expected {
		t.Errorf("Un masing went wrong. Got: <%v>, expected: <%v>", unmasked, expected)
	}
}