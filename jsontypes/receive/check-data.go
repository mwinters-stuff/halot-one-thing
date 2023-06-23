// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    checkData, err := UnmarshalCheckData(bytes)
//    bytes, err = checkData.Marshal()

package receive

import (
	"encoding/json"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
)

func UnmarshalCheckData(data []byte) (CheckData, error) {
	var r CheckData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CheckData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CheckData struct {
	Checkstate int64 `json:"checkstate"`
	jsontypes.MessageCommand
	Key string `json:"key"`
}
