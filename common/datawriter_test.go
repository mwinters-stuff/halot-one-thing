package common_test

import (
	"testing"

	"github.com/mwinters-stuff/halo-one-thing/common"
	"github.com/stretchr/testify/assert"
)

func TestDataWriter(t *testing.T) {

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

	writer := common.NewDataWriter()
	writer.WriteU8(0xff)
	writer.WriteU16(0x1122)
	writer.WriteU32(0x22334455)
	writer.WriteString("astr", false)
	writer.WriteString("astr", true)
	writer.WriteBytes([]byte{0x00, 0x00, 0x22, 0x33, 0x44, 0x55})

	assert.Equal(t, data, writer.Bytes())

}
