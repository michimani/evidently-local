package main

import (
	"os"

	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/repository"
	"github.com/michimani/evidentlylocal/server"
)

const (
	portEnvKey  = "EVIDENTLY_LOCAL_PORT"
	defaultPort = "2306"
	dataDir     = "./data"
)

func main() {
	port := os.Getenv(portEnvKey)
	if len(port) == 0 {
		port = defaultPort
	}

	l, err := logger.NewEvidentlyLocalLogger(os.Stdout)
	if err != nil {
		panic(err)
	}

	fRepo, err := repository.NewFeatureRepositoryWithJSONFile(dataDir, l)
	if err != nil {
		panic(err)
	}

	server.Start(port, l, fRepo)
}
