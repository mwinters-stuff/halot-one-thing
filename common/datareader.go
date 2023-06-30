package common

import "encoding/binary"

// DataReader represents a helper struct for reading binary data
type DataReader interface {
	ReadU8() uint8
	ReadU16() uint16
	ReadU32() uint32
	ReadString(doubleByte bool) string
	ReadBytes(length int) []byte
	Skip(length int)
	Offset() int
}

type DataReaderImpl struct {
	data   []byte
	offset int
}

// NewDataReader creates a new instance of DataReader
func NewDataReader(data []byte, offset int) DataReader {
	return &DataReaderImpl{
		data:   data,
		offset: offset,
	}
}

func (r *DataReaderImpl) Offset() int {
	return r.offset
}

func (r *DataReaderImpl) ReadU8() uint8 {
	val := r.data[r.offset]
	r.offset++
	return val
}

func (r *DataReaderImpl) ReadU16() uint16 {
	val := binary.BigEndian.Uint16(r.data[r.offset : r.offset+2])
	r.offset += 2
	return val
}

func (r *DataReaderImpl) ReadU32() uint32 {
	val := binary.BigEndian.Uint32(r.data[r.offset : r.offset+4])
	r.offset += 4
	return val
}

func (r *DataReaderImpl) ReadString(doubleByte bool) string {
	var length = r.ReadU32()

	str := ""
	if doubleByte {

		for length > 0 {
			r.offset += 1
			str += string(r.data[r.offset])
			r.offset += 1
			length -= 2
		}
		str = str[:len(str)-1]

	} else {
		str = string(r.data[r.offset : r.offset+int(length-1)])
		r.offset += int(length)

	}
	return str

}

func (r *DataReaderImpl) ReadBytes(length int) []byte {
	bytes := r.data[r.offset : r.offset+length]
	r.offset += length
	return bytes
}

func (r *DataReaderImpl) Skip(length int) {
	r.offset += length
}
