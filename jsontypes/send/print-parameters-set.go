// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    printParametersSetIncomingMessage, err := UnmarshalPrintParametersSetIncomingMessage(bytes)
//    bytes, err = printParametersSetIncomingMessage.Marshal()

package send

import "encoding/json"

func UnmarshalPrintParametersSet(data []byte) (PrintParametersSet, error) {
	var r PrintParametersSet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PrintParametersSet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PrintParametersSet struct {
	BottomExposureNum string `json:"bottomExposureNum"`
	TokenMessage
	DelayLight    string `json:"delayLight"`
	EleSpeed      string `json:"eleSpeed"`
	InitExposure  string `json:"initExposure"`
	PrintExposure string `json:"printExposure"`
	PrintHeight   string `json:"printHeight"`
}
