package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/michimani/evidentlylocal/logger"
)

const (
	portEnvKey  = "EVIDENTLY_LOCAL_PORT"
	defaultPort = "2306"
)

func startServer(port string, l logger.Logger) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, world!"))
	})

	l.Info(fmt.Sprintf("Server started on port %s", port))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv(portEnvKey)
	if len(port) == 0 {
		port = defaultPort
	}

	l, err := logger.NewEvidentlyLocalLogger(os.Stdout)
	if err != nil {
		panic(err)
	}

	startServer(port, l)
}
