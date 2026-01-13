package tristar

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func FetchAvailableStops() ([]Stop, error) {
	stopsJson, err := http.Get("https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/4c4025f0-01bf-41f7-a39f-d156d201b82b/download/stops.json")
	if err != nil {
		return nil, err
	}
	defer stopsJson.Body.Close()

	var stopsMap map[string]StopData

	err = json.NewDecoder(stopsJson.Body).Decode(&stopsMap)
	if err != nil {
		return nil, err
	}

	todayDate := time.Now().Format("2006-01-02")

	// fmt.Printf("%s\n", todayDate)

	if stopData, ok := stopsMap[todayDate]; ok {
		return stopData.Stops, nil
	}

	return nil, errors.New("Dataset does not contain data for today")
}

func FetchStopData(stopId int) ([]Departure, error) {

	stopJson, err := http.Get(fmt.Sprintf("https://ckan2.multimediagdansk.pl/departures?stopId=%d", stopId))
	if err != nil {
		return nil, err
	}
	defer stopJson.Body.Close()

	var stopData StopDepartures

	err = json.NewDecoder(stopJson.Body).Decode(&stopData)
	if err != nil {
		return nil, err
	}

	return stopData.Departures, nil
}
