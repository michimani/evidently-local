package handler_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/michimani/evidentlylocal/handler"
	"github.com/michimani/evidentlylocal/logger"
	"github.com/stretchr/testify/assert"
)

func Test_evaluateFeature(t *testing.T) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(io.Discard)
	handler.PrepareForTest(testLogger)

	cases := []struct {
		name           string
		reqBody        string
		reqPath        string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "default rule",
			reqBody:        `{"entityId":"test-entity-id", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"details":"{}","reason":"DEFAULT","value":{"boolValue":false},"variation":"False"}`,
		},
		{
			name:           "override rule",
			reqBody:        `{"entityId":"force-true", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"details":"{}","reason":"OVERRIDE_RULE","value":{"boolValue":true},"variation":"True"}`,
		},
		{
			name:           "feature not found",
			reqBody:        `{"entityId":"test-entity-id", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/not-exists-feature",
			method:         http.MethodPost,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found\n",
		},
		{
			name:           "invalid request body",
			reqBody:        `///`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Bad request\n",
		},
		{
			name:           "method not allowed: GET",
			reqBody:        `{"entityId":"test-entity-id", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "method not allowed: PUT",
			reqBody:        `{"entityId":"test-entity-id", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "method not allowed: PATCH",
			reqBody:        `{"entityId":"test-entity-id", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPatch,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "method not allowed: HEAD",
			reqBody:        `{"entityId":"test-entity-id", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodHead,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "invalid request path",
			reqBody:        `///`,
			reqPath:        "/projects/test-project",
			method:         http.MethodPost,
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

			handler.Exported_evaluateFeature(w, req)

			asst.Equal(c.expectedStatus, w.Code)
			asst.Equal(c.expectedBody, w.Body.String())
		})
	}
}

func Test_batchEvaluateFeature(t *testing.T) {
	testLogger, _ := logger.NewEvidentlyLocalLogger(io.Discard)
	handler.PrepareForTest(testLogger)

	cases := []struct {
		name           string
		reqBody        string
		reqPath        string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "one request",
			reqBody:        `{"requests":[{"entityId":"test-entity-id", "feature": "test-feature-1", "evaluateContext":""}]}`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"results":[{"details":"{}","entityId":"test-entity-id","feature":"test-feature-1","project":"test-project","reason":"DEFAULT","variation":"False","value":{"boolValue":false}}]}`,
		},
		{
			name:           "two requests (default and override)",
			reqBody:        `{"requests":[{"entityId":"test-entity-id", "feature": "test-feature-1", "evaluateContext":""},{"entityId":"force-true", "feature": "test-feature-1", "evaluateContext":""}]}`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"results":[{"details":"{}","entityId":"test-entity-id","feature":"test-feature-1","project":"test-project","reason":"DEFAULT","variation":"False","value":{"boolValue":false}},{"details":"{}","entityId":"force-true","feature":"test-feature-1","project":"test-project","reason":"OVERRIDE_RULE","variation":"True","value":{"boolValue":true}}]}`,
		},
		{
			name:           "with feature not found",
			reqBody:        `{"requests":[{"entityId":"test-entity-id", "feature": "test-feature-1", "evaluateContext":""},{"entityId":"test-entity-id", "feature": "not-exists-feature", "evaluateContext":""}]}`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"results":[{"details":"{}","entityId":"test-entity-id","feature":"test-feature-1","project":"test-project","reason":"DEFAULT","variation":"False","value":{"boolValue":false}},{"details":"","entityId":"test-entity-id","feature":"not-exists-feature","project":"test-project","reason":"Feature not found","variation":"","value":null}]}`,
		},
		{
			name:           "invalid request body",
			reqBody:        `///`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Bad request\n",
		},
		{
			name:           "method not allowed: GET",
			reqBody:        `{"requests":[{"entityId":"test-entity-id", "feature": "test-feature-1", "evaluateContext":""}]}`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "method not allowed: PUT",
			reqBody:        `{"requests":[{"entityId":"test-entity-id", "feature": "test-feature-1", "evaluateContext":""}]}`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "method not allowed: PATCH",
			reqBody:        `{"requests":[{"entityId":"test-entity-id", "feature": "test-feature-1", "evaluateContext":""}]}`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodPatch,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "method not allowed: HEAD",
			reqBody:        `{"requests":[{"entityId":"test-entity-id", "feature": "test-feature-1", "evaluateContext":""}]}`,
			reqPath:        "/projects/test-project/evaluations",
			method:         http.MethodHead,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "invalid request path",
			reqBody:        `///`,
			reqPath:        "/projects/test-project",
			method:         http.MethodPost,
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

			handler.Exported_batchEvaluateFeature(w, req)

			asst.Equal(c.expectedStatus, w.Code)
			asst.Equal(c.expectedBody, w.Body.String())
		})
	}
}
