// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    startPrintStatus, err := UnmarshalStartPrintStatus(bytes)
//    bytes, err = startPrintStatus.Marshal()

package receive

import (
	"encoding/json"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
)

func UnmarshalStartPrintStatus(data []byte) (StartPrintStatus, error) {
	var r StartPrintStatus
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StartPrintStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StartPrintStatus struct {
	jsontypes.MessageCommand
	Filename string `json:"filename"`
	Status   string `json:"status"`
}
