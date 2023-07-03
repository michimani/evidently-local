package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/michimani/evidentlylocal/logger"
)

type ProjectHandler struct {
	l         logger.Logger
	pathParts []string
}

func NewProjectHandler(l logger.Logger) *ProjectHandler {
	return &ProjectHandler{
		l: l,
	}
}

func (h *ProjectHandler) Projects(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	h.pathParts = parts
	h.l.Info(fmt.Sprintf("%s %s", r.Method, path))

	switch len(parts) {
	case 2:
		// GET | POST /projects
		h.handleProjects(w, r)
	case 3:
		// GET | PATCH | DELETE /projects/:project
		h.handleSpecificProject(w, r)
	case 4:
		// POST /projects/:project/evaluations
		// GET | POST /projects/:project/experiments
		// GET | POST /projects/:project/launches
		// GET | POST /projects/:project/features
		h.handleSomeResources(w, r)
	case 5:
		// POST /projects/:project/evaluations/:feature
		// GET | PATCH | DELETE /projects/:project/experiments/:experiment
		// GET | PATCH | DELETE /projects/:project/launches/:launch
		// GET | PATCH | DELETE /projects/:project/features/:feature
		h.handleSpecificResource(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func (h *ProjectHandler) handleProjects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProjectHandler) handleSpecificProject(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPatch, http.MethodDelete:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProjectHandler) handleSomeResources(w http.ResponseWriter, r *http.Request) {
	if len(h.pathParts) != 4 {
		h.l.Error("Invalid path: "+r.URL.Path, nil)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	switch h.pathParts[3] {
	case "evaluations", "experiments", "launches", "features":
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func (h *ProjectHandler) handleSpecificResource(w http.ResponseWriter, r *http.Request) {
	if len(h.pathParts) != 5 {
		h.l.Error("Invalid path: "+r.URL.Path, nil)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	switch h.pathParts[3] {
	case "evaluations":
		// POST /projects/:project/evaluations/:feature
		// https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_EvaluateFeature.html
		eh := newEvaluationHandler(h.l)
		eh.evaluateFeature(w, r)
	case "experiments", "launches", "features":
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
