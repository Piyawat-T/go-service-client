package usecase

import (
	"time"

	"github.com/Piyawat-T/go-service-client/domain"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type sampleUsecase struct {
	contextTimeout time.Duration
}

func NewSampleUsecase(timeout time.Duration) domain.SampleUsecase {
	return &sampleUsecase{
		contextTimeout: timeout,
	}
}

func (usecase *sampleUsecase) GetSample(c *gin.Context) (domain.Sample, error) {
	ctx := c.Request.Context()
	log := otelzap.L()
	log.InfoContext(ctx, "Start Get Sample")

	sample := domain.Sample{
		Data: "sample",
	}
	return sample, nil
}
