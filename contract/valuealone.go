package contract

import (
	"reflect"
	"strconv"
)

type IntegerValue struct {
	Value int64
}

type FloatValue struct {
	Value float64
}

type StringValue struct {
	Value string
}

type BoolValue struct {
	Value bool
}

func NewValueType(field *Field) JsonType {
	switch {
	case field.fieldType == reflect.Int:
		value, _ := strconv.ParseInt(field.stringValue, 0, 64)
		return IntegerValue{Value: value}
	case field.fieldType == reflect.Float64:
		value, _ := strconv.ParseFloat(field.stringValue, 64)
		return FloatValue{Value: value}
	case field.fieldType == reflect.Bool:
		value, _ := strconv.ParseBool(field.stringValue)
		return BoolValue{Value: value}
	case field.fieldType == reflect.String:
		return StringValue{Value: field.stringValue}
	}
	return nil
}
