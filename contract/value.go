package contract

import (
	"encoding/json"
	"github.com/peterhoward42/frost/parse"
)

// This set of XXXValue structures wrap a native integer value, so as to elaborate it with
// properties that are required in the context of the frost contract package.
type IntegerValue struct {
	IntValue int
}

type FloatValue struct {
	FloatValue float64
}

type StringValue struct {
	StringValue string
	Tags []string
}

type BoolValue struct {
	BoolValue bool
}

// The NewXXXValue() function is a factory that makes an instance of one of the XXXValue type
// objects above, depending the type implied by the incoming string representation. It
// then returns a JsonType interface which points to the newly created structure.
func NewXXXValue(inputString string) interface{} {
	var matched bool

	matched, iVal := parse.LooksLikeAnInteger(inputString)
	if matched {
		return IntegerValue{IntValue: iVal}
	}

	matched, fVal := parse.LooksLikeAFloat(inputString)
	if matched {
		return FloatValue{FloatValue: fVal}
	}

	matched, bVal := parse.LooksLikeABool(inputString)
	if matched {
		return BoolValue{BoolValue: bVal}
	}

	// Catch all is to return a string value
	return StringValue{StringValue:inputString, Tags: NewStringTags(inputString)}
}

// We override this interface for the numeric variants to downgrade the JSON created from
// that of an object to that of a raw type. We don't do it for the string variant because we
// want that one to also include the tags.
func (v FloatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.FloatValue)
}

func (v IntegerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.IntValue)
}

func (v BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.BoolValue)
}
