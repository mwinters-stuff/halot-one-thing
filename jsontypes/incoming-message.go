// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    IncomingMessage, err := UnmarshalIncomingMessage(bytes)
//    bytes, err = IncomingMessage.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalIncomingMessage(data []byte) (IncomingMessage, error) {
	var r IncomingMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IncomingMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IncomingMessage struct {
	Cmd string `json:"cmd"`
}
