package main

import (
	"github.com/alfredamos/initializers"
	"github.com/alfredamos/middlewares"
	"github.com/alfredamos/routes"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariable()
	initializers.ConnectDB()
}


func main() {
	//----> Set the gin server.
	server := gin.Default()

	//----> Use the CORS middleware.
	server.Use(middlewares.CorsMiddleware())

	//---->Get the end-points
	routes.RegisteredRoutes(server)

	server.Run()
}