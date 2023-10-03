package route

import (
	"time"

	"github.com/Piyawat-T/go-service-client/bootstrap"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")

	HomeRouter(env, timeout, publicRouterV1)
	SampleRouter(env, timeout, publicRouterV1)
}
