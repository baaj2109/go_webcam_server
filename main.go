package main

import (
	"fmt"
	"net"

	"github.com/baaj2109/webcam_server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// err := engine.Run(":8080")
	router.InitRouter(engine)
	listener, err := net.Listen("tcp", ":8080")
	if nil != err {
		fmt.Println(err)
		return
	}
	engine.RunListener(listener)

}
