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

func Test_EvaluateFeature(t *testing.T) {
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
		{
			name:           "default rule",
			reqBody:        `{"entityId":"test-entity-id", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"details":"{}","reason":"DEFAULT (local)","value":{"boolValue":false},"variation":"False"}`,
		},
		{
			name:           "override rule",
			reqBody:        `{"entityId":"force-true", "evaluateContext":""}`,
			reqPath:        "/projects/test-project/evaluations/test-feature-1",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"details":"{}","reason":"OVERRIDE_RULE (local)","value":{"boolValue":true},"variation":"True"}`,
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
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			body := bytes.NewBufferString(c.reqBody)
			req := httptest.NewRequest(c.method, c.reqPath, body)

			w := httptest.NewRecorder()

			handler.Exported_evaluateFeature(w, req, testLogger)

			asst.Equal(c.expectedStatus, w.Code)
			asst.Equal(c.expectedBody, w.Body.String())
		})
	}
}