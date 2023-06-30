package receive_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/stretchr/testify/assert"
)

func TestDecodeStartPrintStatus(t *testing.T) {
	json := `{"cmd":"START_PRINT","filename":"Frog.cxdlp","status":"CHECKING"}`

	data, err := receive.UnmarshalStartPrintStatus([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "START_PRINT", data.Cmd)

	assert.Equal(t, "Frog.cxdlp", data.Filename)
	assert.Equal(t, "CHECKING", data.Status)

}

func TestEncodeStartPrintStatus(t *testing.T) {
	data := receive.StartPrintStatus{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "START_PRINT"},
		Filename: "Frog.cxdlp",
		Status:   "CHECKING",
	}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"cmd":"START_PRINT","filename":"Frog.cxdlp","status":"CHECKING"}`, string(json))
}
