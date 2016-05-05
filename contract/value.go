package contract

import (
	"github.com/peterhoward42/frost/parse"
	"encoding/json"
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
}

type BoolValue struct {
	BoolValue bool
}

// The NewXXXValue() function is a factory that makes an instance of one of the XXXValue type
// objects above, depending the type implied by the incoming string representation. It
// then returns a JsonType interface which points to the newly created structure.
func NewXXXValue(inputString string) interface{} {
	var matched bool;

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
	return StringValue{StringValue: inputString}
}

// This set of MarshalJSON() methods implement the json.Marshaler interface, and we use them
// to downgrade what is output as JSON to the JSON that would be output for underlying native
// value held within, rather than wrapping it inside an object.
func (v FloatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.FloatValue)
}

func (v IntegerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.IntValue)
}

func (v StringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.StringValue)
}

func (v BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.BoolValue)
}
