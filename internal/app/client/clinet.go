package client

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/dink10/poker/internal/pkg/client"
	"github.com/dink10/poker/internal/pkg/config"
	"github.com/dink10/poker/internal/pkg/logger"
)

// Run runs application.
func Run() error {
	logrus.Info("Client start")
	defer logrus.Info("Client end")

	var cfg Config
	err := config.LoadConfig(&cfg)
	if err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	err = logger.Init(&cfg.Logger)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	c := client.NewClient(&cfg.Client)
	res, err := c.DoRequest(cfg.Client.HttpAddress, http.MethodGet, client.Options{})
	if err != nil {
		return err
	}

	logrus.Infof("Response from server: %s", string(res))

	return nil
}
