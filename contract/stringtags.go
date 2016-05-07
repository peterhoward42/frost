package contract

import (
	"regexp"
	"strings"
)

var startsWithDigitRe = regexp.MustCompile(`^[\d]`)
var digitsAtFrontRe = regexp.MustCompile(`^[\d]+`)
var nonDigitsAtFrontRe = regexp.MustCompile(`^[\D]+`)
var whitespacePresentRe = regexp.MustCompile(`[\s]`)
var digitsVsNonFlipFlopRe = map[bool]*regexp.Regexp{
	true:  digitsAtFrontRe,
	false: nonDigitsAtFrontRe,
}

// The CaptureTagsFromString() function looks for implied sub divisions inside the given string
// and returns them in sequence. For example "AAA_32" is split into the tags "AAA" and "32".
func CaptureTagsFromString(inputString string) []string {
	// Do nothing when whitespace is present
	if whitespacePresentRe.MatchString(inputString) {
		return []string{}
	}
	// If any of the usual-suspect delimiters (like underscore) produces at least two segments,
	// all of which are non empty, then we conclude that this scheme is the only one present,
	// and leave it at that.
	for _, delim := range []string{`_`, `-`, `.`, `/`, `\`} {
		tags := strings.Split(inputString, delim)
		if len(tags) >= 2 {
			allTagsWellFormed := true
			for _, tag := range tags {
				if len(tag) == 0 {
					allTagsWellFormed = false
				}
			}
			if allTagsWellFormed {
				return tags
			}
		}
	}
	// Otherwise we split the input string into a sequence of digit and non digit segments,
	// and accept these as tags, provided there are at least two.
	startWithDigits := startsWithDigitRe.MatchString(inputString)
	tagsFromTransitions := tagsFromTransitions(startWithDigits, inputString)
	if len(tagsFromTransitions) >= 2 {
		return tagsFromTransitions
	} else {
		return []string{}
	}
}

// Recursive helper function.
func tagsFromTransitions(startWithDigits bool, inputString string) []string {
	// Choose the digit or non-digit regexp according to the mandate passed in.
	re := digitsVsNonFlipFlopRe[startWithDigits]
	// Either peel off a matching segment, or detect we've reached the end.
	tag := re.FindString(inputString)
	if tag == "" {
		return []string{}
	} else {
		tagsToReturn := []string{tag}
		// Recurse to extract tags from the remainder of the input string, toggling the
		// regexp of what to look for first.
		remainderOfInputString := strings.Replace(inputString, tag, "", 1)
		flippedStartWithChoice := !startWithDigits
		tailTags := tagsFromTransitions(flippedStartWithChoice, remainderOfInputString)
		tagsToReturn = append(tagsToReturn, tailTags...)
		return tagsToReturn
	}
}
