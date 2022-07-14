package app

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/app/handlers"
	"github.com/escoteirando/mappa-proxy/backend/app/middleware"
	"github.com/escoteirando/mappa-proxy/backend/cache"
	"github.com/escoteirando/mappa-proxy/backend/configuration"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
	"github.com/gofiber/fiber/v2"
	fiberCache "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func CreateServer(config configuration.Configuration, cache *cache.MappaCache, repository repositories.IRepository) (app *fiber.App, err error) {
	app = fiber.New()
	handlers.SetupUserContext(app, config, cache, repository)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST,GET",
		AllowHeaders:     "Origin,Authorization,Content-Type,User-Agent",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	if config.HttpCacheTime > 0 {
		app.Use(fiberCache.New(fiberCache.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.Path() == "/mappa/login"
			},
			Expiration:   time.Duration(config.HttpCacheTime) * time.Minute,
			CacheControl: true,
		}))
	}

	app.Use(favicon.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/", handlers.IndexHandler)
	app.Get("/hc", handlers.MappaHealthCheckHandler)

	mappa := app.Group("/mappa", middleware.NewMappaAuthMiddleware(middleware.MappaAuthMiddlewareConfig{}))
	mappa.Post("/login", handlers.MappaLoginHandler)
	mappa.Get("/escotista/:userId", handlers.MappaEscotistaHandler)
	mappa.Get("/escotista/:userId/secoes",handlers.MappaEscotistaSecoesHandler)
	mappa.Get("/progressoes/:ramo", handlers.MappaProgressoesHandler)
	mappa.Get("/*", handlers.MappaGenericHandler)

	if len(config.StaticFolder) > 0 {
		app.Static("/web", config.StaticFolder, fiber.Static{
			Compress:      true,
			ByteRange:     true,
			Browse:        true,
			Index:         "index.html",
			CacheDuration: time.Minute,
			MaxAge:        60,
		})
	}
	app.Post("/tg/pub", handlers.TelegramPublisherHandler)
	SetupSwagger(app)
	return app, nil
}
