// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    OutgoingMessage, err := UnmarshalOutgoingMessage(bytes)
//    bytes, err = OutgoingMessage.Marshal()

package send

import (
	"encoding/json"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
)

func UnmarshalTokenMessage(data []byte) (TokenMessage, error) {
	var r TokenMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TokenMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TokenMessage struct {
	Token string `json:"token"`
	jsontypes.MessageCommand
}
