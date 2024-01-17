package client

import (
	"context"
	"time"

	"github.com/Piyawat-T/go-service-client/bootstrap"
	"github.com/Piyawat-T/go-service-client/domain"
	"github.com/Piyawat-T/go-service-client/pkg"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

const samplePathV1 = "/v1/sample"

type sampleClient struct {
	Env            *bootstrap.Env
	contextTimeout time.Duration
}

func NewSampleClient(env *bootstrap.Env, timeout time.Duration) domain.SampleClient {
	return &sampleClient{
		Env:            env,
		contextTimeout: timeout,
	}
}

func (sc *sampleClient) Get(ctx context.Context) (domain.Sample, error) {
	log := otelzap.L()
	log.InfoContext(ctx, "Start Get Sample")

	url := viper.GetString(bootstrap.ServiceServerUrl) + samplePathV1
	sample, err := pkg.Get[domain.Sample](ctx, url)
	return sample, err
}
