package send_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
	"github.com/stretchr/testify/assert"
)

func TestDecodeTokenMessage(t *testing.T) {
	json := `{"cmd":"RUN_ME", "Token":"ATOKEN"}`

	data, err := send.UnmarshalTokenMessage([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "RUN_ME", data.Cmd)
	assert.Equal(t, "ATOKEN", data.Token)
}

func TestEncodeTokenMessage(t *testing.T) {
	data := send.TokenMessage{MessageCommand: jsontypes.MessageCommand{Cmd: "RUN_ME"}, Token: "ATOKEN"}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"token":"ATOKEN","cmd":"RUN_ME"}`, string(json))
}
