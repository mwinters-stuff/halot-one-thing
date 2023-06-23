package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/receive"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes/send"
	"golang.org/x/net/websocket"
)

func passToToken(password string) string {
	key, _ := hex.DecodeString("6138356539643638")

	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating DES cipher:", err)
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

	n, err := conn.Write(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d\n", n)
}

func getVersion(ws *websocket.Conn) {
	data, err := json.Marshal(jsontypes.MessageCommand{Cmd: "GET_VERSION"})
	if err != nil {
		panic(err)
	}
	sendData(ws, data)

	var msg = make([]byte, 100)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}

	var versionMessage receive.Version
	if versionMessage, err = receive.UnmarshalVersion(msg[:n]); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received Version: %s.\n", versionMessage.Version)

	if versionMessage.Version != "1" {
		log.Fatal(errors.New(fmt.Sprintf("can no use version %s", versionMessage.Version)))
	}

}

func getStatus(conn *websocket.Conn, token string) {
	msg := send.TokenMessage{
		Token:          token,
		MessageCommand: jsontypes.MessageCommand{Cmd: "GET_PRINT_STATUS"},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	sendData(conn, data)
}

// func startFile(conn *websocket.Conn, token string, filename string) {
// 	msg := jsontypes.StartFileMessage{
// 		Cmd:      "START_FILE",
// 		Token:    token,
// 		Filename: filename,
// 		Offset:   "0",
// 		Size:     "10000",
// 	}

// 	data, err := json.Marshal(msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	sendData(conn, data)
// }

func main() {
	origin := "http://halot-one/"
	url := "ws://halot-one:18188/&password=groot"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	getVersion(ws)

	token := passToToken("groot")
	getStatus(ws, token)

	statusTimer := time.NewTimer(10 * time.Second)
	<-statusTimer.C
	getStatus(ws, token)

	var msg = make([]byte, 512)

	for {
		var n int

		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])

		incomingMsg, err := jsontypes.UnmarshalMessageCommand(msg[:n])
		if err != nil {
			log.Fatal(err)
		}

		if incomingMsg.Cmd == "GET_PRINT_STATUS" {
			printStatus, err := receive.UnmarshalPrinterStatus(msg[:n])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Status %s\n", printStatus.PrintStatus)
		}
		if incomingMsg.Cmd == "START_FILE" {
			startFile, err := receive.UnmarshalStartFile(msg[:n])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("StartFile %s\n", startFile.Filename)
		}
	}
}
