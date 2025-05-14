package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jmoiron/sqlx"
	"github.com/widcha/armada/configs"
	geofencelocation "github.com/widcha/armada/internal/app/delivery/rabbitmq/producer/geofence_location"
)

type LocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

type Subscriber struct {
	DB *sqlx.DB
}

func NewSubscriber(db *sqlx.DB) *Subscriber {
	return &Subscriber{DB: db}
}

func (s *Subscriber) Start() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(configs.Get().MqttBroker)
	opts.SetClientID("armada-subscriber")
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
		if token := c.Subscribe("/fleet/vehicle/+/location", 1, s.handleMessage); token.Wait() && token.Error() != nil {
			log.Println("Subscription error:", token.Error())
		}
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("MQTT connection failed:", token.Error())
	}
}

func (s *Subscriber) handleMessage(client mqtt.Client, msg mqtt.Message) {
	var loc LocationPayload

	if err := json.Unmarshal(msg.Payload(), &loc); err != nil {
		log.Println("Invalid JSON payload:", err)
		return
	}

	if loc.VehicleID == "" || loc.Timestamp == 0 {
		log.Println("Invalid data: missing required fields")
		return
	}

	query := `
		INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp)
		VALUES (:vehicle_id, :latitude, :longitude, to_timestamp(:timestamp))
	`

	params := map[string]interface{}{
		"vehicle_id": loc.VehicleID,
		"latitude":   loc.Latitude,
		"longitude":  loc.Longitude,
		"timestamp":  loc.Timestamp,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := s.DB.NamedExecContext(ctx, query, params); err != nil {
		log.Println("Failed to insert to DB:", err)
		return
	}

	fmt.Printf("Stored: %s @ (%f, %f)\n", loc.VehicleID, loc.Latitude, loc.Longitude)

	// Publish to RabbitMQ if more than 50 meters from geofence
	if err := geofencelocation.CheckAndPublishGeofence(loc.VehicleID, geofencelocation.Location{Latitude: loc.Latitude, Longitude: loc.Longitude}, geofencelocation.Location{}, configs.Get().RabbitMQ); err != nil {
		log.Println("Failed to publish to RabbitMQ:", err)
	}
}
