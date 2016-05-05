package parse

import (
	"regexp"
	"strings"
)

var doubleQuoteRe = regexp.MustCompile(`"[^"]*"`)

const almostSentinelString = `*($&x12#a~` // Improbable to encounter by chance.

// The DisguiseDoubleQuotedSegments() function returns a modified version of the given string, in
// which any internal double quoted strings are replaced with non-quoted variants in which any 
// internal spaces have been replaced with the almost-sentinel string above.
func DisguiseDoubleQuotedSegments(inputString string) string {
	return doubleQuoteRe.ReplaceAllStringFunc(inputString, replacer)
}

// The UnDisguise() method returns a variant of the input string in which any instances of
// the almost-sentinel string above are replaced with a space.
func UnDisguise(inputString string) string {
	if strings.Contains(inputString, almostSentinelString) == false {
		return inputString
	}
	return strings.Replace(inputString, almostSentinelString, " ", -1)
}

func replacer(inputStr string) string {
	a := strings.TrimLeft(inputStr, `"`)
	b := strings.TrimRight(a, `"`)
	c := strings.Replace(b, " ", almostSentinelString, -1)
	return c
}
