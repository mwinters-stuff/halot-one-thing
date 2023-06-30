package main

import (
	"os"
	"time"

	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/mwinters-stuff/halo-one-thing/common"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
	"golang.org/x/net/websocket"
)

func passToToken(password string) string {
	key, _ := hex.DecodeString("6138356539643638")

	block, err := des.NewCipher(key)
	if err != nil {
		common.Logger.Error().Msgf("Error creating DES cipher: %e", err)
		return ""
	}

	ecb := ecb.NewECBEncrypter(block)
	paddedPassword := padPassword(password)

	encrypted := make([]byte, len(paddedPassword))
	ecb.CryptBlocks(encrypted, []byte(paddedPassword))

	finalEncrypted := base64.StdEncoding.EncodeToString(encrypted)
	return finalEncrypted
}

func padPassword(password string) string {
	padLen := des.BlockSize - len(password)%des.BlockSize
	padding := make([]byte, padLen)
	for i := range padding {
		padding[i] = byte(padLen)
	}
	return password + string(padding)
}

func sendData(conn *websocket.Conn, data []byte) {
	common.Logger.Info().Msgf("Sending %s\n", data)
	n, err := conn.Write(data)
	if err != nil {
		common.Logger.Fatal().Err(err)
		os.Exit(1)
	}
	common.Logger.Info().Msgf("wrote %d\n", n)
}

func getVersion(ws *websocket.Conn, startTimerChan chan bool) {
	data, err := json.Marshal(jsontypes.MessageCommand{Cmd: "GET_VERSION"})
	if err != nil {
		common.Logger.Fatal().Err(err)
		os.Exit(1)
	}
	sendData(ws, data)

	var msg = make([]byte, 100)
	var n int
	if n, err = ws.Read(msg); err != nil {
		common.Logger.Fatal().Err(err)
		os.Exit(1)
	}

	var versionMessage receive.Version
	if versionMessage, err = receive.UnmarshalVersion(msg[:n]); err != nil {
		common.Logger.Fatal().Err(err)
		os.Exit(1)

	}

	common.Logger.Info().Msgf("Received Version: %s.\n", versionMessage.Version)

	if versionMessage.Version != "1" {
		common.Logger.Fatal().Msgf("can not use version %s", versionMessage.Version)
		os.Exit(1)
	}
	go func() {
		startTimerChan <- true
	}()
}

func getStatus(conn *websocket.Conn, token string) {
	msg := send.TokenMessage{
		Token:          token,
		MessageCommand: jsontypes.MessageCommand{Cmd: "GET_PRINT_STATUS"},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		common.Logger.Fatal().Err(err)
		os.Exit(1)
	}
	sendData(conn, data)
}

func readLoop(conn *websocket.Conn, target chan []byte) {
	for {
		var n int
		var msg = make([]byte, 512)
		var err error

		if n, err = conn.Read(msg); err != nil {
			common.Logger.Fatal().Err(err)
			os.Exit(1)
		}
		common.Logger.Info().Msgf("Received: %s.\n", msg[:n])
		target <- msg[:n]
	}
}

func main() {
	origin := "http://halot-one/"
	url := "ws://localhost:18188/&password=groot"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		common.Logger.Fatal().Err(err)
		os.Exit(1)
	}
	token := passToToken("groot")

	readSocket := make(chan []byte)
	connectChannel := make(chan bool)
	startTimerChannel := make(chan bool)

	go func() {
		connectChannel <- true
	}()

	statusTimer := time.NewTicker(10 * time.Second)
	defer statusTimer.Stop()
	var runStatusTimer bool = false

	for {
		select {
		case <-statusTimer.C:
			if runStatusTimer {
				getStatus(ws, token)
			}
		case data := <-readSocket:
			readData(data)
		case <-connectChannel:
			common.Logger.Info().Msg("Connected Get Version")
			getVersion(ws, startTimerChannel)
		case runStatusTimer = <-startTimerChannel:
			go readLoop(ws, readSocket)
			common.Logger.Info().Msg("Start Read Loop & Timer")
		}

	}
}

func readData(data []byte) {
	incomingMsg, err := jsontypes.UnmarshalMessageCommand(data)
	if err != nil {
		common.Logger.Fatal().Err(err)
		os.Exit(1)
	}

	switch incomingMsg.Cmd {
	case "GET_PRINT_STATUS":
		{
			printStatus, err := receive.UnmarshalPrinterStatus(data)
			if err != nil {
				common.Logger.Fatal().Err(err)
				os.Exit(1)
			}
			common.Logger.Info().Msgf("Status %s\n", printStatus.PrintStatus)
		}
	case "START_FILE":
		{
			startFile, err := receive.UnmarshalStartFile(data)
			if err != nil {
				common.Logger.Fatal().Err(err)
				os.Exit(1)
			}
			common.Logger.Info().Msgf("StartFile %s\n", startFile.Filename)
		}
	}

}
