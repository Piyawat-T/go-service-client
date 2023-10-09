package domain

import (
	"context"
)

type Sample struct {
	Data string `json:"data"`
}

type SampleClient interface {
	Get(c context.Context) (Sample, error)
}

type SampleUsecase interface {
	GetSample(c context.Context) (Sample, error)
}
