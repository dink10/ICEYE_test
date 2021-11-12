package main

import (
	"github.com/sirupsen/logrus"

	"github.com/dink10/poker/info"
	"github.com/dink10/poker/internal/app/poker"
)

func main() {
	logrus.Infof("Application version: %s", info.Version)

	if err := poker.Run(); err != nil {
		logrus.Fatal(err)
	}
}
