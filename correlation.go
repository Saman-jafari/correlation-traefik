package correlation_id_traefik

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	HeaderName string `yaml:"headerName,omitempty" json:"header_name,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		HeaderName: "",
	}
}

type Correlation struct {
	next       http.Handler
	name       string
	headerName string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.HeaderName == "" {
		config.HeaderName = "X-Correlation-Id"
	}

	return &Correlation{
		next:       next,
		name:       name,
		headerName: config.HeaderName,
	}, nil
}

func (c *Correlation) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	/*
			 0                   1                   2                   3
			 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
			+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
			|                           unix_ts_ms                          |
			+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
			|          unix_ts_ms           |  ver  |       rand_a          |
			+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
			|var|                        rand_b                             |
			+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
			|                            rand_b                             |
			+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		https://github.com/google/uuid/blob/master/version7.go
	*/
	var id, err = uuid.NewV7()
	correlationId := id.String()
	if request.Header.Get(c.headerName) != "" {
		correlationId = request.Header.Get(c.headerName)
	}
	if err == nil {
		request.Header.Set(c.headerName, correlationId)
	}

	c.next.ServeHTTP(writer, request)
}
