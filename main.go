package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guionardo/mappa_proxy/mappa"
	"github.com/guionardo/mappa_proxy/mappa/logging"
	"github.com/guionardo/mappa_proxy/tg"
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
	store := persistence.NewInMemoryStore(time.Minute * 60)

	mappa.StartMappa()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "User-Agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Static("/web", "./web")
	r.GET("/", index)
	r.GET("/hc", healthCheck)
	r.GET("/login/stats", mappa.LoginStatsRoute)
	r.POST("/mappa/login", mappa.LoginRoute)
	r.GET("/mappa/*request", cache.CachePage(store, time.Minute*60, mappa.GetRequest))
	r.POST("/tg/pub", tg.Publish)
	return r
}

func index(context *gin.Context) {
	context.JSON(200, gin.H{"mappa-proxy": "v1.0", "running-by": time.Since(mappa.StartedTime).String()})
}
func parseHostAndPort() (string, int) {
	var nPort = flag.Int("port", 0, "HTTP port")
	var sHost = flag.String("host", "0.0.0.0", "HTTP host")

	flag.Parse()
	port := 0
	if *nPort > 0 {
		port = *nPort
	} else {
		sPort := os.Getenv("HTTP_PORT")
		if len(sPort) == 0 {
			port = 8081
		} else {
			ePort, err := strconv.Atoi(sPort)
			if err != nil {
				log.Fatalf("Invalid HTTP_PORT environment:  %v", sPort)
			}
			port = ePort
		}
	}
	return *sHost, port
}
func main() {
	logger, err := logging.New("logs/mappa-proxy.log", 2)
	if err != nil {
		log.Panicf("Failed to start logging - %v", err)
	}
	log.SetOutput(logger)
	log.Printf("STARTING %v", time.Now())
	host, port := parseHostAndPort()

	err = setupServer().Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
