package tristar

import (
	"time"
)

type StopData struct {
	LastUpdate string `json:"lastUpdate"`
	Stops      []Stop `json:"stops"`
}

type Stop struct {
	StopId             int     `json:"stopId"`
	StopCode           string  `json:"stopCode"`
	StopName           string  `json:"stopName"`
	StopShortName      string  `json:"stopShortName"`
	StopDesc           string  `json:"stopDesc"`
	SubName            string  `json:"subName"`
	Date               string  `json:"date"`
	StopLat            float64 `json:"stopLat"`
	StopLon            float64 `json:"stopLon"`
	Type               string  `json:"type"`
	ZoneId             int     `json:"zoneId"`
	ZoneName           string  `json:"zoneName"`
	StopUrl            string  `json:"stopUrl"`
	LocationType       string  `json:"locationType"`
	ParentStation      string  `json:"parentStation"`
	StopTimezone       string  `json:"stopTimezone"`
	WheelchairBoarding int     `json:"wheelchairBoarding"`
	Virtual            int     `json:"virtual"`
	NonPassenger       int     `json:"nonpassenger"`
	Depot              int     `json:"depot"`
	TicketZoneBorder   int     `json:"ticketZoneBorder"`
	OnDemand           int     `json:"onDemand"`
	ActivationDate     string  `json:"activationDate"`
}

type StopDepartures struct {
	LastUpdate time.Time   `json:"lastUpdate"`
	Departures []Departure `json:"departures"`
}

type Departure struct {
	ID                     string    `json:"id"`
	DelayInSeconds         int       `json:"delayInSeconds"`
	EstimatedTime          time.Time `json:"estimatedTime"`
	Headsign               string    `json:"headsign"`
	RouteID                int       `json:"routeId"`
	RouteShortName         string    `json:"routeShortName"`
	ScheduledTripStartTime time.Time `json:"scheduledTripStartTime"`
	TripID                 int       `json:"tripId"`
	Status                 string    `json:"status"`
	TheoreticalTime        time.Time `json:"theoreticalTime"`
	Timestamp              time.Time `json:"timestamp"`
	Trip                   int       `json:"trip"`
	VehicleCode            int       `json:"vehicleCode"`
	VehicleID              int       `json:"vehicleId"`
	VehicleService         string    `json:"vehicleService"`
}
