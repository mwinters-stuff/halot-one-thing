// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    OutgoingMessage, err := UnmarshalOutgoingMessage(bytes)
//    bytes, err = OutgoingMessage.Marshal()

package receive

import "encoding/json"

func UnmarshalVersion(data []byte) (Version, error) {
	var r Version
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Version) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Version struct {
	Version string `json:"version"`
}
