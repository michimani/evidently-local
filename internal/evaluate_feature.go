package internal

import (
	"github.com/michimani/evidentlylocal/components"
	"github.com/michimani/evidentlylocal/models"
	"github.com/michimani/evidentlylocal/repository"
)

func EvaluateFeature(project, featureName, entityID string) (string, *models.Variation, error) {
	feature, err := repository.FeatureRepositoryInstance().Get(project, featureName)
	if err != nil {
		return "", nil, err
	}

	return components.EvaluateFeature(feature, entityID)
}
