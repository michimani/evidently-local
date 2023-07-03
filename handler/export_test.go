package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/repository"
)

const (
	dataDir = "../testdata"
)

func PrepareForTest(l logger.Logger) {
	repos, _ := repository.NewFeatureRepositoryWithJSONFile(dataDir, l)
	repository.SetFeatureRepositoryInstance(repos)
}

func Exported_handleSomeResources(w http.ResponseWriter, r *http.Request) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	ph := NewProjectHandler(testLogger)
	path := r.URL.Path
	parts := strings.Split(path, "/")
	ph.pathParts = parts
	ph.handleSomeResources(w, r)
}

func Exported_handleSpecificResource(w http.ResponseWriter, r *http.Request) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	ph := NewProjectHandler(testLogger)
	path := r.URL.Path
	parts := strings.Split(path, "/")
	ph.pathParts = parts
	ph.handleSpecificResource(w, r)
}

func Exported_evaluateFeature(w http.ResponseWriter, r *http.Request) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	eh := newEvaluationHandler(testLogger)
	eh.evaluateFeature(w, r)
}
