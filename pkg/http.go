package pkg

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Get[T any](ctx context.Context, url string) (T, error) {
	log := otelzap.L()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	var response T
	if err != nil {
		return response, err
	}
	httpclient := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := httpclient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	log.DebugContext(ctx, resp.Status)
	json.Unmarshal(b, &response)
	return response, err
}
