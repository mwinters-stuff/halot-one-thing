package receive_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/stretchr/testify/assert"
)

func TestDecodeStartFile(t *testing.T) {
	json := `{"cmd":"START_FILE","compress":true,"filename":"Frog.cxdlp","key":"19ebcf3c","offset":"0","size":"538539"}`

	data, err := receive.UnmarshalStartFile([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "START_FILE", data.Cmd)

	assert.True(t, data.Compress)
	assert.Equal(t, "Frog.cxdlp", data.Filename)
	assert.Equal(t, "19ebcf3c", data.Key)
	assert.Equal(t, "0", data.Offset)
	assert.Equal(t, "538539", data.Size)

}

func TestEncodeStartFile(t *testing.T) {
	data := receive.StartFile{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "START_FILE"},
		Compress: true,
		Filename: "Frog.cxdlp",
		Key:      "19ebcf3c",
		Offset:   "0",
		Size:     "538539",
	}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"cmd":"START_FILE","compress":true,"filename":"Frog.cxdlp","key":"19ebcf3c","offset":"0","size":"538539"}`, string(json))
}
