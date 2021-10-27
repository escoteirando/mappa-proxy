package mappa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guionardo/mappa_proxy/mappa/domain"
	"github.com/guionardo/mappa_proxy/mappa/repositories"
)

func Index(context *gin.Context) {
	context.JSON(200, gin.H{"mappa-proxy": "v1.0", "running-by": time.Since(StartedTime).String()})
}

func HealthCheck(context *gin.Context) {
	// Test mappa api
	statusCode, status, err := Ping()
	statusHealthy := "HEALTHY"
	if err != nil || statusCode < 1 || statusCode >= 400 {
		statusHealthy = "UNHEALTHY"
	}
	context.JSON(200, gin.H{"status": statusHealthy, "mappa_server": gin.H{"status_code": statusCode, "status": status, "memory": MemoryStatus()}})
}

func LoginStatsRoute(c *gin.Context) {
	stats := GetStats(repositories.GetRepository())
	c.JSON(200, stats)
}

func LoginRoute(c *gin.Context) {
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Failed to get request body: %s\n", err)
		c.JSON(400, gin.H{"message": "mAPPa request error", "error": err.Error()})
		return
	}
	var loginRequest domain.LoginRequest
	err = json.Unmarshal(requestBody, &loginRequest)
	if err != nil {
		log.Printf("Login request body is invalid: %s\n", err)
		c.JSON(400, gin.H{"message": "mAPPA request error", "error": err.Error()})
		return
	}
	loginResponse, ok := Login(loginRequest.UserName, loginRequest.Password)
	if !ok {
		c.JSON(403, gin.H{"message": "mAPPa login failed"})
		return
	}
	c.JSON(202, loginResponse)
}
