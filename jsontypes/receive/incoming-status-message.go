// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    OutgoingMessage, err := UnmarshalOutgoingMessage(bytes)
//    bytes, err = OutgoingMessage.Marshal()

package receive

import (
	"encoding/json"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
)

func UnmarshalIncomingStatusMessage(data []byte) (IncomingStatusMessage, error) {
	var r IncomingStatusMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IncomingStatusMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IncomingStatusMessage struct {
	jsontypes.MessageCommand
	Status string `json:"status"`
}
