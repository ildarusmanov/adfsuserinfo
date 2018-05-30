package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/adfsuserinfo/controllers"
	"github.com/ildarusmanov/adfsuserinfo/interfaces"
)

func CreateRouter(serviceLocator interfaces.ServiceLocator) *gin.Engine {
	userinfoC := controllers.CreateUserinfoController(
		serviceLocator.Get("Userinfo").(interfaces.UserinfoService),
	)

	statusC := controllers.CreateStatusController()

	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/me", userinfoC.Get)
		apiv1.GET("/status/check", statusC.Check)
	}

	return r
}
