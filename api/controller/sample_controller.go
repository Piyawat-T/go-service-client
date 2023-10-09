package controller

import (
	"net/http"

	"github.com/Piyawat-T/go-service-client/bootstrap"
	"github.com/Piyawat-T/go-service-client/domain"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type SampleController struct {
	Env           *bootstrap.Env
	SampleUsecase domain.SampleUsecase
}

func (s *SampleController) Sample(c *gin.Context) {
	ctx := c.Request.Context()
	log := otelzap.L()
	log.InfoContext(ctx, "Start Get Sample")
	sample, err := s.SampleUsecase.GetSample(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, sample)
}
