package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	probabtest "github.com/yimkh/ws/server/page/probabtest"
)

//Upgrader is websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//WssHandler is to do
func WssHandler(w http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		probabtest.Ptest(w)
		return
	}
}
