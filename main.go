package main

import (
	"github.com/escoteirando/mappa-proxy/backend"
)

// @title          mappa-api
// @version        0.4.0
// @description    This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name   Guionardo Furlan
// @contact.email  guionardo@gmail.com
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath       /
func main() {
	backend.RunMappaProxy()
}
