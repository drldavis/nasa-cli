package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	nasaclient "github.com/drldavis/nasa-cli/api"
	nasaservice "github.com/drldavis/nasa-cli/service"
)

const NasaAPIKey string = "nyObbN41tA4wtOYvOXbBY3YyGQ258j7o8bWhDgOF"

func main() {
	var marsFlag bool
	var solarFlag bool
	flag.BoolVar(&marsFlag, "mars", false, "returns the current weather on mars")
	flag.BoolVar(&solarFlag, "solar", false, "returns data on recent solar flares")
	flag.Parse()
	if marsFlag && solarFlag {
		fmt.Println("Mars flag and solar flag cannot be used together. Please select one.")
		os.Exit(1)
	}
	if !marsFlag && !solarFlag {
		fmt.Println("Please suppy either the -mars or -solar flag")
		os.Exit(1)
	}

	nasaClient := nasaclient.NewNasaClient(NasaAPIKey)
	nasaService := nasaservice.NewNasaService(nasaClient)

	if marsFlag {
		fmt.Println("Getting the mars weather report...")
		if err := nasaService.GetMarsWeather(); err != nil {
			fmt.Println(err)
		}
	} else {
		var startDateInput, endDateInput string
		var startDate, endDate time.Time
		var err error
		for startDateInput == "" {
			fmt.Println("Please enter a start date in the form YYYY-MM-DD")
			fmt.Scan(&startDateInput)
			startDate, err = time.Parse("2006-01-02", startDateInput)
			if err != nil {
				startDateInput = ""
				fmt.Println("Invalid date. Please ensure you are using the format YYYY-MM-DD")
			}
		}
		for endDateInput == "" {
			fmt.Println("Please enter an end date in the form YYYY-MM-DD")
			fmt.Scan(&endDateInput)
			endDate, err = time.Parse("2006-01-02", endDateInput)
			if err != nil {
				endDateInput = ""
				fmt.Println("Invalid date. Please ensure you are using the format YYYY-MM-DD")
			}
		}

		if err := nasaService.GetSolarFlares(startDate, endDate); err != nil {
			fmt.Println(err)
		}
	}
}
