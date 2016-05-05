package contract

import "strings"

func NewStringTags(inputString string) []string {
	return strings.Split(inputString, "_")
}
