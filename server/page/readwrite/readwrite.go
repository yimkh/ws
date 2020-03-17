package readwrite

import (
	"log"

	"github.com/gorilla/websocket"
)

func Read(c *websocket.Conn) {
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		break
	}
	log.Printf("recv: %s", message)
	log.Printf("recv type: %t", mt)
}

func Write(c *websocket.Conn, mt int, message []byte) {
	err = c.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)
		break
	}
}
