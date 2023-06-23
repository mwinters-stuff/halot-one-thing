package send_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
	"github.com/stretchr/testify/assert"
)

func TestDecodeOutgoingMessage(t *testing.T) {
	json := `{"cmd":"RUN_ME", "Token":"ATOKEN"}`

	data, err := send.UnmarshalOutgoingMessage([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "RUN_ME", data.Cmd)
	assert.Equal(t, "ATOKEN", data.Token)
}

func TestEncodeOutgoingMessage(t *testing.T) {
	data := send.OutgoingMessage{MessageCommand: jsontypes.MessageCommand{Cmd: "RUN_ME"}, Token: "ATOKEN"}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"token":"ATOKEN","cmd":"RUN_ME"}`, string(json))
}
