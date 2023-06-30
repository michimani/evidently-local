package repository

import "github.com/michimani/evidentlylocal/logger"

func SetDataDirToFeatureRepositoryWithJSONFile(target *FeatureRepositoryWithJSONFile, dataDir string) {
	target.dataDir = dataDir
}

func SetLoggerToFeatureRepositoryWithJSONFile(target *FeatureRepositoryWithJSONFile, l logger.Logger) {
	target.l = l
}
