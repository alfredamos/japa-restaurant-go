package routes

import (
)

func unProtectedRoutes(server *gin.RouterGroup){
	//----> Auth-routes.
	server.POST("/auth/signup", controllers.SignupController)
	server.POST("/auth/login",controllers.LoginController)

	//----> Pizza-routes.
	//server.GET("/pizzas", controllers.GetAllPizza)

}