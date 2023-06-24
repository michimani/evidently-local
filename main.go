package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	portEnvKey  = "EVIDENTLY_LOCAL_PORT"
	defaultPort = "2306"
)

func startServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, world!"))
	})

	fmt.Printf("Server started on localhost:%s\n", port)
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

	startServer(port)
}
