package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/types"
)

func evaluateFeature(w http.ResponseWriter, r *http.Request, l logger.Logger) {
	if r.Method != http.MethodPost {
		l.Error("Method not allowed: "+r.Method, nil)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 5 {
		l.Error("Invalid path: "+path, nil)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	project := parts[2]
	feature := parts[4]
	l.Info("Project: " + project)
	l.Info("Feature: " + feature)

	res := types.EvaluateFeatureResponse{
		Details: "{}",
		Reason:  "DEFAULT-local",
		Value: types.VariableValue{
			types.VariableValueTypeBool: false,
		},
		Variation: "False",
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	requestID := ""
	uuid, err := uuid.NewV4()
	if err != nil {
		l.Error("Failed to generate UUID. Use constant request id.", err)
		requestID = "xxxxxxxx-0000-0000-0000-xxxxxxxxxxxx"
	} else {
		requestID = uuid.String()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-amzn-RequestId", requestID)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
