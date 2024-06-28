package router

import (
	"github.com/brunobotter/ecommerce-carrinho/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)
		v1.GET("/carrinho/:id", handler.ShowCarrinhoHandler)
		v1.GET("/carrinho", handler.ListCarrinhoHandler)
		v1.POST("/carrinho", handler.CreateCarrinhoHandler)
	}
}
