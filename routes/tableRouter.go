package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/tables", controller.GetTables())
	incomingRoute.GET("/tables/:table_id", controller.GetTable())
	incomingRoute.POST("/tables", controller.CreateTable())
	incomingRoute.PATCH("/tables/:table_id", controller.UpdateTable())
}