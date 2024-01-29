package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/orders", controller.GetOrders())
	incomingRoute.GET("/orders/:order_id", controller.GetOrder())
	incomingRoute.POST("/orders", controller.CreateOrder())
	incomingRoute.PATCH("/orders/:order_id", controller.UpdateOrder())
}