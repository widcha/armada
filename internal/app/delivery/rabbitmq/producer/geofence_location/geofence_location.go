package geofencelocation

import (
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/streadway/amqp"
	"github.com/widcha/armada/configs"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeofenceEvent struct {
	VehicleID string   `json:"vehicle_id"`
	Event     string   `json:"event"`
	Location  Location `json:"location"`
	Timestamp int64    `json:"timestamp"`
}

func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Radius of Earth in meters
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func CheckAndPublishGeofence(vehicleID string, currentLoc Location, geofenceLoc Location, amqpURL string) error {
	if geofenceLoc == (Location{}) {
		geofenceLoc.Latitude = configs.Get().GeofenceLat
		geofenceLoc.Longitude = configs.Get().GeofenceLong
	}

	distance := HaversineDistance(currentLoc.Latitude, currentLoc.Longitude, geofenceLoc.Latitude, geofenceLoc.Longitude)

	if distance > 50 {
		return nil
	}

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

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

	event := GeofenceEvent{
		VehicleID: vehicleID,
		Event:     "geofence_entry",
		Location:  currentLoc,
		Timestamp: time.Now().Unix(),
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"fleet.events",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Geofence entry published for vehicle %s", vehicleID)
	return nil
}
