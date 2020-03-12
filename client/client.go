package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	types "github.com/yimkh/ws/types"
)

//SendReq is to send request
func SendReq(backendServer *httptest.Server) {
	req, _ := http.NewRequest("GET", backendServer.URL, nil)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")

	// can use client to do request to backend
	c := backendServer.Client()

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 101 {
		log.Fatalf("status = %v; want 101", res.Status)
	}
	if types.UpgradeType(res.Header) != "websocket" {
		log.Fatalf("not websocket upgrade; got %#v", res.Header)
	}
	rwc, ok := res.Body.(io.ReadWriteCloser)
	log.Println("frontproxy got ws ReadWriteCloser rwc")
	if !ok {
		log.Fatalf("response body is of type %T; does not implement ReadWriteCloser", res.Body)
	}
	defer rwc.Close()

	io.WriteString(rwc, "Hello\n")
	bs := bufio.NewScanner(rwc)
	if !bs.Scan() {
		log.Fatalf("Scan: %v", bs.Err())
	}
	got := bs.Text()
	log.Println("frontproxy read from response and got:", got)
	want := `backend got "Hello"`
	if got != want {
		log.Println(fmt.Errorf("got %#q, want %#q", got, want))
	}
}
