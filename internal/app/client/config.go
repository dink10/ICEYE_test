package client

import (
	"github.com/dink10/poker/internal/pkg/client"
	"github.com/dink10/poker/internal/pkg/logger"
)

// Config is an application config.
type Config struct {
	Client client.Config
	Logger logger.Config
}
