package parse

import (
	"regexp"
	"strconv"
	"strings"
)

var floatRe = regexp.MustCompile(`^[+-]?[\d]*\.[\d]+$`)
var intRe = regexp.MustCompile(`^[+-]*[\d]+$`)
var boolRe = regexp.MustCompile(`(?i)^(true|false)$`) // (i) means case-insensitive
var keyStringRe = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)

// The LooksLikeAnInteger() function assesses if the input string in its entirety looks like
// a well formed integer, and when so, also provides the converted value.
func LooksLikeAnInteger(inputStr string) (matched bool, value int) {
	matched = intRe.MatchString(inputStr)
	if matched {
		value, _ = strconv.Atoi(inputStr)
		return true, value
	}
	return
}

// The LooksLikeAFloat() function assesses if the input string in its entirety looks like
// a well formed float, and when so, also provides the converted value.
func LooksLikeAFloat(inputStr string) (matched bool, value float64) {
	matched = floatRe.MatchString(inputStr)
	if matched {
		value, _ = strconv.ParseFloat(inputStr, 64)
		return true, value
	}
	return
}

// The LooksLikeABool() function assesses if the input string in its entirety looks like
// a well formed bool, and when so, also provides the converted value.
func LooksLikeABool(inputStr string) (matched bool, value bool) {
	matched = boolRe.MatchString(inputStr)
	if matched {
		// strconv.ParseBool() is too willing for our needs, but we only let it loose
		// once our own Regexp condition is met.
		value, _ = strconv.ParseBool(strings.ToLower(inputStr))
		return true, value
	}
	return
}

// The LooksLikeAKeyString() function assesses if the input string in its entirety looks like
// a string that is suitable to use as a key. I.e. rather like the rules for an identifier func init() {
// a programming language.
func LooksLikeAKeyString(inputStr string) (matched bool) {
	return keyStringRe.MatchString(inputStr)
}
