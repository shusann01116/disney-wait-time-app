package tdr

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func TestFetchStandbys(t *testing.T) {
	urls := []string{
		"https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json",
		"https://www.tokyodisneyresort.jp/_/realtime/tds_attraction.json",
	}
	client := TdrClient{
		Ctx:    context.Background(),
		Client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}

	s, err := client.FetchStandbys(urls)
	assert.NoError(t, err)
	assert.NotEmpty(t, s)
}
