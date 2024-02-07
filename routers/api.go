package routers

import (
	"order-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)

	router.GET("/orders", controllers.GetAllOrder)

	router.GET("/orders/:orderID", controllers.GetOrder)

	// router.PUT("/orders/:orderID", controllers.UpdateOrder)

	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
