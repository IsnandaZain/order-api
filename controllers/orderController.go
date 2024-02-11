package controllers

import (
	"fmt"
	"net/http"
	"order-api/database"
	"order-api/models"

	"github.com/gin-gonic/gin"
)

// GetOrder godoc
// @Summary Get details
// @Description Get details of one order
// @Tags orders
// @Accept json
// @Produce json
// @Param Id path int tru "ID for the order"
// @Success 200 {object} models.Order
// @Router /orders/{orderID} [get]
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

// GetAllOrder godoc
// @Summary Get list
// @Description Get list of all order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} models.Order
// @Router /orders [get]
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

// CreateOrder godoc
// @Summary Post details for a given Id
// @Description Post details of order corresponding to the input Id
// @Tags orders
// @Accept json
// @Produce json
// @Param models.Order body models.Order true "create order"
// @Success 200 {object} models.Order
// @Router /orders [post]
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

// UpdateOrder godoc
// @Summary Update details for a given Id
// @Description Update details of order corresponding to the input Id
// @Tags orders
// @Accept json
// @Produce json
// @Param models.Order body models.Order true "update order"
// @Success 200 {object} models.Order
// @Router /orders [put]
func UpdateOrder(ctx *gin.Context) {
	/*
		Mekanisme update :
		- update data order
		- delete item yang melekat pada order
		- insert item baru yang dikirim melalui endpoint update
	*/
	orderID := ctx.Param("orderID")
	db := database.GetDB()

	orderUpdate := models.Order{}
	if err := ctx.ShouldBindJSON(&orderUpdate); err != nil { // validasi input
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// make sure order exist
	var order models.Order
	err := db.Preload("Items").First(&order, "ID=?", orderID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  fmt.Sprintf("Failed get data order with id %v", orderID),
			"error_message": fmt.Sprintf("Failed get data order with id %v", orderID),
		})
		return
	}

	// make sure item exist
	var orderItem models.Item
	err = db.First(&orderItem, "order_id=?", orderID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  fmt.Sprintf("Failed get data item for order id %v", orderID),
			"error_message": fmt.Sprintf("Failed get data item for order id %v", orderID),
		})
		return
	}

	// update order
	db.Model(&order).Where("ID=?", orderID).Updates(models.Order{
		CustomerName: orderUpdate.CustomerName,
	})

	// update item
	db.Delete(&orderItem)
	err = db.Create(&orderUpdate.Items).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Error insert item data",
			"error_message": "Something wrong when try to insert data",
		})
		return
	}

	orderUpdate.ID = order.ID
	ctx.JSON(http.StatusCreated, gin.H{
		"order": orderUpdate,
	})
}

// DeleteOrder godoc
// @Summary Delete order for a given Id
// @Description Delete order corresponding to the param Id
// @Tags orders
// @Accept json
// @Produce json
// @Param Id path int tru "ID for the order"
// @Success 200 "{'message': 'Order has been deleted'}"
// @Router /orders/{orderID} [delete]
func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	db := database.GetDB()

	var order models.Order
	err := db.Preload("Items").First(&order, "ID=?", orderID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed get data order with items",
			"error_message": fmt.Sprintf("Failed get data order with id %v", orderID),
		})
		return
	}

	// delete items in produk
	var orderItem models.Item
	errItem := db.First(&orderItem, "order_id=?", orderID).Error
	if errItem != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  fmt.Sprintf("Failed get data item for id order %v", orderID),
			"error_message": fmt.Sprintf("Failed get data item for id order %v", orderID),
		})
		return
	}

	db.Delete(&orderItem)

	// delete produk
	db.Delete(&order)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %v has been deleted", orderID),
	})
}
