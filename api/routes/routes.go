package routes

import "github.com/gin-gonic/gin"

func RegisterApiRoutes(router *gin.Engine) {
	api := router.Group("/api")
	v1 := api.Group("/v1")
	{
		registerCustomerRoutes(v1)
	}
}
