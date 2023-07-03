package handler

import (
	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/repository"
)

var (
	Exported_evaluateFeature        = evaluateFeature
	Exported_handleSomeResources    = handleSomeResources
	Exported_handleSpecificResource = handleSpecificResource
)

const (
	dataDir = "../testdata"
)

func PrepareForTest(l logger.Logger) {
	repos, _ := repository.NewFeatureRepositoryWithJSONFile(dataDir, l)
	repository.SetFeatureRepositoryInstance(repos)
}
