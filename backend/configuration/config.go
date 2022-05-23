package configuration

import (
	"log"
	"os"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/infra"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Host         string `envconfig:"MAPPA_PROXY_HOST" default:"0.0.0.0"`
	Port         int    `envconfig:"MAPPA_PROXY_PORT" default:"5000"`
	Repository   string `envconfig:"MAPPA_PROXY_REPOSITORY" default:"json:./mappa_login.json"`
	LogFolder    string `envconfig:"MAPPA_LOG_FOLDER" default:"log"`
	StaticFolder string `envconfig:"MAPPA_STATIC_FOLDER" default:"web"`
}

const (
	DEFAULT_MAPPA_PROXY_PORT = "5000"
	DEFAULT_MAPPA_PROXY_HOST = "0.0.0.0"
	APP_NAME                 = "mappa-proxy"
	APP_VERSION              = "0.1.0"
)

var (
	Config      *Configuration
	StartupTime time.Time
)

func getEnv(env string, defaultValue string, validate func(value interface{}) error) {
	envValue := os.Getenv(env)
	if len(envValue) == 0 {
		envValue = defaultValue
	}
	if err := validate(envValue); err != nil {
		log.Fatalf("INVALID VALUE FOR ENV %s = %v %v", env, envValue, err)
	}
}

func Init() {
	Config = &Configuration{}
	err := envconfig.Process("", Config)
	if err != nil {
		log.Fatalf("Failed to process env variables: %v", err)
	}
	if len(Config.Host) == 0 {
		log.Fatalf("MISSING ENV MAPPA_PROXY_HOST")
	}
	if Config.Port < 80 || Config.Port > 65535 {
		log.Fatalf("MAPPA_PROXY_PORT MUST BE BETWEEN 80 AND 65535: %d", Config.Port)
	}

	if stat, err := os.Stat(Config.LogFolder); err != nil || !stat.IsDir() {
		log.Printf("MAPPA_LOG_FOLDER path not found %s - WILL NOT SAVE LOGS", Config.LogFolder)
		Config.LogFolder = ""
	}

	cs, err := infra.CreateConnectionString(Config.Repository)
	if err != nil {
		log.Fatalf("INVALID MAPPA_PROXY_REPOSITORY %s - %v", Config.Repository, err)
	}
	Config.Repository = cs.String()

	if stat, err := os.Stat(Config.StaticFolder); err != nil || !stat.IsDir() {
		log.Printf("MAPPA_STATIC_FOLDER path not found %s - WILL NOT USE STATIC FILES", Config.StaticFolder)
		Config.StaticFolder = ""
	}

	StartupTime = time.Now()
}
