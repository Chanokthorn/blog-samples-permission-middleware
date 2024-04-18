package main

import (
	"product-service/internal/user"

	"product-service/internal/middleware/permission"
	user_jwt "product-service/internal/middleware/user-jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	// setup gin server
	r := gin.Default()

	r.Use(user_jwt.NewMiddleware("some-jwt-secret"))

	// setup handler (things should be a lot more complex here, but for simplicity's sake we'll leave
	// it as simple as this
	handler := NewHandler()

	// setup api routes
	r.POST(
		"/product",
		permission.NewMiddleware(user.RoleMember, user.RoleAdmin),
		handler.CreateProduct,
	)
	r.DELETE(
		"/product/:id",
		permission.NewMiddleware(user.RoleAdmin),
		handler.DeleteProduct,
	)
	r.POST(
		"/product/:id/need-admin-notification",
		permission.NewMiddleware(user.RoleMember),
		handler.SetNeedAdminNotification,
	)

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
