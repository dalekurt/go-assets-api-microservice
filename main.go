package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/dalekurt/go-assets-api-microservice/models"
  "github.com/dalekurt/go-assets-api-microservice/controllers"
)

func main() {
  router := gin.Default()

  models.ConnectDatabase()

  router.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Healthy"})    
  })

  router.GET("/", controllers.FindAssets)

  router.Run()
}
