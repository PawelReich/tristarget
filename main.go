package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PawelReich/tristarget/tristar"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'find' or 'departures'")
		os.Exit(1)
	}

	findCmd := flag.NewFlagSet("find", flag.ExitOnError)
	findName := findCmd.String("name", "", "Name of the stop")

	departuresCmd := flag.NewFlagSet("departures", flag.ExitOnError)
	departuresStopId := departuresCmd.Int("stop", -1, "ID of the stop")
	departuresRouteId := departuresCmd.Int("route", -1, "ID of the route")
	departuresLimit := departuresCmd.Int("limit", -1, "Maximum number of returned items")

	switch os.Args[1] {
	case "find":
		err := findCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
		if *findName == "" {
			findCmd.PrintDefaults()
			os.Exit(1)
		}
		err = FindStop(*findName)
		if err != nil {
			panic(err)
		}

	case "departures":
		err := departuresCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
		if *departuresStopId == -1 {
			departuresCmd.PrintDefaults()
			os.Exit(1)
		}
		err = FindDepartures(*departuresStopId, *departuresRouteId, *departuresLimit)
		if err != nil {
			panic(err)
		}
	}
}

func FindDepartures(stopId int, routeId int, limit int) error {
	stopData, err := tristar.FetchStopData(stopId)
	if err != nil {
		return err
	}

	printed := 0
	for _, departure := range stopData {
		if routeId == -1 || routeId == departure.RouteID {
			printed++
			if limit != -1 && printed > limit {
				break
			}
			fmt.Printf("%d %s\n", departure.RouteID, departure.EstimatedTime.Local().Format("15:04:05"))
		}
	}
	return nil
}

func FindStop(searchTerm string) error {
	stops, err := tristar.FetchAvailableStops()
	if err != nil {
		return err
	}

	searchTermLower := strings.ToLower(searchTerm)

	for _, stop := range stops {
		nameLower := fmt.Sprintf("%s %s", strings.ToLower(stop.StopName), stop.SubName)
		if strings.Contains(nameLower, searchTermLower) {
			fmt.Printf("%d: %s %s\n", stop.StopId, stop.StopName, stop.SubName)
		}
	}

	return nil
}
