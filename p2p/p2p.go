package p2p

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cumedang/go/utils"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrade.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrade.Upgrade(rw, r, nil)
	utils.HandelERror(err)
	for {
		_, p, err := conn.ReadMessage()
		utils.HandelERror(err)
		fmt.Printf("Just got:%s\n\n", p)
		time.Sleep(5 * time.Second)
		message := fmt.Sprintf("wa also think that: %s", p)
		utils.HandelERror(conn.WriteMessage(websocket.TextMessage, []byte(message)))
	}
}
