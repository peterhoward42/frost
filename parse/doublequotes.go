package parse

import (
	"regexp"
	"strings"
)

var doubleQuoteRe = regexp.MustCompile(`"[^"]*"`)

const almostSentinelString = `9975^-)zz#~foo` // Improbable to encounter by chance.

// The MaskDoubleQuotes() function returns a modified version of the given string, in which
// all double quoted strings have been replaced with non-quoted alternatives, inside which any
// spaces have been replaced by an (almost) sentinel value.
func MaskDoubleQuotes(inputString string) string {
	return doubleQuoteRe.ReplaceAllStringFunc(inputString, replacer)
}

// The UnMaskDoubleQuotes() is a sort of reciprocal to the MaskDoubleQuotes() function, that
// returns a a copy of the input string that is modified only when it contains one or more of
// the almost-sentinel strings.
func UnMaskDoubleQuotes(inputString string) string {
	if strings.Contains(inputString, almostSentinelString) == false {
		return inputString
	}
	s := strings.Replace(inputString, almostSentinelString, " ", -1)
	return `"` + s + `"`
}

func replacer(inputStr string) string {
	a := strings.TrimLeft(inputStr, `"`)
	b := strings.TrimRight(a, `"`)
	c := strings.Replace(b, " ", almostSentinelString, -1)
	return c
}
