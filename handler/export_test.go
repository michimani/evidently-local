package handler

import (
	"net/http"
	"strings"

	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/repository"
)

const (
	dataDir = "../testdata"
)

var testLogger logger.Logger

func PrepareForTest(l logger.Logger) {
	testLogger = l
	repos, _ := repository.NewFeatureRepositoryWithJSONFile(dataDir, l)
	repository.SetFeatureRepositoryInstance(repos)
}

func Exported_handleSomeResources(w http.ResponseWriter, r *http.Request) {
	ph := NewProjectHandler(testLogger)
	path := r.URL.Path
	parts := strings.Split(path, "/")
	ph.pathParts = parts
	ph.handleSomeResources(w, r)
}

func Exported_handleSpecificResource(w http.ResponseWriter, r *http.Request) {
	ph := NewProjectHandler(testLogger)
	path := r.URL.Path
	parts := strings.Split(path, "/")
	ph.pathParts = parts
	ph.handleSpecificResource(w, r)
}

func Exported_evaluateFeature(w http.ResponseWriter, r *http.Request) {
	eh := newEvaluationHandler(testLogger)
	eh.evaluateFeature(w, r)
}

func Exported_batchEvaluateFeature(w http.ResponseWriter, r *http.Request) {
	eh := newEvaluationHandler(testLogger)
	eh.batchEvaluateFeature(w, r)
}
