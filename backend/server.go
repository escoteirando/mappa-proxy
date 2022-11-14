package backend

import (
	"fmt"

	"log"

	"github.com/escoteirando/mappa-proxy/backend/app"
	"github.com/escoteirando/mappa-proxy/backend/cache"
	"github.com/escoteirando/mappa-proxy/backend/configuration"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
)

func RunMappaProxy() {
	configuration.Init()
	repository := repositories.RepositoryFactory.GetRepository(configuration.Config.Repository)
	if repository == nil {
		log.Fatalf("Repository is not configured: %s", configuration.Config.Repository)
	}
	cache, err := cache.CreateMappaCache(repository)
	if err != nil {
		log.Fatalf("Failed to create cache %v", err)
	}

	server, err := app.CreateServer(*configuration.Config, cache, repository)
	if err != nil {
		log.Fatalf("Failed to create http server %v", err)
	} else {
		server.Listen(fmt.Sprintf(":%d", configuration.Config.Port))
	}
}
