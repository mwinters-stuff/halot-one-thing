package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const dataTerm uint16 = 0x0d0a

type CXDLP struct {
	Magic1        string
	Magic2        string
	Model         string
	Version       uint16
	LayerCount    uint16
	ResX          uint16
	ResY          uint16
	Height        uint32
	Thumb         []byte
	Preview1      []byte
	Preview2      []byte
	DimX          string
	DimY          string
	Layer         string
	LightOn       uint16
	LightOff      uint16
	LightPWM      uint16
	LiftDist      uint16
	LiftSpeed     uint16
	DownSpeed     uint16
	BaseLayers    uint16
	BaseLightOn   uint16
	BaseLightPWM  uint16
	BaseLiftDist  uint16
	BaseLiftSpeed uint16
	Layers        []Layer
}

type Layer struct {
	Length    uint32
	LineCount uint32
	Pos       int
	Lines     []Line
}

type Line struct {
	YStart uint16
	YEnd   uint16
	XEnd   uint16
	Color  uint8
}

func NewCXDLP() *CXDLP {
	return &CXDLP{
		Magic1:        "CXSW3DV2",
		Magic2:        "CXSW3DV2",
		Model:         "CL-89",
		Version:       1,
		LayerCount:    0,
		ResX:          3840,
		ResY:          2400,
		Height:        0,
		DimX:          "192.000000",
		DimY:          "120.000000",
		Layer:         "0.050000",
		LightOn:       6,
		LightOff:      2,
		LightPWM:      255,
		LiftDist:      6,
		LiftSpeed:     60,
		DownSpeed:     150,
		BaseLayers:    8,
		BaseLightOn:   60,
		BaseLightPWM:  255,
		BaseLiftDist:  5,
		BaseLiftSpeed: 60,
	}
}

func (c *CXDLP) Read(data []byte) *CXDLP {
	reader := NewDataReader(data, 0)

	c.Magic1 = reader.ReadString(false)
	c.Version = reader.ReadU16()
	c.Model = reader.ReadString(false)
	c.LayerCount = reader.ReadU16()
	c.ResX = reader.ReadU16()
	c.ResY = reader.ReadU16()
	c.Height = reader.ReadU32()
	reader.Skip(60)
	c.Thumb = reader.ReadBytes(26912)
	if reader.ReadU16() != dataTerm {
		panic("invalid data term")
	}
	c.Preview1 = reader.ReadBytes(168200)
	if reader.ReadU16() != dataTerm {
		panic("invalid data term")
	}
	c.Preview2 = reader.ReadBytes(168200)
	if reader.ReadU16() != dataTerm {
		panic("invalid data term")
	}
	c.DimX = reader.ReadString(true)
	c.DimY = reader.ReadString(true)
	c.Layer = reader.ReadString(true)
	c.LightOn = reader.ReadU16()
	c.LightOff = reader.ReadU16()
	c.BaseLightOn = reader.ReadU16()
	c.BaseLayers = reader.ReadU16()
	c.BaseLiftDist = reader.ReadU16()
	c.BaseLiftSpeed = reader.ReadU16()
	c.LiftDist = reader.ReadU16()
	c.LiftSpeed = reader.ReadU16()
	c.DownSpeed = reader.ReadU16()
	c.BaseLightPWM = reader.ReadU16()
	c.LightPWM = reader.ReadU16()

	// Read layer record lengths
	c.Layers = make([]Layer, c.LayerCount)
	for i := 0; i < int(c.LayerCount); i++ {
		layer := Layer{}
		layer.Length = reader.ReadU32()
		fmt.Printf("Layer %d Length %d\n", i, layer.Length)
		c.Layers[i] = layer
	}
	if reader.ReadU16() != dataTerm {
		panic("invalid data term")
	}

	// read layer meta data, not line data
	for i := 0; i < int(c.LayerCount); i++ {
		layer := c.Layers[i]
		size := reader.ReadU32()
		if size != layer.Length {
			panic(fmt.Sprintf("layer length mismatch: %d != %d @ i=%d", size, layer.Length, i))
		}
		lineCount := reader.ReadU32()
		layer.LineCount = lineCount
		layer.Pos = reader.Offset

		layer.Lines = make([]Line, lineCount)

		for j := 0; j < int(lineCount); j++ {
			line := Line{}
			d1 := reader.ReadU16()
			d2 := reader.ReadU16()
			d3 := reader.ReadU16()
			line.YStart = d1 >> 3
			line.YEnd = ((d1 & 0b111) << 10) | (d2 >> 6)
			line.XEnd = ((d2 & 0b111111) << 8) | (d3 >> 8)
			line.Color = uint8(d3 & 0xff)
			layer.Lines[j] = line
		}

	}

	return c
}

// DataReader represents a helper struct for reading binary data
type DataReader struct {
	data   []byte
	Offset int
}

// NewDataReader creates a new instance of DataReader
func NewDataReader(data []byte, offset int) *DataReader {
	return &DataReader{
		data:   data,
		Offset: offset,
	}
}

func (r *DataReader) ReadU8() uint8 {
	val := r.data[r.Offset]
	r.Offset++
	return val
}

func (r *DataReader) ReadU16() uint16 {
	val := binary.BigEndian.Uint16(r.data[r.Offset : r.Offset+2])
	r.Offset += 2
	return val
}

func (r *DataReader) ReadU32() uint32 {
	val := binary.BigEndian.Uint32(r.data[r.Offset : r.Offset+4])
	r.Offset += 4
	return val
}

func (r *DataReader) ReadString(doubleByte bool) string {
	var length = r.ReadU32()

	str := ""
	if doubleByte {

		for length > 0 {
			r.Offset += 1
			str += string(r.data[r.Offset])
			r.Offset += 1
			length -= 2
		}

	} else {
		str = string(r.data[r.Offset : r.Offset+int(length-1)])
		r.Offset += int(length)

	}
	return str

}

func (r *DataReader) ReadBytes(length int) []byte {
	bytes := r.data[r.Offset : r.Offset+length]
	r.Offset += length
	return bytes
}

func (r *DataReader) Skip(length int) {
	r.Offset += length
}

// Write saves the CXDLP data to a binary file
// func (c *CXDLP) Write(filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := NewDataWriter()

// 	// Write header
// 	writer.WriteU16(c.FileType)
// 	writer.WriteString(c.FileVersion, true)
// 	writer.WriteU16(c.Width)
// 	writer.WriteU16(c.Height)
// 	writer.WriteString(c.ProjectName, true)
// 	writer.WriteString(c.PrintTime, true)
// 	writer.WriteString(c.PrintDate, true)
// 	writer.WriteU16(c.ExposureTime)
// 	writer.WriteU16(c.LiftingDistance)
// 	writer.WriteU16(c.BottomLiftDistance)
// 	writer.WriteU16(c.LiftingSpeed)
// 	writer.WriteU16(c.BottomLiftSpeed)
// 	writer.WriteU8(c.LayerCount)
// 	writer.WriteU16(c.DownSpeed)
// 	writer.WriteU16(c.BaseLightPWM)

// 	// Write layers
// 	for _, layer := range c.Layers {
// 		writer.WriteU32(layer.Length)
// 		for _, line := range layer.Lines {
// 			writer.WriteU16(line.YStart)
// 			writer.WriteU16(line.YEnd)
// 			writer.WriteU16(line.XEnd)
// 			writer.WriteU8(line.Color)
// 		}
// 	}

// 	_, err = file.Write(writer.Bytes())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// DataWriter represents a helper struct for writing binary data
type DataWriter struct {
	buffer bytes.Buffer
}

// NewDataWriter creates a new instance of DataWriter
func NewDataWriter() *DataWriter {
	return &DataWriter{}
}

func (w *DataWriter) WriteU8(val uint8) {
	w.buffer.WriteByte(val)
}

func (w *DataWriter) WriteU16(val uint16) {
	binary.Write(&w.buffer, binary.LittleEndian, val)
}

func (w *DataWriter) WriteU32(val uint32) {
	binary.Write(&w.buffer, binary.LittleEndian, val)
}

func (w *DataWriter) WriteString(str string, isFixedLength bool) {
	if isFixedLength {
		w.WriteU16(uint16(len(str)))
	}
	w.buffer.WriteString(str)
}

func (w *DataWriter) Bytes() []byte {
	return w.buffer.Bytes()
}
