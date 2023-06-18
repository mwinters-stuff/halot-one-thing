package main

import (
	"fmt"
	"log"

	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/mwinters-stuff/halo-one-thing/jsontypes"
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

func getStatus(conn *websocket.Conn, token string) {
	msg := jsontypes.OutgoingMessage{
		Cmd:   "GET_PRINT_STATUS",
		Token: token,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	sendData(conn, data)
}

func startFile(conn *websocket.Conn, token string, filename string) {
	msg := jsontypes.StartFileMessage{
		Cmd:      "START_FILE",
		Token:    token,
		Filename: filename,
		Offset:   "0",
		Size:     "10000",
	}

	data, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	sendData(conn, data)
}

func main() {
	origin := "http://halot-one/"
	url := "ws://halot-one:18188/&password=groot"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	token := passToToken("groot")

	getStatus(ws, token)
	startFile(ws, token, "afile.txt")

	var msg = make([]byte, 512)

	for {
		var n int

		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])

		incomingMsg, err := jsontypes.UnmarshalIncomingMessage(msg[:n])
		if err != nil {
			log.Fatal(err)
		}
		if incomingMsg.Cmd == "GET_PRINT_STATUS" {
			printStatus, err := jsontypes.UnmarshalPrintStatusIncoming(msg[:n])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Status %s\n", printStatus.PrintStatus)
		}
		if incomingMsg.Cmd == "START_FILE" {
			startFile, err := jsontypes.UnmarshalStartFileIncoming(msg[:n])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("StartFile %s\n", startFile.Filename)
		}
	}
}
