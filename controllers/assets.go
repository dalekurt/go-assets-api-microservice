// controllers/books.go

package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/dalekurt/go-assets-api-microservice/models"
)

// GET /
// Get all assets
func FindAssets(c *gin.Context) {
	var assets []models.Asset
	models.DB.Find(&assets)

	c.JSON(http.StatusOK, gin.H{"data": assets})
}

// POST /
// Create a new asset
func CreateAsset(c *gin.Context) {
	var input models.CreateAssetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create asset
	createAsset := models.Asset{Filename: input.Filename, Filesize: input.Filesize, Filetype: input.Filetype}
	models.DB.Create(&createAsset)

	c.JSON(http.StatusOK, gin.H{"data":createAsset})
}

// GET /:id
// Find an asset
func FindAsset(c *gin.Context){
	var asset models.Asset

	if err := models.DB.Where("id = ?", c.Param("id")).First(&asset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Asset not found!"})
    	return
	}

	c.JSON(http.StatusOK, gin.H{"data": asset})
}

// PATCH /assets/:id
// Update an asset
func UpdateAsset(c *gin.Context) {
	// Get model if exist
	var asset models.Asset

	if err := models.DB.Where("id = ?", c.Param("id")).First(&asset).Error; err != nil {
	  c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Asset not found!"})
	  return
	}
  
	// Validate input
	var input models.UpdateAssetInput

	if err := c.ShouldBindJSON(&input); err != nil {
	  c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
	
	updateAsset := models.Asset{Filename: input.Filename, Filesize: input.Filesize, Filetype: input.Filetype}

	models.DB.Model(&asset).Updates(updateAsset)
  	c.JSON(http.StatusOK, gin.H{"data": updateAsset})
}
