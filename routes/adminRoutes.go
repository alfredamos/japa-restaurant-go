package routes

import (
	"github.com/alfredamos/controllers"
	"github.com/gin-gonic/gin"
)

func adminRoutes(p gin.IRoutes){
	//----> Orders routes.
	p.GET("/orders", controllers.GetAllOrders)
	p.DELETE("/orders/delete-all-orders", controllers.DeleteAllOrders)

	//----> Menu-item routes.
	p.POST("/menu-items", controllers.CreateMenuItem)
	p.DELETE("/menu-items/:id", controllers.DeleteMenuItemById)
	p.PATCH("/menu-items/:id", controllers.EditMenuItemById)
	
	//----> User routes.
	p.GET("/users", controllers.GetAllUsers)
	p.DELETE("/users/:id", controllers.DeleteUserById)
}