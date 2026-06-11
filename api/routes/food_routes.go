package routes

import (
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/api/handlers/foods"
	"github.com/gin-gonic/gin"
)

func registerFoodRoutes(routerGroup *gin.RouterGroup) {
	foodsRouterGroup := routerGroup.Group("/foods")
	foodsRouterGroup.GET("/getfoods", foods.GetAllFoods)
}
