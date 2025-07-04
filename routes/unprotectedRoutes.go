package routes

import (
	"github.com/alfredamos/controllers"
	"github.com/gin-gonic/gin"
)

func unProtectedRoutes(server *gin.RouterGroup){
	//----> Auth-routes.
	server.POST("/auth/signup", controllers.SignupController)
	server.POST("/auth/login",controllers.LoginController)

	//----> Pizza-routes.
	server.GET("/menu-items", controllers.GetAllMenuItems)

}