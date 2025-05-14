package main

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type VehicleLocation struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	broker := "tcp://mqtt-broker:1883"
	clientID := "mock-vehicle-publisher"
	vehicleID := "B1234XYZ"
	topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to connect to MQTT broker:", token.Error())
		return
	}
	fmt.Println("Connected to Mosquitto broker at", broker)

	latitude := -6.2088
	longitude := 106.8456

	for {
		// latitude += (rand.Float64() - 0.5) / 1000
		// longitude += (rand.Float64() - 0.5) / 1000

		payload := VehicleLocation{
			VehicleID: vehicleID,
			Latitude:  latitude,
			Longitude: longitude,
			Timestamp: time.Now().Unix(),
		}

		data, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Failed to marshal JSON:", err)
			continue
		}

		token := client.Publish(topic, 1, false, data)
		token.Wait()

		fmt.Printf("Published: %s\n", string(data))
		time.Sleep(2 * time.Second)
	}
}
