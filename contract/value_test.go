package contract

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestValueCreationWithJsonOutputForNonStringTypes(t *testing.T) {
	inputStrings := []string{"42", "3.14", "true"}
	typesExpected := []string{"contract.IntegerValue", "contract.FloatValue",
		"contract.BoolValue"}
	jsonExpected := []string{"42", "3.14", "true"}

	for index, inputString := range inputStrings {
		v := NewXXXValue(inputString)
		typeCreated := fmt.Sprintf("%v", reflect.TypeOf(v))
		typeExpected := typesExpected[index]
		if typeCreated != typeExpected {
			t.Errorf("Type inferred wrong. Got: %v, but expected: %v", typeCreated,
				typeExpected)
		}
		json, _ := v.(json.Marshaler).MarshalJSON()
		jsonProduced := string(json)
		jsonExpected := jsonExpected[index]
		if jsonProduced != jsonExpected {
			t.Errorf("json output wrong. Got: <%v>, expected: <%v>", jsonProduced,
				jsonExpected)
		}
	}
}

func TestValueCreationForString(t *testing.T) {
	inputString := "ABB_1_OFF"

	// Ensure correct type created.
	v := NewXXXValue(inputString)
	typeCreated := fmt.Sprintf("%v", reflect.TypeOf(v))
	typeExpected := "contract.StringValue"
	if typeCreated != typeExpected {
		t.Errorf("Type inferred wrong. Got: %v, but expected: %v", typeCreated,
			typeExpected)
	}

	// Ensure there are some tags captured - but leave detailed tag testing to a separate
	// unit test.
	stringValue := v.(StringValue)
	numTagsFound := len(stringValue.Tags)
	numTagsExpected := 3;
	if numTagsFound != numTagsExpected {
		t.Errorf("Number of tags wrong. Got: %v, but expected: %v", numTagsFound,
			numTagsExpected)
	}
}
