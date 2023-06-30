package common_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/common"
	"github.com/stretchr/testify/assert"
)

func TestDataReader(t *testing.T) {

	data := []byte{
		0xff,       // uint8
		0x11, 0x22, // unit16
		0x22, 0x33, 0x44, 0x55, //unit32
		0x00, 0x00, 0x00, 0x05, // uint32 string length
		'a', 's', 't', 'r', 0x00, // string
		0x00, 0x00, 0x00, 0x0a, // uint32 string length
		0x00, 'a', 0x00, 's', 0x00, 't', 0x00, 'r', 0x00, 0x00, // string
		0x00, 0x00, // to skip
		0x22, 0x33, 0x44, 0x55, // bytes
	}

	reader := common.NewDataReader(data, 0)

	assert.Equal(t, uint8(255), reader.ReadU8())
	assert.Equal(t, uint16(0x1122), reader.ReadU16())
	assert.Equal(t, uint32(0x22334455), reader.ReadU32())
	assert.Equal(t, "astr", reader.ReadString(false))
	assert.Equal(t, "astr", reader.ReadString(true))
	reader.Skip(2)
	assert.Equal(t, []byte{0x22, 0x33, 0x44, 0x55}, reader.ReadBytes(4))
	assert.Equal(t, 36, reader.Offset())
}
