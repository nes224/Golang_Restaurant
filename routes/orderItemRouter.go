package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/orderItems", controller.GetOrderItems())
	incomingRoute.GET("/orederItems/:orderItem_id", controller.GetOrderItem())
	incomingRoute.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoute.POST("/orderItems", controller.CreateOrderItem())
	incomingRoute.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

}