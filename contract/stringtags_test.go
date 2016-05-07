package contract

import (
	"testing"
)

// First a bunch of special cases - to ensure these are working and not confusing the
// mainstream tests that follow.

func TestStringTagsSplitDoesNothingWhenASpaceIsPresent(t *testing.T) {
	segments := CaptureTagsFromString(`foo_bar baz`)
	if len(segments) != 0 {
		t.Errorf("Failed to do nothing when a space is present")
	}
}

func TestStringTagsSplitCopesWithEmptyString(t *testing.T) {
	segments := CaptureTagsFromString(``)
	if len(segments) != 0 {
		t.Errorf("Failed to cope with empty string.")
	}
}

func TestStringTagsSplitDoesNothingWithSingleCharacterInput(t *testing.T) {
	segments := CaptureTagsFromString(`a`)
	if len(segments) != 0 {
		t.Errorf("Failed to ignore single character string.")
	}
}

func TestStringTagsSplitAbandonsDelimiterSchemeIfProducesEmptySegments(t *testing.T) {
	segments := CaptureTagsFromString(`foo-bar-`)
	if len(segments) != 0 {
		t.Errorf("Failed to abandon delimiter scheme when produced empty segment.")
	}
}

func TestStringTagsSplitCopesWithDelimitersWithNothingBetween(t *testing.T) {
	segments := CaptureTagsFromString(`--`)
	if len(segments) != 0 {
		t.Errorf("Failed to cope with bunched delimiters.")
	}
}

// Now the mainstream tests.

func TestStringTagsSplitOnCustomaryDelimiters(t *testing.T) {
	// Repeat the same test for each delimiter in the FROST tagging contract.
	for _, delim := range []string{`_`, `-`, `.`, `/`, `\`} {
		segments := CaptureTagsFromString(`foo` + delim + `bar` + delim + `baz`)
		expected := []string{`foo`, `bar`, `baz`}

		// Right number?
		segmentsExpectedCount := len(expected)
		segmentsReceivedCount := len(segments)
		if segmentsExpectedCount != segmentsReceivedCount {
			t.Errorf("Number of tags is wrong. Got %v, expected: %v", segmentsReceivedCount,
				segmentsExpectedCount)
		}
		// Right content?
		for idx, tag := range expected {
			if tag != expected[idx] {
				t.Errorf("Unexpected tag: %v", tag)
			}
		}
	}
}

func TestStringTagsSplitOnWellFormedAlternativesDigitsFirst(t *testing.T) {
	input := "42foo9bar"
	tagsExpected := []string{"42", "foo", "9", "bar"}
	tagsFound := CaptureTagsFromString(input)
	for idx, tagFound := range(tagsFound) {
		expected := tagsExpected[idx]
		if tagFound != expected {
			t.Errorf("Unexpected tag. Got: %v, expected %v", tagFound, expected)
		}
	}
}

func TestStringTagsSplitOnWellFormedAlternativesNonDigitsFirst(t *testing.T) {
	input := "foo9bar42"
	tagsExpected := []string{"foo", "9", "bar", "42"}
	tagsFound := CaptureTagsFromString(input)
	for idx, tagFound := range(tagsFound) {
		expected := tagsExpected[idx]
		if tagFound != expected {
			t.Errorf("Unexpected tag. Got: %v, expected %v", tagFound, expected)
		}
	}
}