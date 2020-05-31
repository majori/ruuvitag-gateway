package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/majori/ruuvitag-gateway/internal/adapters/influxdb2"
	"github.com/majori/ruuvitag-gateway/internal/ruuvi"
)

func main() {
	url := os.Getenv("INFLUXDB_URL")
	if url == "" {
		panic(`"INFLUXDB_URL" missing`)
	}

	token := os.Getenv("INFLUXDB_TOKEN")
	if token == "" {
		panic(`"INFLUXDB_TOKEN" missing`)
	}

	adapter := influxdb2.New(url, token)

	port := 8080
	fmt.Printf("Server started on port %d\n", port)

	http.HandleFunc("/", ruuvi.Handler(adapter))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
