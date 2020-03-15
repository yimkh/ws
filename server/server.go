package main

import (
	"github.com/gin-gonic/gin"
	model "github.com/yimkh/ws/server/model"
)

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		model.WsHandler(c.Writer, c.Request)
	})

	r.Run(":8001")
}
