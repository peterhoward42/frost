package contract

import (
	"parse"
)

// The IntegerValue structure wraps a native integer value, so as to elaborate it with properties
// that are required in the context of the frost contract package.
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

// The NewValueAlone() function is a factory that makes an instance of one of the XXXValue type
// objects above, depending the type implied by the incoming string representation. It
// then returns a JsonType interface which points to the newly created structure.
func NewValueAlone(inputString string) JsonType {
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
