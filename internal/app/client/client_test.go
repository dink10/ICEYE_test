package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dink10/poker/internal/pkg/client"
	"github.com/dink10/poker/internal/pkg/config"
)

var testSuite = []struct {
	url            string
	expectedResult string
}{
	{
		"http://localhost:8080/Den",
		"Greetings, Den",
	},
	{
		"http://localhost:8080/Kate",
		"Greetings, Kate",
	},
	{
		"http://localhost:8080/Ben",
		"Greetings, Ben",
	},
}

func TestClient(t *testing.T) {
	var cfg Config
	err := config.LoadConfig(&cfg)
	assert.NoError(t, err)

	c := client.NewClient(&cfg.Client)

	for _, ts := range testSuite {
		res, rErr := c.DoRequest(ts.url, http.MethodGet, client.Options{})
		assert.NoError(t, rErr)
		assert.Equal(t, string(res), ts.expectedResult)
	}
}
