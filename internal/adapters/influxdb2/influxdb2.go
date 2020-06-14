package influxdb2

import (
	"context"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/majori/ruuvitag-gateway/internal/ruuvi"
)

type InfluxDB2Adapter struct {
	API influxdb2.WriteApiBlocking
}

func New(url, token, organization, bucket string) *InfluxDB2Adapter {
	client := influxdb2.NewClient(url, token)
	defer client.Close()

	api := client.WriteApiBlocking(organization, bucket)

	return &InfluxDB2Adapter{api}
}

func (a *InfluxDB2Adapter) Save(m *ruuvi.Measurement) error {
	for _, tag := range m.Tags {
		tags := make(map[string]string)
		tags["mac"] = tag.ID
		tags["name"] = tag.Name
		tags["dataFormat"] = strconv.Itoa(tag.DataFormat)

		values := make(map[string]interface{})
		values["temperature"] = tag.Temperature
		values["humidity"] = tag.Humidity
		values["pressure"] = tag.Pressure
		values["rssi"] = tag.RSSI
		values["voltage"] = tag.Voltage
		values["movementCounter"] = tag.MovementCounter
		values["measurementSequenceNumber"] = tag.MeasurementSequenceNumber
		values["accelerationX"] = tag.AccelX
		values["accelerationY"] = tag.AccelY
		values["accelerationZ"] = tag.AccelZ

		// TODO: Use time from measurement
		p := influxdb2.NewPoint("ruuvi_measurements", tags, values, time.Now())

		if err := a.API.WritePoint(context.Background(), p); err != nil {
			return err
		}
	}

	return nil
}
