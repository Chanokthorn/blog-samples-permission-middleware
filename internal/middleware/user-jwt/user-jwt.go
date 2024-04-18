package user_jwt

import (
	"errors"
	"net/http"
	"product-service/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const userKey = "user"

type Claims struct {
	ShopOwnerID int       `json:"shopOwnerID"`
	Role        user.Role `json:"role"`
	jwt.RegisteredClaims
}

func NewMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userJWT := c.GetHeader("x-user-jwt")
		claims := &Claims{}

		_, err := jwt.ParseWithClaims(userJWT, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unable to parse claims",
			})
			return
		}

		c.Set(userKey, user.User{
			ShopOwnerID: claims.ShopOwnerID,
			Role:        claims.Role,
		})
		c.Next()
	}
}

func GetUser(c *gin.Context) (user.User, error) {
	u, ok := c.Get(userKey)
	if !ok {
		return user.User{}, errors.New("user not found")
	}

	uParsed, ok := u.(user.User)
	if !ok {
		return user.User{}, errors.New("unable to parse user from context")
	}

	return uParsed, nil
}
