package api

import (
	"fmt"

	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
	logging "github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/logger"
	"github.com/gin-gonic/gin"
)

var (
	cfg    = config.GetConfig()
	logger = logging.NewLogger(cfg)
)

func InitApiServer() {
	serverPort := cfg.Server.Port
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	logger.Info(logging.General, logging.Api,
		fmt.Sprintf("Api Server Is Listening To Port:%s", serverPort), nil)
	err := r.Run(fmt.Sprintf(":%s", serverPort))
	if err != nil {
		logger.Fatal(logging.Internal, logging.Api,
			fmt.Sprintf("Error In Starting Api Server Error is :%s", err.Error()), nil)
	}
}
