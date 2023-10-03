package domain

import "github.com/gin-gonic/gin"

type Sample struct {
	Data string `json:"data"`
}

type SampleUsecase interface {
	GetSample(c *gin.Context) (Sample, error)
}
