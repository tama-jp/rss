package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/frameworks/config"
	"github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/frameworks/logger"
	"github.com/tama-jp/rss/pkg/wire"
)

func main() {
	fmt.Println("*** 開始 ***")

	conf, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	dbConn, err := db.NewDB(conf)
	if err != nil {
		panic(err.Error())
	}

	logConn, err := logger.NewLogger(conf)
	if err != nil {
		panic(err.Error())
	}

	dbConn.Migration()
	dbConn.UserRolesSeqIDReset()

	r := gin.Default()

	appInstance := wire.InitializeApp(conf, dbConn, r, logConn)
	appInstance.Start()
}
