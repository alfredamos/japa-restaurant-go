package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckForSameUserAndAdmin(userId string, c *gin.Context) error {
		//----> Get user id from context.
		userIdInt := GetUserIdFromContext(c)

		//----> Check for equality of userId.
		userIsSame := IsSameUser(userIdInt, userId) 

		//----> Get admin user.
		_, isAdmin := GetRoleFromContext(c)

		//----> Admin and same user are allowed.
		if isAdmin || userIsSame {
			return nil
		}

		//----> You are not admin neither is same user, hence you are not allowed.
		return errors.New("you are not permitted to access this page or perform this function")

}

//----> Check for checking for same user.
func IsSameUser(userId1, userId2 string) bool{
	return userId1 == userId2
}

