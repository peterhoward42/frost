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

type BoolValue struct {
	BoolValue bool
}

type StringValue struct {
	Value string
	Tags  []string
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
		Value: inputString,
		Tags:  CaptureTagsFromString(inputString),
	}
}

// We implement the HasFrostType interface for each type of value

func (v IntegerValue) GetFrostType() FrostType { return FrostInt }
func (v FloatValue) GetFrostType() FrostType   { return FrostFloat }
func (v BoolValue) GetFrostType() FrostType    { return FrostBool }
func (v StringValue) GetFrostType() FrostType  { return FrostString }

// We implement the MarshalJSON interface to customise the JSON produced.

func (v IntegerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.IntValue)
}

func (v FloatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.FloatValue)
}

func (v BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.BoolValue)
}

func (v StringValue) MarshalJSON() ([]byte, error) {
	type template struct {
		Type  string
		Value string
		Tags  []string
	}
	toMarshal := template{
		Type:  "string",
		Value: v.Value,
		Tags:  v.Tags,
	}
	return json.Marshal(toMarshal)
}
