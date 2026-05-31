package routes

import (
	handlers "github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/api/handlers/customers"
	"github.com/gin-gonic/gin"
)

func registerCustomerRoutes(routerGroup *gin.RouterGroup) {
	customers := routerGroup.Group("/customers")
	customers.GET("/getall", handlers.GetAllCustomers)
}
