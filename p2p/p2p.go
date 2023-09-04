package p2p

import (
	"fmt"
	"net/http"

	"github.com/cumedang/go/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	openProt := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openProt != "" && ip != ""
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandelERror(err)
	initPeer(conn, ip, openProt)

}

func AddoPeer(address, port, openPort string) {
	//Porst :4000이 porst:3000과 연결하기를 원해요
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandelERror(err)
	initPeer(conn, address, port)
}
