package main

import (
	"github.com/gin-gonic/gin"
)

// home index handler
func indexHandler(c *gin.Context) {
	response := map[string]interface{}{
		"useage": "go to grafana add /webhook/ for webhook",
	}
	c.JSON(200, response)
}
