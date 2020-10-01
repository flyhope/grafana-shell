package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	loadConfig()

	router := gin.Default()
	router.NoRoute(indexHandler)

	webhook := router.Group("/webhook", basicAuth())
	webhook.POST("/", handlerWebhook)
	webhook.PUT("/", handlerWebhook)
	err := router.Run(":8850")
	if err != nil {
		logrus.Warnln(err)
	}
}
