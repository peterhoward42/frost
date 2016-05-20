package testutils

import (
	"regexp"
	"testing"
)

// The AssertPageContainsSamples() function provides a shorthand for testing a page of
// text for the presence of a series of regular expressions in one go, and encapsulates the
// test error reporting when a failure occurs. It is dumb about line breaks, and it is left
// to the caller of the function to decide how to reconcile the regular expressions provided
// with the page provided. You might for example choose to replace newlines in the page with
// spaces before passing the page in.
func AssertPageContainsSamples(t *testing.T, page string, matchAssertions []*MatchAssertion) {
	for _, matchAssertion := range matchAssertions {
		reStr := matchAssertion.MustExistRegexp
		re := regexp.MustCompile(reStr)
		if re.MatchString(page) == false {
			t.Errorf("Page does not contain sample regexp.\n"+
				"Regexp: %v\nDescription: %v",
				reStr, matchAssertion.MetaDescription)
		}
	}
}

// The MatchAssertion type models a regular expression that can be used in a regexp matching
// test, along with a human readable string fragment that can be used to augment error handling.
type MatchAssertion struct {
	MustExistRegexp string
	MetaDescription string
}
