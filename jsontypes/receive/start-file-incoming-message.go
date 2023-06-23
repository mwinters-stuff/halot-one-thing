// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    StartFileIncoming, err := UnmarshalStartFileIncoming(bytes)
//    bytes, err = StartFileIncoming.Marshal()

package receive

import "encoding/json"

func UnmarshalStartFileIncoming(data []byte) (StartFileIncoming, error) {
	var r StartFileIncoming
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StartFileIncoming) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StartFileIncoming struct {
	Cmd       string `json:"cmd"`
	Compress  bool   `json:"compress"`
	Errorcode int64  `json:"errorcode"`
	Filename  string `json:"filename"`
	Key       string `json:"key"`
	Offset    string `json:"offset"`
	Size      string `json:"size"`
}
