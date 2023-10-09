package main

import (
	"context"
	"fmt"
	"time"

	routeV1 "github.com/Piyawat-T/go-service-client/api/route/v1"
	"github.com/Piyawat-T/go-service-client/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()
	tp := bootstrap.NewTraceProvider()
	tc := tp.Tracer("service-client")
	bootstrap.NewContext(ctx, tc)

	app := bootstrap.App()
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second

	logger := otelzap.New(zap.Must(zap.NewDevelopment()),
		otelzap.WithTraceIDField(true),
		otelzap.WithMinLevel(zapcore.DebugLevel),
	)
	defer logger.Sync()
	undo := otelzap.ReplaceGlobals(logger)
	defer undo()

	gin.SetMode(env.GinMode)
	r := gin.Default()
	r.Use(otelgin.Middleware("service-name"))

	contextPath := fmt.Sprintf("%s/v1", env.ContextPath)
	routerV1 := r.Group(contextPath)
	routeV1.Setup(env, timeout, routerV1)

	r.Run(env.ServerAddress)
}
