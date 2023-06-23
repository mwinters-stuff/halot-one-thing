package receive_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/stretchr/testify/assert"
)

func TestDecodeCheckData(t *testing.T) {
	json := `{"checkstate":1,"cmd":"CHECK_DATA","key":"19ebcf3c"}`

	data, err := receive.UnmarshalCheckData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "CHECK_DATA", data.Cmd)

	assert.Equal(t, int64(1), data.Checkstate)
	assert.Equal(t, "19ebcf3c", data.Key)

}

func TestEncodeCheckData(t *testing.T) {
	data := receive.CheckData{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "CHECK_DATA"},
		Checkstate: 1,
		Key:        "19ebcf3c",
	}

	json, err := data.Marshal()
	assert.Nil(t, err, "Err is not nil")
	assert.Equal(t, `{"checkstate":1,"cmd":"CHECK_DATA","key":"19ebcf3c"}`, string(json))
}
