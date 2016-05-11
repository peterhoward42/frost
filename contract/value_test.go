package contract

import (
	"encoding/json"
	"testing"
)

func TestGetFrostTypeWorks(t *testing.T) {
	if NewXXXValue("42").(HasFrostType).GetFrostType() != FrostInt {
		t.Errorf("Wrong type")
	}
	if NewXXXValue("42.3").(HasFrostType).GetFrostType() != FrostFloat {
		t.Errorf("Wrong type")
	}
	if NewXXXValue("true").(HasFrostType).GetFrostType() != FrostBool {
		t.Errorf("Wrong type")
	}
	if NewXXXValue("foo").(HasFrostType).GetFrostType() != FrostString {
		t.Errorf("Wrong type")
	}
}

func TestValueCreationWithJsonOutputForNonStringTypes(t *testing.T) {
	inputStrings := []string{"42", "3.14", "true"}
	jsonExpected := []string{"42", "3.14", "true"}

	for index, inputString := range inputStrings {
		v := NewXXXValue(inputString)
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
	v := NewXXXValue(inputString)
	// Ensure there are some tags captured - but leave detailed tag testing to a separate
	// unit test.
	stringValue := v.(StringValue)
	numTagsFound := len(stringValue.Tags)
	numTagsExpected := 3
	if numTagsFound != numTagsExpected {
		t.Errorf("Number of tags wrong. Got: %v, but expected: %v", numTagsFound,
			numTagsExpected)
	}
}
