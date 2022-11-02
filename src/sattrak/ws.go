package main

import (
	"github.com/go-yaml/yaml"
	"github.com/gorilla/websocket"

	"encoding/json"
	"net/http"
	"sync"
	"time"
)

var (
	upgrader       = websocket.Upgrader{}
	wsConns        map[int]*wsConn
	newWSConnMutex sync.RWMutex
	wsMutex        sync.RWMutex
	wsConnNum      = 0
)

type wsConn struct {
	conn          *websocket.Conn
	mutex         sync.RWMutex
	authenticated bool
	display       bool
	connNum       int
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
		return
	}
	wc := new(wsConn)
	wc.conn = conn
	wc.authenticated = true
	newWSConnMutex.Lock()
	wc.connNum = wsConnNum
	wsConns[wc.connNum] = wc
	wsConnNum++
	packet := WSPacket{"info", infoData}
	wsContent, _ := json.Marshal(packet)
	sendWS(wc, wsContent)
	wsContent = getModifyDataValues()
	sendWS(wc, wsContent)
	newWSConnMutex.Unlock()
	go connRead(wc)
}

func connRead(wc *wsConn) {
	defer wc.conn.Close()
	defer delete(wsConns, wc.connNum)
	//wc.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	for {
		if wc.authenticated {
			wc.conn.SetReadDeadline(time.Time{})
		}
		//_, _, err := wc.conn.ReadMessage()
		_, content, err := wc.conn.ReadMessage()
		if err != nil {
			break
		}
		go func(content []byte) {
			contentVals := new(WSPacket)
			json.Unmarshal(content, &contentVals)
			if contentVals.PacketType == "devices" {
				d := new(DevicesPacket)
				json.Unmarshal(content, &d)
				if len(d.Packet.Devices) > 0 {
					modifyDevicesStatus(d.Packet.Devices[0].Index, d.Packet.Devices[0].Status)
					packet := WSPacket{"devices", deviceData}
					wsContent, _ := json.Marshal(packet)
					sendAllWS(wsContent)
				}
			} else if contentVals.PacketType == "config" {
				c := new(ConfigPacket)
				yaml.Unmarshal(content, &c)
				editConfig(c.Packet)
				sendModifyDataValues()
			}
		}(content)
	}
}

func sendAllWS(content []byte) {
	wsMutex.Lock()
	for _, wc := range wsConns {
		sendWS(wc, content)
	}
	wsMutex.Unlock()
}

func sendWS(wc *wsConn, content []byte) {
	if !wc.authenticated {
		return
	}
	wc.mutex.Lock()
	err := wc.conn.WriteMessage(websocket.TextMessage, content)
	if err != nil {
		wc.conn.Close()
		delete(wsConns, wc.connNum)
	}
	wc.mutex.Unlock()
}
