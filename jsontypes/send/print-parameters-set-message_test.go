package send_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
	"github.com/stretchr/testify/assert"
)

func TestDecodePrintParametersSet(t *testing.T) {
	json := `{"bottomExposureNum":"2","cmd":"PRINT_PARA_SET","delayLight":"4","eleSpeed":"2","initExposure":"40","printExposure":"3","printHeight":"6","token":"ATOKEN"}`

	data, err := send.UnmarshalPrintParametersSetMessage([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "PRINT_PARA_SET", data.Cmd)

	assert.Equal(t, "2", data.BottomExposureNum)
	assert.Equal(t, "4", data.DelayLight)
	assert.Equal(t, "2", data.EleSpeed)
	assert.Equal(t, "40", data.InitExposure)
	assert.Equal(t, "3", data.PrintExposure)
	assert.Equal(t, "6", data.PrintHeight)
	assert.Equal(t, "ATOKEN", data.Token)

}

func TestEncodePrintParametersSet(t *testing.T) {
	data := send.PrintParametersSet{
		OutgoingMessage: send.OutgoingMessage{
			MessageCommand: jsontypes.MessageCommand{
				Cmd: "PRINT_PARA_SET"},
			Token: "ATOKEN"},
		BottomExposureNum: "2",
		DelayLight:        "4",
		EleSpeed:          "2",
		InitExposure:      "40",
		PrintExposure:     "3",
		PrintHeight:       "6",
	}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"bottomExposureNum":"2","token":"ATOKEN","cmd":"PRINT_PARA_SET","delayLight":"4","eleSpeed":"2","initExposure":"40","printExposure":"3","printHeight":"6"}`, string(json))
}
