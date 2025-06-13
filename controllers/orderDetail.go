package controllers

import (
	"fmt"
	"net/http"
	"github.com/alfredamos/models"
	"github.com/gin-gonic/gin"
)

func CreateOrderDetail(context *gin.Context) {
	orderDetail := models.OrderDetail{} //----> OrderDetail variable
	
	//----> Get the cart-item payload from the request.
	err := context.ShouldBindJSON(&orderDetail)

	//----> Check for binding error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Insert the cart-item into the database.
	err = orderDetail.CreateOrderDetail()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusBadRequest, gin.H{"status": "Success", "message": "Cart-item has been created successfully!"})
}

func DeleteOrderDetailById(context *gin.Context) {
	orderDetail := models.OrderDetail{} //----> Cart-item variable.
	
	//----> Get the id from param.
	id := context.Param("id")
	
	//----> Delete the cart-item from the database.
	err := orderDetail.DeleteOrderDetailById(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Cart-item has been deleted successfully!"})
}

func EditOrderDetailById(context *gin.Context) {
	orderDetail := models.OrderDetail{} //----> Cart-item variable.

	//----> Get the id from param.
	id := context.Param("id")

	//----> Get the request payload
	err := context.ShouldBindJSON(&orderDetail)

	//----> Check for error.
	if err != nil {
	 context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	 return
	}

	//----> Update cart-item in the database.
	err = orderDetail.EditOrderDetailId(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Cart-item has been edited successfully!"})
}

func GetAllOrderDetails(context *gin.Context) {
	orderDetail := models.OrderDetail{} //----> Cart-item variable.

	//----> Retrieve all the cart-items from database.
	orderDetails, err := orderDetail.GetAllOrderDetails()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back response.
	context.JSON(http.StatusOK, orderDetails)
}

func GetOrderDetailById(context *gin.Context) {
	orderDetail := models.OrderDetail{} //----> Cart-item variable.
	
	//----> Get the id from param.
	id := context.Param("id")

	//----> Retrieve cart-item from database.
	orderDetail, err := orderDetail.GetOrderDetailById(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusOK, orderDetail)
}