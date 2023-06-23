package receive_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/stretchr/testify/assert"
)

func TestDecodeVersion(t *testing.T) {
	json := `{"version":"1.0.0"}`

	data, err := receive.UnmarshalVersion([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "1.0.0", data.Version)
}

func TestEncodeVersion(t *testing.T) {
	data := receive.Version{Version: "1.0.0"}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"version":"1.0.0"}`, string(json))
}
