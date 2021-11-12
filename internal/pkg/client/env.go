package client

// Config keeps configuration of HTTP server.
type Config struct {
	HttpAddress      string `env:"HTTP_ADDRESS" envDefault:"http://localhost:8080/test"`
	LogRequests      bool   `env:"SERVER_LOG_REQUESTS" envDefault:"true"`
	LogRequestBody   bool   `env:"SERVER_LOG_REQUEST_BODY" envDefault:"true"`
	ClientTimeOut    int    `env:"CLIENT_TIMEOUT" envDefault:"10"`
	MaxConn          int    `env:"MAX_CONN" envDefault:"1024"`
	HandshakeTimeout int    `env:"HANDSHAKE_TIMEOUT" envDefault:"0"`
}
