package route

import (
	"time"

	"github.com/Piyawat-T/go-service-client/api/controller"
	"github.com/Piyawat-T/go-service-client/bootstrap"
	"github.com/Piyawat-T/go-service-client/client"
	"github.com/Piyawat-T/go-service-client/usecase"
	"github.com/gin-gonic/gin"
)

func SampleRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	client := client.NewSampleClient(env, timeout)
	sc := controller.SampleController{
		Env:           env,
		SampleUsecase: usecase.NewSampleUsecase(timeout, client),
	}
	group.GET("/sample", sc.Sample)
}
