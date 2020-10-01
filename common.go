package main

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func fechCurrentPath() string {
	// fetch workdir from env
	workdir := os.Getenv("WORKDIR")
	if workdir != "" {
		return workdir + "/"
	}

	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		logrus.Panicln(err)
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		logrus.Panicln(err)
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		logrus.Panicln(errors.New(`Can't find "/" or "\"`))
		return ""
	}
	return path[0 : i+1]
}
