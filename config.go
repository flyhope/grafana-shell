package main

import (
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type configXml struct {
	Webhook struct {
		Username string
		Password string
	}
	Shell struct {
		Cmd struct {
			Name string
			Args []string
		}
	}
}

var Config = new(configXml)

func loadConfig() {
	// load xml file
	configFile := fechCurrentPath() + "config.xml"
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.Panic(err)
	}

	// parse xml
	err = xml.Unmarshal(content, Config)
	if err != nil {
		logrus.Panicln(err)
	}
}
