package send_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
	"github.com/stretchr/testify/assert"
)

func TestDecodeStartPrint(t *testing.T) {
	json := `{"cmd":"START_PRINT","filename":"Frog.cxdlp","token":"kz4M32rFVvI="}`

	data, err := send.UnmarshalStartPrint([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "START_PRINT", data.Cmd)
	assert.Equal(t, "Frog.cxdlp", data.Filename)
	assert.Equal(t, "kz4M32rFVvI=", data.Token)
}

func TestEncodeStartPrint(t *testing.T) {
	data := send.StartPrint{
		TokenMessage: send.TokenMessage{
			MessageCommand: jsontypes.MessageCommand{
				Cmd: "START_PRINT"},
			Token: "kz4M32rFVvI=",
		},
		Filename: "Frog.cxdlp",
	}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"token":"kz4M32rFVvI=","cmd":"START_PRINT","filename":"Frog.cxdlp"}`, string(json))
}
