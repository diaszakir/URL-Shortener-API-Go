package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", checkAPI)

	r.POST("/shorten")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API works"})
}
