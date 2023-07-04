package components

import (
	"github.com/michimani/evidentlylocal/models"
	"github.com/michimani/evidentlylocal/types"
)

func EvaluateFeature(feature *models.Feature, entityID string) (types.EvaluationReason, models.Variation, error) {
	// check override rules
	for overrideEntityID, overrideVariationName := range feature.EntityOverrides {
		if overrideEntityID == entityID {
			value := feature.GetValue(overrideVariationName)

			v := models.Variation{
				Name: overrideVariationName,
				Value: map[types.VariableValueType]any{
					feature.VariableValueType(): value,
				},
			}
			return types.EvaluationReasonOverride, v, nil
		}
	}

	// TODO: check percentage rules

	// return default variation
	v := models.Variation{
		Name: feature.DefaultVariation,
		Value: map[types.VariableValueType]any{
			feature.VariableValueType(): feature.GetDefaultValue(),
		},
	}

	return types.EvaluationReasonDefault, v, nil
}
