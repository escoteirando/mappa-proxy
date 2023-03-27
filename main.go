package main

import (
	"log"

	"github.com/escoteirando/mappa-proxy/backend"
	"github.com/escoteirando/mappa-proxy/backend/build"
	"github.com/escoteirando/mappa-proxy/backend/configuration"
	"github.com/escoteirando/mappa-proxy/backend/mappa"
)

// @title         mappa-proxy
// @version       0.5.4
// @description   Proxy and data analysis for Mappa
// @contact.name  Guionardo Furlan
// @contact.email guionardo@gmail.com
// @BasePath      /

var Build = "development"

func main() {
	build.BuildTime = Build
	log.Printf("Starting %s v%s - %s @ %v", configuration.APP_NAME, configuration.APP_VERSION, mappa.URL, build.BuildTime)
	backend.RunMappaProxy()

}
