package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SameUserAndAdmin(c *gin.Context) {
	//----> Get the user-id .
	userId := c.Param("userId")
	//----> Get user id from context.
	userIdInt := GetUserIdFromContext(c)

	//----> Check for equality of userId.
	userIsSame := IsSameUser(userIdInt, userId) 

	//----> Get admin user.
	_, isAdmin := GetUserAuthFromContext(c)

	//----> Admin and same user are not allowed.
	if !isAdmin && !userIsSame {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "You are not authorized on this page!"})
		return 
	}

	//----> You are either an admin or same user.
	c.Next()
}

//----> Check for checking for same user.
func IsSameUser(userId1, userId2 string) bool{
	return userId1 == userId2
}

