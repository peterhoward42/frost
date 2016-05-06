package contract

import (
	"strings"
	"regexp"
)

var startsWithDigitRe = regexp.MustCompile(`^[\d]`)
var digitsAtFrontRe = regexp.MustCompile(`^[\d]+`)
var nonDigitsAtFrontRe = regexp.MustCompile(`^[\D]+`)
var whitespacePresentRe = regexp.MustCompile(`[\s]`)
var digitsVsNonFlipFlopRe = map[bool]*regexp.Regexp {true: digitsAtFrontRe,false: nonDigitsAtFrontRe,}

// The CaptureTagsFromString() function looks for implied sub divisions inside the given string
// and returns them in sequence. For example "AAA_32" is split into the tags "AAA" and "32".
func CaptureTagsFromString(inputString string) (segments []string) {
	// Do nothing when whitespace is present
	if whitespacePresentRe.MatchString(inputString) {
		return
	}
	// If any of the usual-suspect delimiters (like underscore) produces at least two segments,
	// all of which are non empty, then we conclude that this scheme is the only one present,
	// and leave it at that.
	for _, delim := range([]string{`_`, `-`, `.`, `/`, `\`}) {
		segments = strings.Split(inputString, delim)
		if len(segments) >= 2 {
			allSegmentsWellFormed := true
			for _, segment := range(segments) {
				if len(segment) == 0 {
					allSegmentsWellFormed = false
				}
			}
			if allSegmentsWellFormed {
				return segments
			}
		}
	}
	// Otherwise we look for sequences of digits, and sequences of non-digits.
	startWithDigits :=  startsWithDigitRe.MatchString(inputString);
	return captureBasedOnTransitions(startWithDigits, inputString)
}

// Recursive helper for CaptureTagsFromString()
func captureBasedOnTransitions(startWithDigits bool, inputString string) (segments []string) {
	re := digitsVsNonFlipFlopRe[startWithDigits]
	// Either peel off the right sort of segment, or detect we've reached the end.
	segment := re.FindString(inputString)
	if segment == "" {
		return
	} else {
		segments = append(segments, segment)
		// Recurse to consume the remainder, toggling the regexp to look for.
		remainder := strings.Replace(inputString, segment, "", 1)
		flippedStartWithChoice := !startWithDigits
		return captureBasedOnTransitions(flippedStartWithChoice, remainder)
	}
}