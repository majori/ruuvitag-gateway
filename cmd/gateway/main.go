package main

import (
	"io/ioutil"
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

	http.HandleFunc("/", handler(adapter))
	http.ListenAndServe(":8080", nil)
}

func handler(adapter ruuvi.DatabaseAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		b, _ := ioutil.ReadAll(req.Body)
		m, err := ruuvi.Parse(b)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%+v", m)

		if err := adapter.Save(m); err != nil {
			panic(err)
		}
	}
}
