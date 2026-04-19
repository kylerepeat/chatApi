package router

import (
	"one-api/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	SetApiRouter(router)
	SetDashboardRouter(router)
	SetRelayRouter(router)
	router.NoRoute(controller.RelayNotFound)
}
