package controllers

import (
	"fmt"
	"net/http"
	"order-api/database"
	"order-api/models"

	"github.com/gin-gonic/gin"
)

func GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	db := database.GetDB()

	order := models.Order{}
	err := db.Preload("Items").First(&order, "ID=?", orderID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed get data order with items",
			"error_message": fmt.Sprintf("Failed get data order with items with id %v", orderID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func GetAllOrder(ctx *gin.Context) {
	db := database.GetDB()

	orders := models.Order{}
	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed get data order with items",
			"error_message": fmt.Sprintf("Failed get data order"),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	newOrder := models.Order{}

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&newOrder).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Error insert data",
			"error_message": "Something wrong when try to insert data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

// func UpdateOrder(ctx *gin.Context) {
// 	orderID := ctx.Param("orderID")
// 	db := database.GetDB()

// 	orderUpdate := models.Order{}
// 	if err := ctx.ShouldBindJSON(&orderUpdate); err != nil { // validasi input
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	// delete all items
// }

func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")

}
