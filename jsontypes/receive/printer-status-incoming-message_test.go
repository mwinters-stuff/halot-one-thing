package receive_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/stretchr/testify/assert"
)

func TestDecodePrinterStatusIncoming(t *testing.T) {
	json := `{"bottomExposureNum":"2","cmd":"GET_PRINT_STATUS","curSliceLayer":"0","delayLight":"1","eleSpeed":"1","filename":"Frog.cxdlp","initExposure":"40","layerThickness":"0.050000","printExposure":"3","printHeight":"6","printRemainTime":"790","printStatus":"PRINT_STOP","resin":"","sliceLayerCount":"43"}`

	data, err := receive.UnmarshalPrintStatusIncoming([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "GET_PRINT_STATUS", data.Cmd)

	assert.Equal(t, "2", data.BottomExposureNum)
	assert.Equal(t, "0", data.CurSliceLayer)
	assert.Equal(t, "1", data.DelayLight)
	assert.Equal(t, "1", data.EleSpeed)
	assert.Equal(t, "Frog.cxdlp", data.Filename)
	assert.Equal(t, "40", data.InitExposure)
	assert.Equal(t, "0.050000", data.LayerThickness)
	assert.Equal(t, "3", data.PrintExposure)
	assert.Equal(t, "6", data.PrintHeight)
	assert.Equal(t, "790", data.PrintRemainTime)
	assert.Equal(t, "PRINT_STOP", data.PrintStatus)
	assert.Equal(t, "", data.Resin)
	assert.Equal(t, "43", data.SliceLayerCount)

}

func TestEncodePrinterStatusIncoming(t *testing.T) {
	data := receive.PrintStatusIncoming{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "GET_PRINT_STATUS"},
		BottomExposureNum: "2",
		CurSliceLayer:     "0",
		DelayLight:        "1",
		EleSpeed:          "1",
		Filename:          "Frog.cxdlp",
		InitExposure:      "40",
		LayerThickness:    "0.050000",
		PrintExposure:     "3",
		PrintHeight:       "6",
		PrintRemainTime:   "790",
		PrintStatus:       "PRINT_STOP",
		Resin:             "",
		SliceLayerCount:   "43",
	}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"bottomExposureNum":"2","cmd":"GET_PRINT_STATUS","curSliceLayer":"0","delayLight":"1","eleSpeed":"1","filename":"Frog.cxdlp","initExposure":"40","layerThickness":"0.050000","printExposure":"3","printHeight":"6","printRemainTime":"790","printStatus":"PRINT_STOP","resin":"","sliceLayerCount":"43"}`, string(json))
}
