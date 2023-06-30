// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    startPrint, err := UnmarshalStartPrint(bytes)
//    bytes, err = startPrint.Marshal()

package send

import "encoding/json"

func UnmarshalStartPrint(data []byte) (StartPrint, error) {
	var r StartPrint
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StartPrint) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StartPrint struct {
	TokenMessage
	Filename string `json:"filename"`
}
