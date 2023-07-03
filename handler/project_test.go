package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/michimani/evidentlylocal/handler"
	"github.com/michimani/evidentlylocal/logger"
	"github.com/stretchr/testify/assert"
)

func Test_Project(t *testing.T) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	handler.PrepareForTest(testLogger)
	ph := handler.NewProjectHandler(testLogger)

	cases := []struct {
		name           string
		reqBody        string
		reqPath        string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		// invalid path
		{
			name:           "GET /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "POST /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "DELETE /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PUT /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PATCH /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "HEAD /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		// /projects
		{
			name:           "GET /projects",
			reqPath:        "/projects",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects",
			reqPath:        "/projects",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects",
			reqPath:        "/projects",
			method:         http.MethodDelete,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "PUT /projects",
			reqPath:        "/projects",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "PATCH /projects",
			reqPath:        "/projects",
			method:         http.MethodPatch,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "HEAD /projects",
			reqPath:        "/projects",
			method:         http.MethodHead,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		// /projects/:project
		{
			name:           "GET /projects/:project",
			reqPath:        "/projects/test-project",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project",
			reqPath:        "/projects/test-project",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "DELETE /projects/:project",
			reqPath:        "/projects/test-project",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project",
			reqPath:        "/projects/test-project",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "PATCH /projects/:project",
			reqPath:        "/projects/test-project",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project",
			reqPath:        "/projects/test-project",
			method:         http.MethodHead,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		// /projects/:project/evaluations
		{
			name:           "GET /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/experiments
		{
			name:           "GET /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/launches
		{
			name:           "GET /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/features
		{
			name:           "GET /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/invalid-resource
		{
			name:           "GET /projects/:project/invalid-resource",
			reqPath:        "/projects/test-project/invalid-resource",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "POST /projects/:project/invalid-resource",
			reqPath:        "/projects/test-project/invalid-resource",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "DELETE /projects/:project/invalid-resource",
			reqPath:        "/projects/test-project/invalid-resource",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PUT /projects/:project/invalid-resource",
			reqPath:        "/projects/test-project/invalid-resource",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PATCH /projects/:project/invalid-resource",
			reqPath:        "/projects/test-project/invalid-resource",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "HEAD /projects/:project/invalid-resource",
			reqPath:        "/projects/test-project/invalid-resource",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		// /projects/:project/evaluations/:feature
		{
			name:           "GET /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "POST /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Bad request\n",
		},
		{
			name:           "DELETE /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "PUT /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "PATCH /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "HEAD /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		// /projects/:project/experiments/:experiment
		{
			name:           "GET /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/launches/:launch
		{
			name:           "GET /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/features/:feature
		{
			name:           "GET /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/invalid-resource/:feature
		{
			name:           "GET /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "POST /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "DELETE /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PUT /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PATCH /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "HEAD /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			body := bytes.NewBufferString(c.reqBody)
			req := httptest.NewRequest(c.method, c.reqPath, body)

			w := httptest.NewRecorder()

			ph.Projects(w, req)

			asst.Equal(c.expectedStatus, w.Code)
			asst.Equal(c.expectedBody, w.Body.String())
		})
	}
}

func Test_handleSomeResources(t *testing.T) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	handler.PrepareForTest(testLogger)

	cases := []struct {
		name           string
		reqBody        string
		reqPath        string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		// invalid path
		{
			name:           "GET /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "POST /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "DELETE /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PUT /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PATCH /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "HEAD /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		// /projects/:project/evaluations
		{
			name:           "GET /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/evaluations",
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/experiments
		{
			name:           "GET /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/experiments",
			reqPath:        "/projects/test-project/experiments",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/launches
		{
			name:           "GET /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/launches",
			reqPath:        "/projects/test-project/launches",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/features
		{
			name:           "GET /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/features",
			reqPath:        "/projects/test-project/features",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			body := bytes.NewBufferString(c.reqBody)
			req := httptest.NewRequest(c.method, c.reqPath, body)

			w := httptest.NewRecorder()

			handler.Exported_handleSomeResources(w, req)

			asst.Equal(c.expectedStatus, w.Code)
			asst.Equal(c.expectedBody, w.Body.String())
		})
	}
}

func Test_handleSpecificResource(t *testing.T) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(os.Stdout)
	handler.PrepareForTest(testLogger)

	cases := []struct {
		name           string
		reqBody        string
		reqPath        string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		// invalid path
		{
			name:           "GET /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "POST /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "DELETE /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PUT /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PATCH /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "HEAD /projects/invalid/path/invalid/path/invalid/path",
			reqPath:        "/projects/invalid/path/invalid/path/invalid/path",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		// /projects/:project/evaluations/:feature
		{
			name:           "GET /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "POST /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Bad request\n",
		},
		{
			name:           "DELETE /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "PUT /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "PATCH /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "HEAD /projects/:project/evaluations/:feature",
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		// /projects/:project/experiments/:experiment
		{
			name:           "GET /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/experiments/:experiment",
			reqPath:        "/projects/test-project/experiments/test-experiment-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/launches/:launch
		{
			name:           "GET /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/launches/:launch",
			reqPath:        "/projects/test-project/launches/test-launch-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/features/:feature
		{
			name:           "GET /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "POST /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "DELETE /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PUT /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "PATCH /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		{
			name:           "HEAD /projects/:project/features/:feature",
			reqPath:        "/projects/test-project/features/test-feature-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotImplemented,
			expectedBody:   "Not implemented\n",
		},
		// /projects/:project/invalid-resource/:feature
		{
			name:           "GET /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "POST /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "DELETE /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PUT /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "PATCH /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "HEAD /projects/:project/invalid-resource/:feature",
			reqPath:        "/projects/test-project/invalid-resource/test-feature-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			body := bytes.NewBufferString(c.reqBody)
			req := httptest.NewRequest(c.method, c.reqPath, body)

			w := httptest.NewRecorder()

			handler.Exported_handleSpecificResource(w, req)

			asst.Equal(c.expectedStatus, w.Code)
			asst.Equal(c.expectedBody, w.Body.String())
		})
	}
}
