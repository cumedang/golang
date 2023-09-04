package p2p

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cumedang/go/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	//port :3000이 port :4000에서 온 request를 upgrad합니다
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandelERror(err)
	openProt := r.URL.Query().Get("openPort")
	result := strings.Split(r.RemoteAddr, ":")
	initPeer(conn, result[0], openProt)

}

func AddoPeer(address, port, openPort string) {
	//Porst :4000이 porst:3000과 연결하기를 원해요
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandelERror(err)
	initPeer(conn, address, port)
}
