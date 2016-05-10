package contract

type KeyValuePair struct {
	Type  string
	Key   string
	Value interface{}
}

func NewKeyValuePair(key string, valueString string) *KeyValuePair {
	return &KeyValuePair{
		Type:  "KeyValue",
		Key:   key,
		Value: NewXXXValue(valueString),
	}
}
