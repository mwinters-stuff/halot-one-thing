package jsontypes_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/stretchr/testify/assert"
)

func TestDecodeMessageCommand(t *testing.T) {
	json := `{"cmd":"RUN_ME"}`

	data, err := jsontypes.UnmarshalMessageCommand([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "RUN_ME", data.Cmd)
}

func TestEncodeMessageCommand(t *testing.T) {
	data := jsontypes.MessageCommand{Cmd: "RUN_ME"}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"cmd":"RUN_ME"}`, string(json))
}
