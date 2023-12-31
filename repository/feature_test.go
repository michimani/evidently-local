package repository_test

import (
	"os"
	"testing"

	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/models"
	"github.com/michimani/evidentlylocal/repository"
	"github.com/michimani/evidentlylocal/types"
	"github.com/stretchr/testify/assert"
)

func Test_NewFeatureRepositoryWithJSONFile(t *testing.T) {
	t.Parallel()

	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	testRepo := repository.FeatureRepositoryWithJSONFile{}
	repository.SetDataDirToFeatureRepositoryWithJSONFile(&testRepo, "testdata")
	repository.SetLoggerToFeatureRepositoryWithJSONFile(&testRepo, testLogger)

	cases := []struct {
		name    string
		dataDir string
		l       logger.Logger
		wantErr bool
		expect  *repository.FeatureRepositoryWithJSONFile
	}{
		{
			name:    "dataDir is empty",
			dataDir: "",
			l:       testLogger,
			wantErr: true,
			expect:  nil,
		},
		{
			name:    "logger is nil",
			dataDir: "testdata",
			l:       nil,
			wantErr: true,
			expect:  nil,
		},
		{
			name:    "success",
			dataDir: "testdata",
			l:       testLogger,
			wantErr: false,
			expect:  &testRepo,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			got, err := repository.NewFeatureRepositoryWithJSONFile(c.dataDir, c.l)
			if c.wantErr {
				asst.Nil(got)
				asst.Error(err)
				return
			}

			asst.NoError(err)
			asst.Equal(c.expect, got)
		})
	}
}

func Test_FeatureRepositoryWithJSONFile_Get(t *testing.T) {
	t.Parallel()

	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	testRepo, _ := repository.NewFeatureRepositoryWithJSONFile("../testdata", testLogger)

	cases := []struct {
		name        string
		repo        *repository.FeatureRepositoryWithJSONFile
		project     string
		featureName string
		wantErr     bool
		expect      *models.Feature
	}{
		{
			name:        "repo is nil",
			repo:        nil,
			project:     "test-project",
			featureName: "test-feature-1",
			wantErr:     true,
			expect:      nil,
		},
		{
			name:        "project not found",
			repo:        testRepo,
			project:     "not-exists-project",
			featureName: "test-feature-1",
			wantErr:     true,
			expect:      nil,
		},
		{
			name:        "feature not found",
			repo:        testRepo,
			project:     "test-project",
			featureName: "not-exists-feature",
			wantErr:     true,
			expect:      nil,
		},
		{
			name:        "success: bool value",
			repo:        testRepo,
			project:     "test-project",
			featureName: "test-feature-1",
			wantErr:     false,
			expect: &models.Feature{
				Name:             "test-feature-1",
				DefaultVariation: "False",
				EntityOverrides: models.EntityOverride{
					"force-true": "True",
				},
				Project:   "test-project",
				Status:    "AVAILABLE",
				ValueType: "BOOLEAN",
				Variations: []models.Variation{
					{
						Name: "True", Value: map[types.VariableValueType]any{
							types.VariableValueTypeBool: true,
						},
					},
					{
						Name: "False", Value: map[types.VariableValueType]any{
							types.VariableValueTypeBool: false,
						},
					},
				},
			},
		},
		{
			name:        "success: string value",
			repo:        testRepo,
			project:     "test-project",
			featureName: "test-feature-2",
			wantErr:     false,
			expect: &models.Feature{
				Name:             "test-feature-2",
				DefaultVariation: "String1",
				EntityOverrides: models.EntityOverride{
					"force-2": "String2",
				},
				Project:   "test-project",
				Status:    "AVAILABLE",
				ValueType: "STRING",
				Variations: []models.Variation{
					{
						Name: "String1", Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "string-1",
						},
					},
					{
						Name: "String2", Value: map[types.VariableValueType]any{
							types.VariableValueTypeString: "string-2",
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			got, err := c.repo.Get(c.project, c.featureName)
			if c.wantErr {
				asst.Nil(got)
				asst.Error(err)
				return
			}

			asst.NoError(err)
			asst.Equal(*c.expect, *got)
		})
	}
}

func Test_FeatureRepositoryWithJSONFile_List(t *testing.T) {
	t.Parallel()

	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	testRepo, _ := repository.NewFeatureRepositoryWithJSONFile("../testdata", testLogger)

	cases := []struct {
		name    string
		repo    *repository.FeatureRepositoryWithJSONFile
		project string
		wantErr bool
		expect  []*models.Feature
	}{
		{
			name:    "repo is nil",
			repo:    nil,
			project: "test-project",
			wantErr: true,
			expect:  nil,
		},
		{
			name:    "project not found",
			repo:    testRepo,
			project: "not-exists-project",
			wantErr: true,
			expect:  nil,
		},
		{
			name:    "has no features directory project",
			repo:    testRepo,
			project: "has-no-features-dir-project",
			wantErr: false,
			expect:  []*models.Feature{},
		},
		{
			name:    "has no feature project",
			repo:    testRepo,
			project: "has-no-feature-project",
			wantErr: false,
			expect:  []*models.Feature{},
		},
		{
			name:    "has invalid feature json project",
			repo:    testRepo,
			project: "has-invalid-json-project",
			wantErr: false,
			expect:  []*models.Feature{},
		},
		{
			name:    "success",
			repo:    testRepo,
			project: "test-project",
			wantErr: false,
			expect: []*models.Feature{
				{
					Name:             "test-feature-1",
					DefaultVariation: "False",
					EntityOverrides: models.EntityOverride{
						"force-true": "True",
					},
					Project:   "test-project",
					Status:    "AVAILABLE",
					ValueType: "BOOLEAN",
					Variations: []models.Variation{
						{
							Name: "True", Value: map[types.VariableValueType]any{
								types.VariableValueTypeBool: true,
							},
						},
						{
							Name: "False", Value: map[types.VariableValueType]any{
								types.VariableValueTypeBool: false,
							},
						},
					},
				},
				{
					Name:             "test-feature-2",
					DefaultVariation: "String1",
					EntityOverrides: models.EntityOverride{
						"force-2": "String2",
					},
					Project:   "test-project",
					Status:    "AVAILABLE",
					ValueType: "STRING",
					Variations: []models.Variation{
						{
							Name: "String1", Value: map[types.VariableValueType]any{
								types.VariableValueTypeString: "string-1",
							},
						},
						{
							Name: "String2", Value: map[types.VariableValueType]any{
								types.VariableValueTypeString: "string-2",
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			got, err := c.repo.List(c.project)
			if c.wantErr {
				asst.Nil(got)
				asst.Error(err)
				return
			}

			asst.NoError(err)
			asst.Equal(len(c.expect), len(got))
			for i, e := range c.expect {
				asst.Equal(*e, *got[i])
			}
		})
	}
}

func Test_SetAndGetFeatureRepositoryInstance(t *testing.T) {
	t.Parallel()

	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	testRepo, _ := repository.NewFeatureRepositoryWithJSONFile("../testdata", testLogger)

	asst := assert.New(t)

	asst.Nil(repository.FeatureRepositoryInstance())

	repository.SetFeatureRepositoryInstance(testRepo)
	instance := repository.FeatureRepositoryInstance()
	asst.NotNil(instance)
}
