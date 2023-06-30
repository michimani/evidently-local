package models_test

import (
	"testing"

	"github.com/michimani/evidentlylocal/models"
	"github.com/michimani/evidentlylocal/types"
	"github.com/stretchr/testify/assert"
)

func Test_Feature_GetValue(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		feature   *models.Feature
		variation string
		expect    any
	}{
		{
			name: "string value",
			feature: &models.Feature{
				ValueType: types.FeatureValueTypeString,
				Variations: []models.Variation{
					{
						Name: "string-value-1",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "string-value-1",
						},
					},
					{
						Name: "string-value-2",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "string-value-2",
						},
					},
				},
			},
			variation: "string-value-1",
			expect:    "string-value-1",
		},
		{
			name: "bool value",
			feature: &models.Feature{
				ValueType: types.FeatureValueTypeBoolean,
				Variations: []models.Variation{
					{
						Name: "bool-value-true",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeBool: true,
						},
					},
					{
						Name: "bool-value-false",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeBool: false,
						},
					},
				},
			},
			variation: "bool-value-false",
			expect:    false,
		},
		{
			name: "long value",
			feature: &models.Feature{
				ValueType: types.FeatureValueTypeLong,
				Variations: []models.Variation{
					{
						Name: "long-value-1",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeLong: 1.0,
						},
					},
					{
						Name: "long-value-2",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeLong: 2.0,
						},
					},
				},
			},
			variation: "long-value-2",
			expect:    2.0,
		},
		{
			name: "double value",
			feature: &models.Feature{
				ValueType: types.FeatureValueTypeDouble,
				Variations: []models.Variation{
					{
						Name: "double-value-1",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeDouble: 1.0,
						},
					},
					{
						Name: "double-value-2",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeDouble: 2.0,
						},
					},
				},
			},
			variation: "double-value-2",
			expect:    2.0,
		},
		{
			name: "not exists value",
			feature: &models.Feature{
				ValueType: types.FeatureValueTypeString,
				Variations: []models.Variation{
					{
						Name: "double-value-1",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "value-1",
						},
					},
					{
						Name: "double-value-2",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "value-2",
						},
					},
				},
			},
			variation: "not-exists-variation",
			expect:    nil,
		},
		{
			name: "undefined value type",
			feature: &models.Feature{
				ValueType: types.FeatureValueType("undefined"),
				Variations: []models.Variation{
					{
						Name: "string-value-1",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "string-value-1",
						},
					},
					{
						Name: "string-value-2",
						Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "string-value-2",
						},
					},
				},
			},
			variation: "string-value-1",
			expect:    nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			value := c.feature.GetValue(c.variation)
			asst.Equal(c.expect, value, value)
		})
	}
}
