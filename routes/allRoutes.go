package routes

import (
	"github.com/alfredamos/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(server *gin.Engine){
	//----> Unprotected routes.
	unAuthenticatedRoutes := server.Group("/api")
	
	unProtectedRoutes(unAuthenticatedRoutes)
	
	//----> Apply middleware for protected routes
	authenticatedRoutes := server.Group("/api").Use(middlewares.VerifyTokenJwt)

	//----> Protected routes.
	protectedRoutes(authenticatedRoutes)
	
	//----> Admin role permitted routes middleware.
	routesOfAdmin := server.Group("/api").Use(middlewares.VerifyTokenJwt, middlewares.RolePermission("Admin"))
	
	//----> Admin routes
	adminRoutes(routesOfAdmin) 
}