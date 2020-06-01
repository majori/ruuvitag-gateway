package ruuvi

import (
	"encoding/json"
)

type Measurement struct {
	Tags []struct {
		AccelX                    float32 `json:"accelX"`
		AccelY                    float32 `json:"accelY"`
		AccelZ                    float32 `json:"accelZ"`
		DataFormat                int     `json:"dataFormat"`
		DefaultBackground         int     `json:"defaultBackground"`
		Favorite                  bool    `json:"favorite"`
		Humidity                  float32 `json:"humidity"`
		ID                        string  `json:"id"`
		MeasurementSequenceNumber int     `json:"measurementSequenceNumber"`
		MovementCounter           int     `json:"movementCounter"`
		Name                      string  `json:"name,omitempty"`
		Pressure                  float32 `json:"pressure"`
		RSSI                      int     `json:"rssi"`
		Temperature               float32 `json:"temperature"`
		TXPower                   float32 `json:"txPower"`
		UpdateAt                  string  `json:"updateAt"`
		Voltage                   float32 `json:"voltage"`
		RawDataBlob               struct {
			blob []int
		} `json:"rawDataBlob"`
	} `json:"tags"`
	BatteryLevel int    `json:"batteryLevel"`
	DeviceID     string `json:"deviceId"`
	EventID      string `json:"eventId"`
	Location     struct {
		Accuracy  float32 `json:"accuracy"`
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"location"`
	Time string `json:"time"`
}

func Parse(b []byte) (*Measurement, error) {
	m := &Measurement{}

	if err := json.Unmarshal(b, m); err != nil {
		return nil, err
	}

	return m, nil
}
