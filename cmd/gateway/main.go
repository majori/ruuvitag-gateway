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

	org := os.Getenv("INFLUXDB_ORGANIZATION")
	if org == "" {
		fmt.Println("Warning: INFLUXDB_ORGANIZATION environment variable is missing")
	}

	bucket := os.Getenv("INFLUXDB_BUCKET")
	if bucket == "" {
		fmt.Println("Warning: INFLUXDB_BUCKET environment variable is missing")
	}

	adapter := influxdb2.New(url, token, org, bucket)

	port := 8080
	fmt.Printf("Server started on port %d\n", port)

	http.HandleFunc("/", ruuvi.Handler(adapter))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
