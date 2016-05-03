package contract

import (
	"reflect"
	"strconv"
	"strings"
	"appengine"
)

type Field struct {
	// We capture the type during parsing but normalise the value held to a string
	// representation.
	fieldType   reflect.Kind
	stringValue string
}

func NewField(fromString string, ctx appengine.Context) *Field {
	// Tests have significant inline returns and precedence.
	_, err := strconv.ParseInt(fromString, 0, 64)
	if err != nil {
		ctx.Infof("XXX as int")
		return &Field{
			fieldType:   reflect.Int,
			stringValue: fromString,
		}
	}
	// Now as float
	_, err = strconv.ParseFloat(fromString, 64)
	if err != nil {
		ctx.Infof("XXX as float")
		return &Field{
			fieldType:   reflect.Float64,
			stringValue: fromString,
		}
	}
	// Now as bool
	if strings.ToLower(fromString) == "true" {
		ctx.Infof("XXX as bool")
		return &Field{
			fieldType:   reflect.Bool,
			stringValue: fromString,
		}
	}
	if strings.ToLower(fromString) == "false" {
		ctx.Infof("XXX as bool")
		return &Field{
			fieldType:   reflect.Bool,
			stringValue: fromString,
		}
	}
	// Default to string
	ctx.Infof("XXX as string")
	return &Field{
		fieldType:   reflect.String,
		stringValue: fromString,
	}
}
