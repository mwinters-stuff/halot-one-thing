// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    StartFileMessage, err := UnmarshalStartFileMessage(bytes)
//    bytes, err = StartFileMessage.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalStartFileMessage(data []byte) (StartFileMessage, error) {
	var r StartFileMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StartFileMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StartFileMessage struct {
	Cmd      string `json:"cmd"`
	Token    string `json:"token"`
	Filename string `json:"filename"`
	Offset   string `json:"offset"`
	Size     string `json:"size"`
}
