package send_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
	"github.com/stretchr/testify/assert"
)

func TestDecodeStartFile(t *testing.T) {
	json := `{"token":"ATOKEN","cmd":"START_FILE","filename":"filename.abc","offset":"0","size":"212345"}`

	data, err := send.UnmarshalStartFile([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "START_FILE", data.Cmd)
	assert.Equal(t, "ATOKEN", data.Token)
	assert.Equal(t, "filename.abc", data.Filename)
	assert.Equal(t, "0", data.Offset)
	assert.Equal(t, "212345", data.Size)
}

func TestEncodeStartFile(t *testing.T) {
	data := send.StartFile{
		TokenMessage: send.TokenMessage{
			MessageCommand: jsontypes.MessageCommand{
				Cmd: "START_FILE"},
			Token: "ATOKEN"},
		Filename: "filename.abc",
		Offset:   "0",
		Size:     "212345"}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"token":"ATOKEN","cmd":"START_FILE","filename":"filename.abc","offset":"0","size":"212345"}`, string(json))
}
