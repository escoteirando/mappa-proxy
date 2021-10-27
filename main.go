package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guionardo/mappa_proxy/mappa"
	"github.com/guionardo/mappa_proxy/mappa/configuration"
	"github.com/guionardo/mappa_proxy/mappa/logging"
	"github.com/guionardo/mappa_proxy/mappa/repositories"
	"github.com/guionardo/mappa_proxy/tg"
)

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
	r.GET("/", mappa.Index)
	r.GET("/hc", mappa.HealthCheck)
	r.GET("/login/stats", mappa.LoginStatsRoute)
	r.POST("/mappa/login", mappa.LoginRoute)
	r.GET("/mappa/associado/*id")
	r.GET("/mappa/*request", cache.CachePage(store, time.Minute*60, mappa.GetRequest))
	r.POST("/tg/pub", tg.Publish)
	return r
}

func main() {
	config, err := configuration.GetConfiguration()
	if err != nil {
		log.Panicf("Failed to read configuration - %v", err)
	}
	repositories.SetRepository(config.Repository)
	logger, err := logging.New(filepath.Join(config.LogFolder, "mappa-proxy.log"), 5)
	if err != nil {
		log.Panicf("Failed to start logging - %v", err)
	}
	log.SetOutput(logger)
	log.Printf("STARTING Log = %s", logger.GetFile())

	err = setupServer().Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
