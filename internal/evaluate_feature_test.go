package internal_test

import (
	"os"
	"testing"

	"github.com/michimani/evidentlylocal/internal"
	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/models"
	"github.com/michimani/evidentlylocal/repository"
	"github.com/michimani/evidentlylocal/types"
	"github.com/stretchr/testify/assert"
)

func Test_EvaluateFeature(t *testing.T) {
	l, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	repo, _ := repository.NewFeatureRepositoryWithJSONFile("../testdata", l)
	repository.SetFeatureRepositoryInstance(repo)

	cases := []struct {
		name            string
		project         string
		feature         string
		entityID        string
		wantErr         bool
		expectReason    string
		expectVariation *models.Variation
	}{
		{
			name:            "project not found",
			project:         "not-exists-project",
			feature:         "test-feature-1",
			entityID:        "test-entity-1",
			wantErr:         true,
			expectReason:    "",
			expectVariation: nil,
		},
		{
			name:            "feature not found",
			project:         "test-project",
			feature:         "not-exists-feature",
			entityID:        "test-entity-1",
			wantErr:         true,
			expectReason:    "",
			expectVariation: nil,
		},
		{
			name:         "entity override",
			project:      "test-project",
			feature:      "test-feature-1",
			entityID:     "force-true",
			wantErr:      false,
			expectReason: "OVERRIDE_RULE (local)",
			expectVariation: &models.Variation{
				Name: "True",
				Value: map[types.VariableValueType]any{
					types.VariableValueTypeBool: true,
				},
			},
		},
		{
			name:         "default variation",
			project:      "test-project",
			feature:      "test-feature-1",
			entityID:     "test-entity-1",
			wantErr:      false,
			expectReason: "DEFAULT (local)",
			expectVariation: &models.Variation{
				Name: "False",
				Value: map[types.VariableValueType]any{
					types.VariableValueTypeBool: false,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			reason, variation, err := internal.EvaluateFeature(c.project, c.feature, c.entityID)

			if c.wantErr {
				asst.Empty(reason, reason)
				asst.Nil(variation, variation)
				asst.Error(err)
				return
			}

			asst.NoError(err)
			asst.Equal(c.expectReason, reason, reason)
			asst.Equal(*c.expectVariation, *variation, variation)
		})
	}
}
