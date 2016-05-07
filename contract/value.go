package contract

import (
	"encoding/json"
	"github.com/peterhoward42/frost/parse"
)

// This set of XXXValue structures wrap a native typed primitive value, so as to elaborate them with
// properties and behaviour that are required in the context of the frost contract package.
type IntegerValue struct {
	IntValue int
}

type FloatValue struct {
	FloatValue float64
}

type StringValue struct {
	Type  string
	Value string
	Tags  []string
}

type BoolValue struct {
	BoolValue bool
}

// The NewXXXValue() function is a polymorphic factory function that makes an instance of one
// of the XXXValue types above, depending on the type implied by the incoming string
// representation. It returns an empty interface object that points to the concrete type created.
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
	return StringValue{
		Type:  "String",
		Value: inputString,
		Tags:  CaptureTagsFromString(inputString)}
}

// We override these interface functions to customize the JSON produced

func (v FloatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.FloatValue)
}

func (v IntegerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.IntValue)
}

func (v BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.BoolValue)
}
