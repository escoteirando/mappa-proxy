package app

import (
	"fmt"
	"log"
	"strings"

	"net/http"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/app/handlers"
	"github.com/escoteirando/mappa-proxy/backend/app/middleware"
	"github.com/escoteirando/mappa-proxy/backend/app/scheduled"
	"github.com/escoteirando/mappa-proxy/backend/build"
	"github.com/escoteirando/mappa-proxy/backend/cache"
	"github.com/escoteirando/mappa-proxy/backend/configuration"
	"github.com/escoteirando/mappa-proxy/backend/infra/fiberfilestorage"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
	"github.com/escoteirando/mappa-proxy/backend/static"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	fiberCache "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cacheStorage fiber.Storage
)

func CreateServer(config configuration.Configuration, cache *cache.MappaCache, repository repositories.IRepository) (app *fiber.App, err error) {	
	app = fiber.New(fiber.Config{
		AppName:     fmt.Sprintf("%s - v%s @ %v", configuration.APP_NAME, configuration.APP_VERSION, build.BuildTime),
		ColorScheme: fiber.DefaultColors,
		// EnablePrintRoutes: true,
	})

	handlers.SetupUserContext(app, config, cache, repository)
	scheduled.Setup(handlers.GetCurrentUserContextData())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST,GET",
		AllowHeaders:     "Origin,Authorization,Content-Type,User-Agent",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	cacheStorage = fiberfilestorage.New(&fiberfilestorage.Config{
		BasePath:   config.CachePath,
		GCInterval: time.Duration(config.HttpCacheTime) * time.Minute,
	})
	RevalidateCacheAfterVersionChange(cacheStorage)

	if config.HttpCacheTime > 0 {
		app.Use(fiberCache.New(fiberCache.Config{
			Next: func(c *fiber.Ctx) bool {
				cPath := c.Route().Path
				return cPath == "/mappa/login" ||
					strings.HasPrefix(cPath, "/swagger") ||
					c.Response().StatusCode() > 299 ||
					strings.HasPrefix(cPath, "/admin")
			},
			//			Expiration: time.Duration(config.HttpCacheTime) * time.Minute,
			ExpirationGenerator: func(c *fiber.Ctx, cfg *fiberCache.Config) time.Duration {
				if r, ok := handlers.Routes[c.Route().Path]; ok {
					return r.CacheTime
				}

				return time.Duration(config.HttpCacheTime) * time.Minute
			},
			CacheControl:         true,
			StoreResponseHeaders: true,
			Storage:              cacheStorage,
		}))
	}

	app.Use(favicon.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	mappa := app.Group("/mappa", middleware.NewMappaAuthMiddleware(middleware.MappaAuthMiddlewareConfig{}))
	for routePath, route := range handlers.Routes {
		var router fiber.Router
		isMappa := ""
		if route.Mappa {
			router = mappa
			isMappa = "[Mappa] "
		} else {
			router = app
		}
		cacheStr := "NO"

		switch route.Method {
		case "POST":
			router.Post(routePath, route.Handler)
		default:
			router.Get(routePath, route.Handler)
		}
		if route.CacheTime > 0 {
			cacheStr = fmt.Sprintf("%v", route.CacheTime)
		}

		log.Printf("Route %s%s registered: %s %s (cache %s)", isMappa, route.Name, route.Method, routePath, cacheStr)
	}

	handlers.AddAdminHandlers(app, cacheStorage)

	app.Use("/web", filesystem.New(filesystem.Config{
		Root:       http.FS(static.EmbedStaticWeb),
		PathPrefix: "web",
		Browse:     true,
		MaxAge:     60,
	}))

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	SetupSwagger(app)

	scheduled.Run()
	return app, nil
}

func RevalidateCacheAfterVersionChange(storage fiber.Storage) {
	const versionKey = configuration.APP_NAME + "_version"
	if storage == nil {
		return
	}
	lastVersion, _ := storage.Get(versionKey)

	if string(lastVersion) != configuration.APP_VERSION {
		if err := storage.Reset(); err != nil {
			log.Printf("Error cleaning cache: %v", err)
		} else {
			storage.Set(versionKey, []byte(configuration.APP_VERSION), 0)
			log.Printf("Cache cleaned after version change: %v -> %v", string(lastVersion), configuration.APP_VERSION)
		}
	}
}
