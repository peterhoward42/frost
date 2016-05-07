package contract

type RowOfValues struct {
	Type string
	Values []interface{} // E.g. FloatValue
}

func NewRowOfValues(values []interface{}) *RowOfValues {
	return &RowOfValues {
		Type: "RowOfValues",
		Values: values,
	}
}