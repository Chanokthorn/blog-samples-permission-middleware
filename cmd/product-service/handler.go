package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) CreateProduct(c *gin.Context) {
	println("Product created")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product created",
	})
	return
}

func (h Handler) DeleteProduct(c *gin.Context) {
	println("Product deleted")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
	return
}

func (h Handler) SetNeedAdminNotification(c *gin.Context) {
	println("Need admin notification set")
	c.JSON(http.StatusOK, gin.H{
		"message": "Need admin notification set",
	})
	return
}
