package routes

import (
	"github.com/alfredamos/controllers"
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

	//----> Same user and admin routes.
	routesOfSameUserAndAdmin := server.Group("/api").Use(middlewares.VerifyTokenJwt, middlewares.SameUserAndAdmin)
	sameUserAndAdminRoutes(routesOfSameUserAndAdmin)

	//----> Owner and admin routes.
	routesOfOwnerAndAdmin := server.Group("/api").Use(middlewares.VerifyTokenJwt, controllers.OwnerAndAdmin)
	ownerAndAdminRoutes(routesOfOwnerAndAdmin)
	
	//----> Admin role permitted routes middleware.
	routesOfAdmin := server.Group("/api").Use(middlewares.VerifyTokenJwt, middlewares.RolePermission("Admin"))
	
	//----> Admin routes
	adminRoutes(routesOfAdmin) 
}