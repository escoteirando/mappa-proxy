package handlers

import (
	"context"

	"github.com/escoteirando/mappa-proxy/backend/cache"
	"github.com/escoteirando/mappa-proxy/backend/configuration"
	"github.com/escoteirando/mappa-proxy/backend/mappa"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
	"github.com/gofiber/fiber/v2"
)

const userContextKey = "uc_key"

type MappaUserContextData struct {
	Config        configuration.Configuration
	Cache         *cache.MappaCache
	MappaService  *mappa.MappaService
	Authorization string
}

func getUserContext(c *fiber.Ctx) *MappaUserContextData {
	if userContext := c.UserContext(); userContext != nil {
		anyValue := userContext.Value(userContextKey)
		switch f := anyValue.(type) {
		case *MappaUserContextData:
			auth := string(c.Request().Header.Peek("Authorization"))
			f.Authorization = auth
			return f
		}
	}
	return nil
}

func SetupUserContext(app *fiber.App, config configuration.Configuration, cache *cache.MappaCache, repository repositories.IRepository) {
	mappaUserContextData := &MappaUserContextData{
		Config: config,
		Cache:  cache,
		MappaService: &mappa.MappaService{
			Cache:      cache,
			Repository: repository,
			API:        mappa.NewMappaAPI(),
		},
	}

	app.Use(func(c *fiber.Ctx) error {
		userContext := context.WithValue(c.Context(), userContextKey, mappaUserContextData)
		c.SetUserContext(userContext)
		return c.Next()
	})
}
