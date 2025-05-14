package geofence

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type GeofenceEvent struct {
	VehicleID string `json:"vehicle_id"`
	Event     string `json:"event"`
	Location  struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Timestamp int64 `json:"timestamp"`
}

func StartGeofenceWorker(amqpURL string) error {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(
		"fleet.events",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"geofence_alerts",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		"",
		"fleet.events",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("Geofence worker listening for events")

	go func() {
		for d := range msgs {
			var event GeofenceEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Println("Failed to parse event:", err)
				continue
			}

			log.Printf("[GEOFENCE] Vehicle %s entered zone at (%f, %f) at %d\n",
				event.VehicleID, event.Location.Latitude, event.Location.Longitude, event.Timestamp)
		}
	}()

	select {}
}
