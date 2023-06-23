package receive_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/stretchr/testify/assert"
)

func TestDecodeIncomingStatusMessage(t *testing.T) {
	json := `{"cmd":"PRINT_STOP", "status":"PRINT_STOP"}`

	data, err := receive.UnmarshalIncomingStatusMessage([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "PRINT_STOP", data.Cmd)
	assert.Equal(t, "PRINT_STOP", data.Status)
}

func TestEncodeIncomingStatusMessage(t *testing.T) {
	data := receive.IncomingStatusMessage{MessageCommand: jsontypes.MessageCommand{Cmd: "PRINT_STOP"}, Status: "PRINT_STOP"}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"cmd":"PRINT_STOP","status":"PRINT_STOP"}`, string(json))
}
