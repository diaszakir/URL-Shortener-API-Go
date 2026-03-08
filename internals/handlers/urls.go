package handlers

import (
	"net/http"
	"sync/atomic"

	"github.com/diaszakir/URL-Shortener-API-Go/internals/models"
	"github.com/diaszakir/base62-go"
	"github.com/gin-gonic/gin"
)

const baseURL = "http://localhost:8080"

var nextID int64 = 0

var urlStore = make(map[string]*models.ShortURL)

// CreateShortURL godoc
// @Summary Creating short url
// @Description Creating a new short url
// @Tags urls
// @Produce json
// @Accept json
// @Param todo body models.URLRequest true "URL payload"
// @Success 200 {object} map[string]string
// @Router /shorten [post]
func CreateShortURL(c *gin.Context) {

	var url models.URLRequest
	err := c.ShouldBindJSON(&url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect data"})
		return
	}

	id := atomic.AddInt64(&nextID, 1)
	maskedID := id*7919 + 1000000
	code, err := base62.Encode(maskedID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if len(code) > 1 {
		code = code[1:]
	}

	var res models.URLResponse

	res = models.URLResponse{
		OriginalURL: url.URL,
		ShortCode:   code,
		ShortURL:    baseURL + "/" + code,
	}

	urlStore[code] = &models.ShortURL{
		Code:        code,
		OriginalURL: res.OriginalURL,
		Clicks:      0,
	}

	c.JSON(http.StatusCreated, res)
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	url, ok := urlStore[code]

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	urlStore[code].Clicks++

	c.Redirect(http.StatusFound, url.OriginalURL)
}

func InfoURL(c *gin.Context) {
	code := c.Param("code")

	url, ok := urlStore[code]

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, url)
}
