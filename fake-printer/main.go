package main

import (
	"bytes"
	"compress/zlib"
	"embed"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/mwinters-stuff/halo-one-thing/common"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
)

//go:embed cr.html
var content embed.FS

type FakePrinter interface {
	webSocketServer(websocket *websocket.Conn)
	getVersion()
	getPrintStatus([]byte)
	sendData([]byte)
	printParametersSet([]byte)
	printPause([]byte)
	printStop([]byte)
	startPrint([]byte)
	startFile([]byte)
	sendStartData(offset int)
	readFileData([]byte)
	sendPrintStatus()
	Run()
}

type FakePrinterImpl struct {
	printerStatus    receive.PrinterStatus
	printParameters  send.PrintParametersSet
	startFileMessage send.StartFile
	startDataMessage receive.StartData
	ws               *websocket.Conn
	fileSize         int64
}

func (p *FakePrinterImpl) webSocketServer(ws *websocket.Conn) {
	p.ws = ws
	for {
		messageType, msg, err := p.ws.ReadMessage()
		if err != nil {
			common.Logger.Info().Err(err)
			return
		}

		if messageType == websocket.TextMessage {

			var messageCommand jsontypes.MessageCommand
			if messageCommand, err = jsontypes.UnmarshalMessageCommand(msg); err != nil {
				common.Logger.Fatal().Msg(err.Error())
			}

			common.Logger.Info().Msgf("received Command: %s\n%s", messageCommand.Cmd, msg)

			switch messageCommand.Cmd {
			case "GET_VERSION":
				p.getVersion()
			case "GET_PRINT_STATUS":
				p.getPrintStatus(msg)
			case "PRINT_PARA_SET":
				p.printParametersSet(msg)
			case "PRINT_PAUSE":
				p.printPause(msg)
			case "PRINT_STOP":
				p.printStop(msg)
			case "START_PRINT":
				p.startPrint(msg)
			case "START_FILE":
				p.startFile(msg)
			}
		}
		if messageType == websocket.BinaryMessage {
			common.Logger.Info().Msgf("Received binary message %d long", len(msg))
			p.readFileData(msg)
		}
	}
}

func (p *FakePrinterImpl) printPause(data []byte) {
	var tokenMessage send.TokenMessage
	var err error
	if tokenMessage, err = send.UnmarshalTokenMessage(data); err != nil {
		common.Logger.Fatal().Msg(err.Error())

	}
	if tokenMessage.Token == "" {
		common.Logger.Info().Msgf("token incorrect %s", tokenMessage.Token)
	}

	message := receive.Status{
		MessageCommand: jsontypes.MessageCommand{Cmd: "PRINT_PAUSE"}, Status: p.printerStatus.PrintStatus,
	}

	if p.printerStatus.PrintStatus == "PRINT_PROCESSING" {
		p.printerStatus.PrintStatus = "PRINT_STOP"
	}

	writeData, err := message.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	p.sendData(writeData)
	p.sendPrintStatus()
}

// printStop implements FakePrinter.
func (p *FakePrinterImpl) printStop(data []byte) {
	var tokenMessage send.TokenMessage
	var err error
	if tokenMessage, err = send.UnmarshalTokenMessage(data); err != nil {
		common.Logger.Fatal().Msg(err.Error())

	}

	if tokenMessage.Token == "" {
		common.Logger.Info().Msgf("token incorrect %s", tokenMessage.Token)
	}
	message := receive.Status{
		MessageCommand: jsontypes.MessageCommand{Cmd: "PRINT_STOP"}, Status: p.printerStatus.PrintStatus,
	}

	if p.printerStatus.PrintStatus == "PRINT_PROCESSING" || p.printerStatus.PrintStatus == "PRINT_STOP" {
		p.printerStatus.PrintStatus = "PRINT_END"
	}

	writeData, err := message.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())

	}
	p.sendData(writeData)
	p.sendPrintStatus()
}

func (p *FakePrinterImpl) startFile(data []byte) {

	var err error
	if p.startFileMessage, err = send.UnmarshalStartFile(data); err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}

	var f *os.File
	if f, err = os.Create(p.startFileMessage.Filename); err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	f.Close()

	sendStartFile := receive.StartFile{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "START_FILE",
		},
		Compress: true,
		Filename: p.startFileMessage.Filename,
		Key:      p.startFileMessage.Key,
		Offset:   "0",
		Size:     p.startFileMessage.Size,
	}

	sendData, err := sendStartFile.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	p.sendData(sendData)

	i, _ := strconv.Atoi(p.startFileMessage.Size)
	p.fileSize = int64(i)

	// prepare start data message, but dont send it.
	p.startDataMessage = receive.StartData{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "START_DATA",
		},
		Errorcode: int64(0),
		Key:       p.startFileMessage.Key,
		Received:  "0",
		Size:      p.startFileMessage.Size,
	}
}

func (p *FakePrinterImpl) sendStartData(offset int) {
	p.startDataMessage.Received = fmt.Sprintf("%d", offset)

	sendData, err := p.startDataMessage.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	// p.packetSize = 0
	p.sendData(sendData)
}

func (p *FakePrinterImpl) readFileData(data []byte) {
	packetSize := int(binary.BigEndian.Uint32(data[:4]))

	common.Logger.Debug().Msgf("Uncompressed Expected %d bytes, buffer is %d", packetSize, len(data[4:]))
	var bytesBuffer bytes.Buffer

	bytesBuffer.Write(data[4:])
	zr, err := zlib.NewReader(&bytesBuffer)
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}

	defer zr.Close()

	f, err := os.OpenFile(p.startFileMessage.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	if fileInfo.Size() < p.fileSize {

		var nn int64
		if nn, err = io.Copy(f, zr); err != nil {
			common.Logger.Fatal().Msg(err.Error())
		}

		fileInfo, err = f.Stat()
		if err != nil {
			common.Logger.Fatal().Msg(err.Error())
		}
		common.Logger.Info().Msgf("Wrote %d to %s, now %d remaining %d", nn, p.startFileMessage.Filename, fileInfo.Size(), p.fileSize-fileInfo.Size())
	}
	if fileInfo.Size() < p.fileSize {
		p.sendStartData(int(fileInfo.Size()))
	} else {
		cd := receive.CheckData{
			MessageCommand: jsontypes.MessageCommand{
				Cmd: "CHECK_DATA",
			},
			Checkstate: 1,
			Key:        p.startFileMessage.Key,
		}
		sendData, err := cd.Marshal()
		if err != nil {
			common.Logger.Fatal().Msg(err.Error())
		}
		p.sendData(sendData)
	}

}

func (p *FakePrinterImpl) startPrint(data []byte) {
	var startMessage send.StartPrint
	var err error
	if startMessage, err = send.UnmarshalStartPrint(data); err != nil {
		common.Logger.Fatal().Msg(err.Error())

	}
	if startMessage.Token == "" {
		common.Logger.Info().Msgf("token incorrect %s", startMessage.Token)
	}

	printStartStatus := receive.StartPrintStatus{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "START_PRINT",
		},
		Filename: startMessage.Filename,
		Status:   "CHECKING",
	}

	sendData, err := printStartStatus.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	p.sendData(sendData)

	timer := time.NewTimer(time.Second)
	go func() {
		<-timer.C
		{
			printStartStatus.Status = "STARTED"
			sendData, err := printStartStatus.Marshal()
			if err != nil {
				common.Logger.Fatal().Msg(err.Error())
			}
			p.sendData(sendData)

			p.printerStatus.PrintStatus = "PRINT_PROCESSING"

			p.printerStatus.BottomExposureNum = "2"
			p.printerStatus.CurSliceLayer = "0"
			p.printerStatus.DelayLight = "4"
			p.printerStatus.EleSpeed = "1"
			p.printerStatus.Filename = startMessage.Filename
			p.printerStatus.InitExposure = "40"
			p.printerStatus.LayerThickness = "0.050000"
			p.printerStatus.PrintExposure = "3"
			p.printerStatus.PrintHeight = "6"
			p.printerStatus.PrintRemainTime = "1304"
			p.printerStatus.Resin = ""
			p.printerStatus.SliceLayerCount = "67"
		}
	}()

}

func (p *FakePrinterImpl) getPrintStatus(data []byte) {
	var tokenMessage send.TokenMessage
	var err error
	if tokenMessage, err = send.UnmarshalTokenMessage(data); err != nil {
		common.Logger.Fatal().Msg(err.Error())

	}
	if tokenMessage.Token == "" {
		common.Logger.Info().Msgf("token incorrect %s", tokenMessage.Token)
	}
	p.sendPrintStatus()
}

func (p *FakePrinterImpl) sendPrintStatus() {
	writeData, err := p.printerStatus.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	p.sendData(writeData)

}

func (p *FakePrinterImpl) printParametersSet(data []byte) {
	var pps send.PrintParametersSet
	var err error
	if pps, err = send.UnmarshalPrintParametersSet(data); err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	if pps.Token == "" {
		common.Logger.Info().Msgf("token incorrect %s", pps.Token)
	}
	p.printParameters = pps

	message := receive.Status{MessageCommand: jsontypes.MessageCommand{Cmd: "PRINT_PARA_SET"}, Status: ""}
	writeData, err := message.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())

	}
	p.sendData(writeData)
}

func (p *FakePrinterImpl) sendData(data []byte) {
	common.Logger.Info().Msgf("Sending %s", data)
	err := p.ws.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
}

func (p *FakePrinterImpl) getVersion() {
	msg := receive.Version{Version: "1"}
	data, err := msg.Marshal()
	if err != nil {
		common.Logger.Fatal().Msg(err.Error())
	}
	p.sendData(data)

}

func (p *FakePrinterImpl) Run() {
	p.printerStatus = receive.PrinterStatus{
		MessageCommand: jsontypes.MessageCommand{
			Cmd: "GET_PRINT_STATUS"},
		BottomExposureNum: "",
		CurSliceLayer:     "",
		DelayLight:        "",
		EleSpeed:          "",
		Filename:          "",
		InitExposure:      "",
		LayerThickness:    "",
		PrintExposure:     "",
		PrintHeight:       "",
		PrintRemainTime:   "",
		PrintStatus:       "PRINT_GENERAL",
		Resin:             "",
		SliceLayerCount:   "",
	}

	p.printParameters = send.PrintParametersSet{
		TokenMessage: send.TokenMessage{
			MessageCommand: jsontypes.MessageCommand{
				Cmd: "PRINT_PARA_SET"},
			Token: ""},
		BottomExposureNum: "2",
		DelayLight:        "4",
		EleSpeed:          "2",
		InitExposure:      "40",
		PrintExposure:     "3",
		PrintHeight:       "6",
	}

	common.Logger.Info().Msgf("Waiting for data.")

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
	}

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		common.Logger.Info().Msgf("Request %s", request.URL)
		if strings.HasPrefix(request.URL.Path, "/cr.html") {
			http.FileServer(http.FS(content)).ServeHTTP(response, request)
		} else {
			conn, err := upgrader.Upgrade(response, request, nil)
			if err != nil {
				common.Logger.Fatal().Msg(err.Error())

			}
			p.webSocketServer(conn)
			//websocket.Handler.ServeHTTP(p.webSocketServer, response, request)
		}

	})

	//http.Handle("/", websocket.Handler(p.webSocketServer))
	for {
		http.ListenAndServe(":18188", nil)
	}

}

// This example demonstrates a trivial echo server.
func main() {
	var fp FakePrinter = &FakePrinterImpl{}
	fp.Run()

}
