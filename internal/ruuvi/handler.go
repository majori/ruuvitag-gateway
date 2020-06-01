package ruuvi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Handler parses measurement from HTTP body and saves it to database
func Handler(adapter DatabaseAdapter) http.HandlerFunc {
	whitelist := make(map[string]struct{})
	dw := os.Getenv("DEVICE_ID_WHITELIST")
	if dw != "" {
		devices := strings.Split(dw, ",")
		fmt.Printf("Whitelisted device identifiers: %s\n", strings.Join(devices, ", "))
		for _, device := range devices {
			whitelist[device] = struct{}{}
		}
	}

	return func(w http.ResponseWriter, req *http.Request) {
		b, _ := ioutil.ReadAll(req.Body)

		m, err := Parse(b)
		if err != nil {
			fmt.Println("Request rejected: malformed payload")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(whitelist) > 0 {
			if _, ok := whitelist[m.DeviceID]; !ok {
				fmt.Printf("Request rejected: unauthorized device %s\n", m.DeviceID)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		fmt.Printf("Device %s: received %d tags\n", m.DeviceID, len(m.Tags))

		if err := adapter.Save(m); err != nil {
			panic(err)
		}
	}
}
