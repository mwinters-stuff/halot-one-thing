package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var ws *WebSocket
var stat Status
var statusTimer *time.Timer
var file *os.File
var token string

type Status struct {
	PrintStatus       string  `json:"printStatus"`
	Filename          string  `json:"filename"`
	PrintRemainTime   int     `json:"printRemainTime"`
	CurSliceLayer     int     `json:"curSliceLayer"`
	SliceLayerCount   int     `json:"sliceLayerCount"`
	PrintExposure     float64 `json:"printExposure"`
	LayerThickness    float64 `json:"layerThickness"`
	PrintHeight       float64 `json:"printHeight"`
	BottomExposureNum int     `json:"bottomExposureNum"`
	InitExposure      float64 `json:"initExposure"`
	DelayLight        float64 `json:"delayLight"`
	EleSpeed          float64 `json:"eleSpeed"`
	Resin             string  `json:"resin"`
}

func log(text string) {
	line := time.Now().Format("15:04:05") + " " + text + "<br>\n"
	fmt.Printf("%s\n", line)
}

func passToToken(password string) string {
	// from: https://github.com/tarequeh/DES/blob/master/run_des.c#L136

	key := "6138356539643638"

	encrypted := encryptDES(password, key)

	finalEncrypted := base64Encode(encrypted)

	return finalEncrypted
}

func printStatus(stat Status) {
	code := ""
	code += "Status: " + stat.PrintStatus + "<br>\n"
	code += "Filename: " + stat.Filename + "<br>\n"
	code += "Time left: " + time.Unix(int64(stat.PrintRemainTime), 0).UTC().Format("15:04:05") + "<br>\n"
	code += "Progress: " + fmt.Sprintf("%.1f", (float64(stat.CurSliceLayer)/float64(stat.SliceLayerCount))*100) + "%<br>\n"
	code += "Current layer: " + fmt.Sprintf("%d", stat.CurSliceLayer) + "<br>\n"
	code += "Total layers: " + fmt.Sprintf("%d", stat.SliceLayerCount) + "<br>\n"
	code += "Print exposure: " + fmt.Sprintf("%.1f", stat.PrintExposure) + "s<br>\n"
	code += "Layer thickness: " + fmt.Sprintf("%.1f", stat.LayerThickness) + "mm<br>\n"
	code += "Rising height: " + fmt.Sprintf("%.1f", stat.PrintHeight) + "mm<br>\n"
	code += "Bottom layers: " + fmt.Sprintf("%d", stat.BottomExposureNum) + "<br>\n"
	code += "Initial exposure: " + fmt.Sprintf("%.1f", stat.InitExposure) + "s<br>\n"
	code += "Turn off delay: " + fmt.Sprintf("%.1f", stat.DelayLight) + "s<br>\n"
	code += "Motor speed: " + fmt.Sprintf("%.1f", stat.EleSpeed) + "mm/s<br>\n"
	code += "Exposure time: " + fmt.Sprintf("%.1f", stat.DelayLight) + "s<br>\n"
	code += "Resin: " + stat.Resin + "<br>\n"

	fmt.Println(code)
}

func printFileStatus(stat Status) {
	code := ""
	code += "Sent: " + fmt.Sprintf("%.2f", float64(stat.Received)/(1024*1024)) + " MB<br>\n"
	code += "Size: " + fmt.Sprintf("%.2f", float64(stat.Size)/(1024*1024)) + " MB<br>\n"
	code += "Progress: " + fmt.Sprintf("%.0f", (float64(stat.Received)/float64(stat.Size))*100) + "%<br>\n"
	code += "Checked: " + func() string {
		if stat.Checkstate {
			return "OK"
		}
		return "ERROR"
	}() + "<br>\n"

	fmt.Println(code)
}

func sendCmd(cmd string, extras map[string]interface{}) {
	if ws == nil || ws.ReadyState() != WebSocketStateOpen {
		log("Not connected, please connect first.")
		return
	}

	msg := map[string]interface{}{
		"cmd":   cmd,
		"token": token,
	}

	for k, v := range extras {
		msg[k] = v
	}

	jmsg, err := json.Marshal(msg)
	if err != nil {
		log("Failed to encode JSON message: " + err.Error())
		return
	}

	log("> " + string(jmsg))
	ws.Send(jmsg)
}

func getStatus() {
	sendCmd("GET_PRINT_STATUS", nil)
}

func uploadFile() {
	if file == nil {
		log("Please choose a file first.")
		return
	}

	extras := map[string]interface{}{
		"filename": file.Name(),
		"offset":   "0",                            // why string?
		"size":     fmt.Sprintf("%d", file.Size()), // why?
	}

	sendCmd("START_FILE", extras)
}

func startPrint() {
	if confirm("Are you sure you want to start new print?") {
		sendCmd("START_PRINT", map[string]interface{}{"filename": file.Name()})
	}
}

func stopPrint() {
	if confirm("Are you sure you want to stop the print?") {
		sendCmd("PRINT_STOP", nil)
	}
}

func disconnect() {
	log("Disconnecting...")
	ws.Close()
}

func readEventHandler(evt FileReaderLoadEvent) {
	if evt.Err == nil {
		data := evt.Target.Result

		fmt.Println(data)

		zdata := deflate(data)

		len := make([]byte, 4)
		len[0] = byte((len(data) >> 24) & 0xff)
		len[1] = byte((len(data) >> 16) & 0xff)
		len[2] = byte((len(data) >> 8) & 0xff)
		len[3] = byte((len(data) >> 0) & 0xff)

		blob := append(len, zdata...)

		fmt.Println(blob)
		ws.Send(blob)
	} else {
		log("Read error: " + evt.Err.Error())
		return
	}
}

func sendChunk(offset int) {
	dataSlice := make([]byte, 0x10000)
	_, err := file.ReadAt(dataSlice, int64(offset))
	if err != nil {
		log("Failed to read data chunk: " + err.Error())
		return
	}

	reader := &FileReader{}

	reader.OnLoadEnd(func(evt FileReaderLoadEvent) {
		readEventHandler(evt)
	})

	reader.ReadAsArrayBuffer(dataSlice)
}

func connect() {
	log("Connecting...")
	url := document.GetElementById("url").Value
	password := document.GetElementById("password").Value
	token = passToToken(password)

	window.Location().Hash = "url=" + url + "&password=" + password

	// Let us open a web socket
	ws, err := NewWebSocket(url)
	if err != nil {
		log("Failed to create WebSocket: " + err.Error())
		return
	}

	ws.OnOpen(func(evt Event) {
		log("Connected.")

		getStatus()

		statusTimer = time.NewTimer(5 * time.Second)
		go func() {
			for range statusTimer.C {
				getStatus()
			}
		}()
	})

	ws.OnMessage(func(evt MessageEvent) {
		receivedMsg := evt.Data
		msg := make(map[string]interface{})
		err := json.Unmarshal([]byte(receivedMsg), &msg)
		if err != nil {
			log("Failed to decode JSON message: " + err.Error())
			return
		}

		cmd := msg["cmd"].(string)

		if cmd == "GET_PRINT_STATUS" {
			printStatus(msg)
		} else if cmd == "START_FILE" {
			offset := int(msg["offset"].(float64))
			sendChunk(offset)
		} else if cmd == "START_DATA" {
			offset := int(msg["received"].(float64))
			printFileStatus(msg)
			sendChunk(offset)
		} else if cmd == "CHECK_DATA" {
			printFileStatus(msg)
		}

		log("< " + receivedMsg)
	})

	ws.OnClose(func(evt CloseEvent) {
		statusTimer.Stop()
		log("Connection is closed...")
	})

}

func updateParamsFromHash() {
	params := strings.SplitN(window.Location().Hash[1:], "&", -1)

	var url, password string

	for _, param := range params {
		parts := strings.SplitN(param, "=", 2)
		if len(parts) == 2 {
			switch parts[0] {
			case "url":
				url = parts[1]
			case "password":
				password = parts[1]
			}
		}
	}

	if url != "" {
		document.GetElementById("url").Value = url
	}

	if password != "" {
		document.GetElementById("password").Value = password
	}
}

func onPageLoad() {
	fmt.Println(window.Location().Hash)

	if _, ok := window.(interface {
		OnHashChange(func()) func()
	}); ok {
		window.OnHashChange(func() {
			updateParamsFromHash()
		})
	}

	updateParamsFromHash()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(data))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
