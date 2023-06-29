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

	switch r.Method {
	case http.MethodGet:
		get(w, r, l)
	case http.MethodPost:
		post(w, r, l)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func get(w http.ResponseWriter, r *http.Request, l logger.Logger) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	l.Info(fmt.Sprintf("GET %s", path))

	switch len(parts) {
	case 2:
		// GET /projects
		// return all projects, but not supported yet
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	case 3:
		// GET /projects/:project
		// return the project, but not supported yet
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	case 4:
		// GET /projects/:project/evaluations
		// return all evaluations of the project, but not supported yet
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	case 5:
		// GET /projects/:project/evaluations/:feature
		// return the evaluation of the feature, but not supported yet
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func post(w http.ResponseWriter, r *http.Request, l logger.Logger) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	l.Info(fmt.Sprintf("POST %s", path))

	switch len(parts) {
	case 2:
		// POST /projects
		// create a project, but not supported yet
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	case 3:
		// method not allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case 4:
		// POST /projects/:project/evaluations
		// create an evaluation, but not supported yet
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	case 5:
		// POST /projects/:project/evaluations/:feature
		// evaluate feature
		evaluateFeature(w, r, l)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
