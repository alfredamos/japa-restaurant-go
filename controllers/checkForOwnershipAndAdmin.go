package controllers

import (
	"github.com/alfredamos/middlewares"
	"github.com/alfredamos/models"
	"github.com/gin-gonic/gin"
)

func checkForOwnershipAndAdmin(id string, c *gin.Context) (bool, bool) {
	ord := models.Order{} //----> Order declaration and initiation.
	
	//----> Get user id from context.
	userIdFromAuth := middlewares.GetUserIdFromContext(c)

	//----> Get the order with the given id
	order, _ := ord.GetOrderById(id)
	userId := order.ID //----> User id attached to order.

	//----> Check for equality of userId.
	isOwner := middlewares.IsSameUser(userIdFromAuth, userId) 

	//----> Get admin user.
	_, isAdmin := middlewares.GetRoleFromContext(c)

	//----> Send back results.
	return isAdmin, isOwner

}