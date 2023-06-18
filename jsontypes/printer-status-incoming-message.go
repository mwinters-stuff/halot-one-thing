// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    printStatus, err := UnmarshalPrintStatus(bytes)
//    bytes, err = printStatus.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalPrintStatusIncoming(data []byte) (PrintStatusIncoming, error) {
	var r PrintStatusIncoming
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PrintStatusIncoming) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PrintStatusIncoming struct {
	BottomExposureNum string `json:"bottomExposureNum"`
	Cmd               string `json:"cmd"`
	CurSliceLayer     string `json:"curSliceLayer"`
	DelayLight        string `json:"delayLight"`
	EleSpeed          string `json:"eleSpeed"`
	Filename          string `json:"filename"`
	InitExposure      string `json:"initExposure"`
	LayerThickness    string `json:"layerThickness"`
	PrintExposure     string `json:"printExposure"`
	PrintHeight       string `json:"printHeight"`
	PrintRemainTime   string `json:"printRemainTime"`
	PrintStatus       string `json:"printStatus"`
	Resin             string `json:"resin"`
	SliceLayerCount   string `json:"sliceLayerCount"`
}
