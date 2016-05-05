package contract

import (
	"testing"
	"encoding/json"
	"reflect"
	"fmt"
)

func TestValueCreationWithJsonOutputForAllTypes(t *testing.T) {
	inputStrings := []string{"42", "3.14", "true", "hello"}
	typesExpected := []string{"contract.IntegerValue", "contract.FloatValue",
		"contract.BoolValue", "contract.StringValue"}
	jsonExpected := []string{"42", "3.14", "true", `"hello"`}

	for index, inputString := range (inputStrings) {
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
