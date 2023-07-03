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

type evaluationHandler struct {
	l logger.Logger
}

func newEvaluationHandler(l logger.Logger) *evaluationHandler {
	return &evaluationHandler{
		l: l,
	}
}

func (h *evaluationHandler) evaluateFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.l.Error("Method not allowed: "+r.Method, nil)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 5 {
		h.l.Error("Invalid path: "+path, nil)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	project := parts[2]
	featureName := parts[4]

	feature, err := repository.FeatureRepositoryInstance().Get(project, featureName)
	if err != nil {
		h.l.Error("Failed to get feature", err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	request := &types.EvaluateFeatureRequest{}
	err = json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		h.l.Error("Failed to decode request body", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	entityID := request.EntityID

	reason, variation, err := components.EvaluateFeature(feature, entityID)
	if err != nil {
		h.l.Error("Failed to evaluate feature", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.l.Info(fmt.Sprintf("return variation: %+v", variation))

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
		h.l.Error("Failed to generate UUID. Use constant request id.", err)
		requestID = "xxxxxxxx-0000-0000-0000-xxxxxxxxxxxx"
	} else {
		requestID = uuid.String()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-amzn-RequestId", requestID)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
