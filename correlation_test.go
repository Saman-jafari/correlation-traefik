package correlation_id_traefik_test

import (
	"context"
	"github.com/saman-jafari/correlation-id-traefik"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemo(t *testing.T) {
	cfg := correlation_id_traefik.CreateConfig()
	cfg.HeaderName = "correlation-id"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := correlation_id_traefik.New(ctx, next, cfg, "plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertCorrelationId(t, req, "correlation-id")

}

func assertCorrelationId(t *testing.T, req *http.Request, key string) {
	t.Helper()

	if req.Header.Get(key) == "" {
		t.Errorf("correlation id should exists")
	}
}
