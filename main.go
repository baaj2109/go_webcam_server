package main

import (
	"fmt"
	"net"

	"github.com/baaj2109/webcam_server/config"
	"github.com/baaj2109/webcam_server/global"
	"github.com/baaj2109/webcam_server/router"
	"github.com/gin-gonic/gin"
)

func main() {

	initSettings()

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

func initSettings() {
	if err := config.InitConfig(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
	// 	fmt.Printf("init logger failed, err:%v\n", err)
	// 	return
	// }
	if err := global.InitMySqlDb(config.Conf.MySqlConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer global.MySQLDb.Close()

	if err := global.InitRedisDb(config.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	defer global.RedisDb.Close()
}
