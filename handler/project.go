package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/michimani/evidentlylocal/logger"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	l, err := logger.NewEvidentlyLocalLogger(os.Stdout)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")
	l.Info(fmt.Sprintf("%s %s", r.Method, path))

	switch len(parts) {
	case 2:
		// GET | POST /projects
		handleProjects(w, r, l)
	case 3:
		// GET | PATCH | DELETE /projects/:project
		handleSpecificProject(w, r, parts[2], l)
	case 4:
		// POST /projects/:project/evaluations
		// GET | POST /projects/:project/experiments
		// GET | POST /projects/:project/launches
		// GET | POST /projects/:project/features
		handleSomeResources(w, r, parts, l)
	case 5:
		// POST /projects/:project/evaluations/:feature
		// GET | PATCH | DELETE /projects/:project/experiments/:experiment
		// GET | PATCH | DELETE /projects/:project/launches/:launch
		// GET | PATCH | DELETE /projects/:project/features/:feature
		handleSpecificResource(w, r, parts, l)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func handleProjects(w http.ResponseWriter, r *http.Request, l logger.Logger) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleSpecificProject(w http.ResponseWriter, r *http.Request, project string, l logger.Logger) {
	switch r.Method {
	case http.MethodGet, http.MethodPatch, http.MethodDelete:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleSomeResources(w http.ResponseWriter, r *http.Request, pathPart []string, l logger.Logger) {
	if len(pathPart) != 4 {
		l.Error("Invalid path: "+r.URL.Path, nil)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	switch pathPart[3] {
	case "evaluations", "experiments", "launches", "features":
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func handleSpecificResource(w http.ResponseWriter, r *http.Request, pathPart []string, l logger.Logger) {
	if len(pathPart) != 5 {
		l.Error("Invalid path: "+r.URL.Path, nil)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	switch pathPart[3] {
	case "evaluations":
		// POST /projects/:project/evaluations/:feature
		// https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_EvaluateFeature.html
		evaluateFeature(w, r, l)
	case "experiments", "launches", "features":
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
