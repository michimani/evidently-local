package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/michimani/evidentlylocal/components"
	"github.com/michimani/evidentlylocal/internal"
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

	bytes, requestID, err := internal.GenerateResponseBody(res)
	if err != nil {
		h.l.Error("Failed to generate response body", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-amzn-RequestId", requestID)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func (h *evaluationHandler) batchEvaluateFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.l.Error("Method not allowed: "+r.Method, nil)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 4 {
		h.l.Error("Invalid path: "+path, nil)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	project := parts[2]

	request := &types.BatchEvaluateFeatureRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		h.l.Error("Failed to decode request body", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	results := make([]types.EvaluationResult, len(request.Requests))

	wg := sync.WaitGroup{}
	for i, req := range request.Requests {
		wg.Add(1)

		go func(i int, req types.EvaluationRequest) {
			defer wg.Done()
			feature, err := repository.FeatureRepositoryInstance().Get(project, req.Feature)
			if err != nil {
				h.l.Error("Failed to get feature", err)
				results[i] = types.EvaluationResult{
					EntityID: req.EntityID,
					Feature:  req.Feature,
					Project:  project,
					Reason:   "Feature not found",
				}
				return
			}

			reason, variation, err := components.EvaluateFeature(feature, req.EntityID)
			if err != nil {
				h.l.Error("Failed to evaluate feature", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				results[i] = types.EvaluationResult{
					EntityID: req.EntityID,
					Feature:  req.Feature,
					Project:  project,
					Reason:   "Failed to evaluate feature",
				}
				return
			}

			res := types.EvaluationResult{
				Details:   "{}",
				EntityID:  req.EntityID,
				Feature:   req.Feature,
				Project:   project,
				Reason:    reason,
				Value:     variation.Value,
				Variation: variation.Name,
			}

			results[i] = res
		}(i, req)
	}

	wg.Wait()

	res := types.BatchEvaluateFeatureResponse{
		Results: results,
	}

	bytes, requestID, err := internal.GenerateResponseBody(res)
	if err != nil {
		h.l.Error("Failed to generate response body", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-amzn-RequestId", requestID)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
