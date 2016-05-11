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

func (row *RowOfValues) HasSameSignatureAs(otherRow *RowOfValues) bool {
	if len(row.Values) != len(otherRow.Values) {
		return false
	}
	for i := 0; i < len(row.Values); i++ {
		rowSig := row.Values[i].(HasFrostType).GetFrostType();
		otherRowSig := otherRow.Values[i].(HasFrostType).GetFrostType();
		if rowSig != otherRowSig {
			return false
		}
	}
	return true
}