// controllers/books.go

package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/dalekurt/go-assets-api-microservice/models"
)

// GET /assets
// Get all assets
func FindAssets(c *gin.Context) {
	var assets []models.Asset
	models.DB.Find(&assets)

	c.JSON(http.StatusOK, gin.H{"data": assets})
}
