package receive_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/stretchr/testify/assert"
)

func TestDecodeStartData(t *testing.T) {
	json := `{"cmd":"START_DATA","errorcode":0,"key":"19ebcf3c","received":"538539","size":"538539"}`

	data, err := receive.UnmarshalStartData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "START_DATA", data.Cmd)

	assert.Equal(t, int64(0), data.Errorcode)
	assert.Equal(t, "19ebcf3c", data.Key)
	assert.Equal(t, "538539", data.Received)
	assert.Equal(t, "538539", data.Size)

}

func TestEncodeStartData(t *testing.T) {
	data := receive.StartData{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "START_FILE"},
		Errorcode: 0,
		Key:       "19ebcf3c",
		Received:  "538539",
		Size:      "538539",
	}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"cmd":"START_DATA","errorcode":0,"key":"19ebcf3c","received":"538539","size":"538539"}`, string(json))
}
