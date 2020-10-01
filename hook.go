package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os/exec"
)

// Grafana web hook struct
type hookAlert struct {
	EvalMatches []struct {
		Metric string      `json:"metric"`
		Tags   interface{} `json:"tags"`
		Value  float64     `json:"value"`
	} `json:"evalMatches"`
	ImageUrl string `json:"imageUrl"`
	Message  string `json:"message"`
	RuleId   int    `json:"ruleId"`
	RuleName string `json:"ruleName"`
	RuleUrl  string `json:"ruleUrl"`
	State    string `json:"state"`
	Title    string `json:"title"`
}

var webhookLockChan = make(chan bool, 1)

// basicAuth
func basicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		Config.Webhook.Username: Config.Webhook.Password,
	})
}

// grafana webhook handler
func handlerWebhook(c *gin.Context) {
	select {
	case webhookLockChan <- true:
		defer func() {
			<-webhookLockChan
		}()
		break
	default:
		logrus.WithFields(logrus.Fields{"controller": "webhook"}).Warnln("busy")
		return
	}

	// parse json
	data := new(hookAlert)
	err := c.BindJSON(data)
	if err != nil {
		logrus.Warnln(err)
		return
	}

	// only except exec
	if data.State != "ok" {
		cmd := exec.Command(Config.Shell.Cmd.Name, Config.Shell.Cmd.Args...)
		if err = cmd.Run(); err != nil {
			logrus.Warnln(err)
			return
		}
		c.JSON(200, false)
		return
	}

	c.JSON(200, false)
}
