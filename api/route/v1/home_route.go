package route

import (
	"time"

	"github.com/Piyawat-T/go-service-client/api/controller"
	"github.com/Piyawat-T/go-service-client/bootstrap"
	"github.com/gin-gonic/gin"
)

func HomeRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	sc := controller.HomeController{
		Env: env,
	}
	group.GET("/ping", sc.Home)
}
