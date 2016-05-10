package contract

type RowOfValues struct {
	Type   string
	Values []interface{} // E.g. FloatValue
}

func NewRowOfValues(valueStrings []string) *RowOfValues {
	rowOfValues := RowOfValues{
		Type:   "RowOfValues",
		Values: []interface{}{},
	}
	for _, valueString := range valueStrings {
		rowOfValues.Values = append(rowOfValues.Values, NewXXXValue(valueString))
	}
	return &rowOfValues
}
