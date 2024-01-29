package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)

func MenuRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/menus", controller.GetMenus())
	incomingRoute.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoute.POST("/menus", controller.CreateMenu())
	incomingRoute.PATCH("/menus/:menu_id", controller.UpdateMenu())
}