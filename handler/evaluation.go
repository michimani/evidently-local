package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/michimani/evidentlylocal/components"
	"github.com/michimani/evidentlylocal/logger"
	"github.com/michimani/evidentlylocal/repository"
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

	feature, err := repository.FeatureRepositoryInstance().Get(project, featureName)
	if err != nil {
		l.Error("Failed to get feature", err)
		http.Error(w, "Resource Not Found", http.StatusNotFound)
		return
	}

	// TODO: get entityID from request body
	request := &types.EvaluateFeatureRequest{}
	err = json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		l.Error("Failed to decode request body", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	entityID := request.EntityID

	reason, variation, err := components.EvaluateFeature(feature, entityID)
	if err != nil {
		l.Error("Failed to evaluate feature", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	l.Info(fmt.Sprintf("return variation: %+v", variation))

	res := types.EvaluateFeatureResponse{
		Details:   "{}",
		Reason:    reason,
		Value:     variation.Value,
		Variation: variation.Name,
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
