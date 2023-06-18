// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    OutgoingMessage, err := UnmarshalOutgoingMessage(bytes)
//    bytes, err = OutgoingMessage.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalOutgoingMessage(data []byte) (OutgoingMessage, error) {
	var r OutgoingMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OutgoingMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OutgoingMessage struct {
	Cmd   string `json:"cmd"`
	Token string `json:"token"`
}
