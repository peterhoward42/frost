package contract

import "encoding/json"

type ValueAlone struct {
	Json []byte
}

func NewValueAlone(stringValue string) *ValueAlone {
	marshalled, _ := json.Marshal(stringValue)
	return &ValueAlone{
		Json: marshalled,
	}
}
