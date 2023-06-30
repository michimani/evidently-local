package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/models"
)

var featureRepositoryInstance FeatureRepository

func SetFeatureRepositoryInstance(r FeatureRepository) {
	featureRepositoryInstance = r
}

func FeatureRepositoryInstance() FeatureRepository {
	return featureRepositoryInstance
}

type FeatureRepository interface {
	Get(project, feature string) (*models.Feature, error)
	List(project string) ([]*models.Feature, error)
}

var _ FeatureRepository = (*FeatureRepositoryWithJSONFile)(nil)

type FeatureRepositoryWithJSONFile struct {
	dataDir string
	l       logger.Logger
}

func NewFeatureRepositoryWithJSONFile(dataDir string, l logger.Logger) (*FeatureRepositoryWithJSONFile, error) {
	if len(dataDir) == 0 {
		return nil, errors.New("dataDir is empty")
	}

	if l == nil {
		return nil, errors.New("logger is nil")
	}

	return &FeatureRepositoryWithJSONFile{
		dataDir: dataDir,
		l:       l,
	}, nil
}

func (r *FeatureRepositoryWithJSONFile) Get(project, featureName string) (*models.Feature, error) {
	if r == nil {
		return nil, errors.New("FeatureRepositoryWithJSONFile is nil")
	}

	projectDir := filepath.Join(r.dataDir, "projects", project)
	if _, err := os.Stat(projectDir); err != nil {
		r.l.Error("project directory not found", err)
		return nil, fmt.Errorf("Project not found: %s", project)
	}

	featureFile := path.Join(projectDir, "features", featureName+".json")
	if _, err := os.Stat(featureFile); err != nil {
		r.l.Error("feature file not found", err)
		return nil, fmt.Errorf("Feature not found: %s", featureName)
	}

	feature, err := r.getFeatureByFilePath(featureFile)
	if err != nil {
		return nil, err
	}

	return feature, nil
}

func (r *FeatureRepositoryWithJSONFile) List(project string) ([]*models.Feature, error) {
	if r == nil {
		return nil, errors.New("FeatureRepositoryWithJSONFile is nil")
	}

	projectDir := filepath.Join(r.dataDir, "projects", project)
	if _, err := os.Stat(projectDir); err != nil {
		r.l.Error("project directory not found", err)
		return nil, fmt.Errorf("Project not found: %s", project)
	}

	files, err := os.ReadDir(filepath.Join(projectDir, "features"))
	if err != nil {
		r.l.Error("failed to read project directory", err)
		return nil, err
	}

	res := []*models.Feature{}
	for _, file := range files {
		if !file.IsDir() && file.Name()[len(file.Name())-5:] == ".json" {
			featureFilePath := filepath.Join(projectDir, "features", file.Name())
			feature, err := r.getFeatureByFilePath(featureFilePath)
			if err != nil {
				r.l.Error("failed to get feature", err)
				continue
			}

			res = append(res, feature)
		}
	}

	return res, nil
}

func (r *FeatureRepositoryWithJSONFile) getFeatureByFilePath(path string) (*models.Feature, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		r.l.Error("failed to read feature file", err)
		return nil, err
	}

	feature := &models.Feature{}

	if err = json.Unmarshal([]byte(f), feature); err != nil {
		r.l.Error("failed to unmarshal feature file", err)
		return nil, err
	}

	return feature, nil
}
