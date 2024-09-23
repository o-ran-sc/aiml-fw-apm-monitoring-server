package api

import (
	"github.com/gin-gonic/gin"
	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-server/pkg/api/monitoring"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	monitoringServer monitoring.Command
)

func init() {
	monitoringServer = monitoring.NewServer()
}

func setupRouter() (router *gin.Engine) {
	router = gin.Default()

	v1 := router.Group("/v1/monitoring")
	{
		monitoringServer.Register(v1)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return
}

// ListenAndServe is listening rest api call & setting up serving function
func ListenAndServe() {
	router := setupRouter()
	router.Run()
}
