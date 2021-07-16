package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guionardo/mappa_proxy/mappa"
)

func healthCheck(context *gin.Context) {
	// Test mappa api
	statusCode, status, err := mappa.Ping()
	statusHealthy := "HEALTHY"
	if err != nil || statusCode < 1 || statusCode >= 400 {
		statusHealthy = "UNHEALTHY"
	}
	context.JSON(200, gin.H{"status": statusHealthy, "mappa_server": gin.H{"status_code": statusCode, "status": status}})

}

func setupServer() *gin.Engine {
	r := gin.Default()
	mappa.StartMappa()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "User-Agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/hc", healthCheck)
	r.GET("/login/stats", mappa.MappaLoginStatsRoute)
	r.POST("/mappa/login", mappa.MappaLoginRoute)
	r.GET("/mappa/*request", mappa.MappaGetRequest)
	// r.POST("/mappa/*request", mappa.MappaPostRequest)
	return r
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	setupServer().Run(":" + port)
}
