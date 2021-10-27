package configuration

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Configuration struct {
	Host       string
	Port       int
	Repository string
	LogFolder  string
}

var executableName string

const (
	DEFAULT_MAPPA_PROXY_PORT = 8081
	DEFAULT_MAPPA_PROXY_HOST = "0.0.0.0"
)

func init() {
	en, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to identify current executable: %v", err)
	}
	executableName = en
}

func getEnv(env string, defaultValue string) string {
	envValue := os.Getenv(env)
	if len(envValue) > 0 {
		return envValue
	}
	return defaultValue
}
func GetConfiguration() (Configuration, error) {
	config := Configuration{}
	envPort := getEnv("MAPPA_PROXY_PORT", "8081")
	nPort, err := strconv.Atoi(envPort)
	if err != nil || nPort < 1024 {
		log.Printf("Invalid MAPPA_PROXY_PORT (%s) -> Using default (%d)", envPort, DEFAULT_MAPPA_PROXY_PORT)
	}
	envHost := getEnv("MAPPA_PROXY_HOST", "0.0.0.0")
	envLog := getEnv("MAPPA_LOG_FOLDER", filepath.Join(filepath.Dir(executableName), "log"))
	envRepository := getEnv("MAPPA_PROXY_REPOSITORY", "json:./mappa_login.json")

	var fPort = flag.Int("port", nPort, "HTTP port")
	var fHost = flag.String("host", envHost, "HTTP host")
	var fLog = flag.String("log", envLog, "Log folder")
	var fRepo = flag.String("repository", envRepository, "Repository")

	flag.Parse()

	if *fPort < 1024 {
		return config, fmt.Errorf("invalid port %d", *fPort)
	}
	config.Host = *fHost
	config.LogFolder = *fLog
	config.Repository = *fRepo
	config.Port = *fPort

	// else {
	// 	sPort := os.Getenv("HTTP_PORT")
	// 	if len(sPort) == 0 {
	// 		port = 8081
	// 		} else {
	// 			ePort, err := strconv.Atoi(sPort)
	// 			if err != nil {
	// 				log.Fatalf("Invalid HTTP_PORT environment:  %v", sPort)
	// 			}
	// 			port = ePort
	// 		}
	// 	}

	return config, err
}
