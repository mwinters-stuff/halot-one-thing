// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    IncomingMessage, err := UnmarshalIncomingMessage(bytes)
//    bytes, err = IncomingMessage.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalMessageCommand(data []byte) (MessageCommand, error) {
	var r MessageCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessageCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageCommand struct {
	Cmd string `json:"cmd"`
}
