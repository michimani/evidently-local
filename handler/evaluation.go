package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/michimani/evidentlylocal/internal"
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
	featureName := parts[4]

	request := &types.EvaluateFeatureRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		l.Error("Failed to decode request body", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	entityID := request.EntityID

	reason, variation, eerr := internal.EvaluateFeature(project, featureName, entityID)
	if eerr != nil {
		l.Error("Failed to evaluate feature", eerr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := types.EvaluateFeatureResponse{
		Details:   "{}",
		Reason:    reason,
		Value:     variation.Value,
		Variation: variation.Name,
	}

	bytes, requestID, gerr := internal.GenerateResponseBody(res)
	if gerr != nil {
		l.Error("Failed to generate response body", gerr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-amzn-RequestId", requestID)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
