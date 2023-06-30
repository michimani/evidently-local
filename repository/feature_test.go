package repository_test

import (
	"os"
	"testing"

	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/repository"
	"github.com/stretchr/testify/assert"
)

func Test_NewFeatureRepositoryWithJSONFile(t *testing.T) {
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
