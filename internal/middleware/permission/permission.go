package permission

import (
	"net/http"

	user_jwt "product-service/internal/middleware/user-jwt"
	"product-service/internal/user"

	"github.com/gin-gonic/gin"
)

func NewMiddleware(roles ...user.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user from gin context
		user, err := user_jwt.GetUser(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "user not found",
			})
			return
		}

		// check if user has the required role
		for _, r := range roles {
			if user.Role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "insufficient role",
		})

		return
	}
}
