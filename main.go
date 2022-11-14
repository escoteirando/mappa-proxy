package main

import (
	"log"

	"github.com/escoteirando/mappa-proxy/backend"
	"github.com/escoteirando/mappa-proxy/backend/build"
	"github.com/escoteirando/mappa-proxy/backend/configuration"
	"github.com/escoteirando/mappa-proxy/backend/mappa"
)

// @title          mappa-api
// @version        0.5.1
// @description    This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name   Guionardo Furlan
// @contact.email  guionardo@gmail.com
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath       /

var Build = "development"

func main() {
	build.BuildTime = Build
	log.Printf("Starting %s v%s - %s @ %v", configuration.APP_NAME, configuration.APP_VERSION, mappa.URL, build.BuildTime)
	backend.RunMappaProxy()
}
