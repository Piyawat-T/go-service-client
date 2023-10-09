package usecase

import (
	"context"
	"time"

	"github.com/Piyawat-T/go-service-client/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type sampleUsecase struct {
	contextTimeout time.Duration
	sampleClient   domain.SampleClient
}

func NewSampleUsecase(timeout time.Duration, sampleClient domain.SampleClient) domain.SampleUsecase {
	return &sampleUsecase{
		contextTimeout: timeout,
		sampleClient:   sampleClient,
	}
}

func (usecase *sampleUsecase) GetSample(ctx context.Context) (domain.Sample, error) {
	log := otelzap.L()
	log.InfoContext(ctx, "Start Get Sample")
	return usecase.sampleClient.Get(ctx)
}
