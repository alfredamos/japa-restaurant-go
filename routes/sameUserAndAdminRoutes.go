package routes

import (
	"github.com/alfredamos/controllers"
	"github.com/gin-gonic/gin"
)

func sameUserAndAdminRoutes(r gin.IRoutes) {
	//----> Order routes.
	r.GET("/orders/orders-by-user-id/:userId", controllers.GetAllOrderByUserId)
	r.DELETE("/orders/delete-all-orders-by-user-id/:userId", controllers.DeleteOrderByUserId)

	//----> User routes.
	r.DELETE("/users/:id", controllers.DeleteUserById)
	r.GET("/users/:id", controllers.GetUserById)
}