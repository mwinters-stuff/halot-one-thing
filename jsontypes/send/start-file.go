// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    StartFileMessage, err := UnmarshalStartFileMessage(bytes)
//    bytes, err = StartFileMessage.Marshal()

package send

import "encoding/json"

func UnmarshalStartFile(data []byte) (StartFile, error) {
	var r StartFile
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StartFile) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StartFile struct {
	TokenMessage
	Filename string `json:"filename"`
	Offset   string `json:"offset"`
	Size     string `json:"size"`
}
