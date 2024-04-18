package main

import "github.com/gin-gonic/gin"

func main() {
	// setup gin server
	r := gin.Default()

	// setup handler (things should be a lot more complex here, but for simplicity's sake we'll leave
	// it as simple as this
	handler := NewHandler()

	// setup api routes
	r.POST("/product", handler.CreateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)
	r.POST("/product/:id/need-admin-notification", handler.SetNeedAdminNotification)

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
