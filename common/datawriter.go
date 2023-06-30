package common

import (
	"bytes"
	"encoding/binary"
)

// DataWriter represents a helper struct for writing binary data
type DataWriter interface {
	WriteU8(val uint8)
	WriteU16(val uint16)
	WriteU32(val uint32)
	WriteString(str string, isDoubleByte bool)
	WriteBytes([]byte)
	Bytes() []byte
}

type DataWriterImpl struct {
	buffer bytes.Buffer
}

// NewDataWriter creates a new instance of DataWriter
func NewDataWriter() DataWriter {
	return &DataWriterImpl{}
}

func (w *DataWriterImpl) WriteU8(val uint8) {
	w.buffer.WriteByte(val)
}

func (w *DataWriterImpl) WriteU16(val uint16) {
	binary.Write(&w.buffer, binary.BigEndian, val)
}

func (w *DataWriterImpl) WriteU32(val uint32) {
	binary.Write(&w.buffer, binary.BigEndian, val)
}

func (w *DataWriterImpl) WriteString(str string, isDoubleByte bool) {
	if isDoubleByte {
		w.WriteU32(uint32((len(str) + 1) * 2))
		b := []byte(str)
		for i := 0; i < len(b); i++ {
			w.buffer.WriteByte(0x00)
			w.buffer.WriteByte(b[i])
		}
		w.buffer.WriteByte(0x00)
		w.buffer.WriteByte(0x00)
	} else {
		w.WriteU32(uint32(len(str) + 1))
		w.buffer.WriteString(str)
		w.buffer.WriteByte(0x00)
	}
}

func (w *DataWriterImpl) WriteBytes(data []byte) {
	w.buffer.Write(data)
}

func (w *DataWriterImpl) Bytes() []byte {
	return w.buffer.Bytes()
}
