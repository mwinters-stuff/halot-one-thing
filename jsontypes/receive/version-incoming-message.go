// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    OutgoingMessage, err := UnmarshalOutgoingMessage(bytes)
//    bytes, err = OutgoingMessage.Marshal()

package receive

import "encoding/json"

func UnmarshalVersionIncomingMessage(data []byte) (VersionIncoming, error) {
	var r VersionIncoming
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VersionIncoming) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type VersionIncoming struct {
	Version string `json:"version"`
}
