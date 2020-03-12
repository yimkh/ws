package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/yimkh/ws/client"
	types "github.com/yimkh/ws/types"
)

func main() {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if types.UpgradeType(r.Header) != "websocket" {
			log.Println("unexpected backend request")
			http.Error(w, "unexpected request", 400)
			return
		}
		log.Println("backend server get the websocket connection")
		c, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			log.Println(err)
			return
		}
		defer c.Close()
		log.Println("backend server upgrade http/1.1 101 to websocket")
		io.WriteString(c, "HTTP/1.1 101 Switching Protocols\r\nConnection: upgrade\r\nUpgrade: WebSocket\r\n\r\n")
		bs := bufio.NewScanner(c)
		if !bs.Scan() {
			log.Println(fmt.Errorf("backend failed to read line from client: %v", bs.Err()))
			return
		}
		fmt.Fprintf(c, "backend got %q\n", bs.Text())
	}))
	defer backendServer.Close()

	client.SendReq(backendServer)
}
