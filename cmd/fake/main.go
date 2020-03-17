package main

import (
	"github.com/gin-gonic/gin"
	server "github.com/yimkh/ws/server"
)

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		server.WsHandler(c.Writer, c.Request)
	})

	r.Run(":8001")
}
