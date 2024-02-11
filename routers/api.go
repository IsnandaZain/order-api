package routers

import (
	"order-api/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "order-api/docs"
)

// @title Order API
// @version 1.0
// @description Thisi s a simple service for Manage Order
// @termOfService https://swagger.io/terms
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://ww.apache.org/
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)

	router.GET("/orders", controllers.GetAllOrder)

	router.GET("/orders/:orderID", controllers.GetOrder)

	router.PUT("/orders/:orderID", controllers.UpdateOrder)

	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
