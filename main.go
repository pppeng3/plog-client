package main

import (
	"plog_client/handler"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}

func main() {
	handler.SendLog("pp")
}
