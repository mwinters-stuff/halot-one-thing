// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    startData, err := UnmarshalStartData(bytes)
//    bytes, err = startData.Marshal()

package receive

import (
	"encoding/json"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
)

func UnmarshalStartData(data []byte) (StartData, error) {
	var r StartData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StartData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StartData struct {
	jsontypes.MessageCommand
	Errorcode int64  `json:"errorcode"`
	Key       string `json:"key"`
	Received  string `json:"received"`
	Size      string `json:"size"`
}
