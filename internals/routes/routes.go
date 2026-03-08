package routes

import (
	"net/http"

	"github.com/diaszakir/URL-Shortener-API-Go/internals/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", checkAPI)

	r.POST("/shorten", handlers.CreateShortURL)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// CheckAPI godoc
// @Summary Check API
// @Description Checking API and services working
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API works"})
}
