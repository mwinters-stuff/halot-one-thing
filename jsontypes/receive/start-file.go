// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    startFile, err := UnmarshalStartFile(bytes)
//    bytes, err = startFile.Marshal()

package receive

import (
	"encoding/json"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
)

func UnmarshalStartFile(data []byte) (StartFile, error) {
	var r StartFile
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StartFile) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StartFile struct {
	jsontypes.MessageCommand
	Compress bool   `json:"compress"`
	Filename string `json:"filename"`
	Key      string `json:"key"`
	Offset   string `json:"offset"`
	Size     string `json:"size"`
}
