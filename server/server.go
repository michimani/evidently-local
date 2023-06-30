package server

import (
	"fmt"
	"net/http"

	"github.com/michimani/evidentlylocal/handler"
	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/repository"
)

func Start(port string, l logger.Logger, repoInstance repository.FeatureRepository) {
	repository.SetFeatureRepositoryInstance(repoInstance)

	http.HandleFunc("/projects/", handler.Projects)

	l.Info(fmt.Sprintf("Server started on port %s", port))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
