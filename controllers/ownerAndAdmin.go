package controllers

import (
	"net/http"
	"github.com/alfredamos/middlewares"
	"github.com/alfredamos/models"
	"github.com/gin-gonic/gin"
)

func OwnerAndAdmin(c *gin.Context)  {
	ord := models.Order{} //----> Order declaration and initiation.
	//----> Get the order id from param.
	id := c.Param("id")

	//----> Get user id from context.
	userIdFromAuth := middlewares.GetUserIdFromContext(c)

	//----> Get the order with the given id
	order, _ := ord.GetOrderById(id)
	userId := order.ID //----> User id attached to order.

	//----> Check for equality of userId.
	isOwner := middlewares.IsSameUser(userIdFromAuth, userId) 

	//----> Get admin user.
	_, isAdmin := middlewares.GetUserAuthFromContext(c)

	//----> Only owner and admin is allowed to pass.
	if !isAdmin && !isOwner {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "You are not authorized on this page!"})
		return 
	}

	//----> Admin and owner can pass.
	c.Next()

}