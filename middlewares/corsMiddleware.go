package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

// CORS middleware function definition
func CorsMiddleware() gin.HandlerFunc {
	// Define allowed origins as a comma-separated string
	originsString := "http://localhost:3000,http://localhost:4200,http://localhost:5173,http://localhost:5174"
	var allowedOrigins []string
	if originsString != "" {
	 // Split the originsString into individual origins and store them in allowedOrigins slice
	 allowedOrigins = strings.Split(originsString, ",")
	}
 
	// Return the actual middleware handler function
	return func(c *gin.Context) {
	 //----> Get the Origin header from the request
	 origin := GetOrigin(c)

	 //----> Function to check if a given origin is allowed
	 isOriginAllowed := getAllAllowedOrigins(origin, allowedOrigins)
 
	 //----> Check if the origin is allowed
	 if isOriginAllowed {
		//----> If the origin is allowed, set CORS headers in the response
		setCorsHeaders(c, origin)
	 }

	 fmt.Println("origin : ", origin)
 
	 fmt.Println("isOriginAllowed : ", isOriginAllowed)
 
	 //----> Handle preflight OPTIONS requests by aborting with status 204
	 if c.Request.Method == "OPTIONS" {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"message": "wrongly configured"})
		return
	 }
 
	 //----> Call the next handler
	 c.Next()
	}
 }

 func getAllAllowedOrigins(origin string, allowedOrigins []string) bool{
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
		 return true
		}
	 }
	 return false
 }

 func setCorsHeaders(c *gin.Context, origin string){
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, DELETE, OPTIONS, GET, PUT")
 }

 func GetOrigin(c *gin.Context) string{
	return c.Request.Header.Get("Origin")
 }