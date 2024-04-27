package nasaservice

import (
	"fmt"
	"time"

	"github.com/drldavis/nasa-cli/models"
)

type nasaClient interface {
	GetMarsWeather() (models.MarsWeather, error)
	GetSolarFlares(startDate time.Time, endDate time.Time) ([]models.SolarFlare, error)
}

type nasaService struct {
	client nasaClient
}

func NewNasaService(client nasaClient) nasaService {
	return nasaService{client: client}
}

func (s nasaService) GetMarsWeather() error {
	marsWeather, err := s.client.GetMarsWeather()
	if err != nil {
		return fmt.Errorf("error occured while getting mars weather: %w", err)
	}
	fmt.Println("The most recently available mars weather report...")
	fmt.Println("Temperature: ", marsWeather.Temperature)
	fmt.Println("Wind Speed: ", marsWeather.WindSpeed)
	fmt.Println("Pressure: ", marsWeather.Pressure)
	fmt.Println("Season: ", marsWeather.Season)

	return nil
}

func (s nasaService) GetSolarFlares(startDate time.Time, endDate time.Time) error {
	solarFlares, err := s.client.GetSolarFlares(startDate, endDate)
	if err != nil {
		return fmt.Errorf("error occured while getting solar flares report: %w", err)
	}
	fmt.Println("The following data was found on solar flares within the specified date range")
	for _, sf := range solarFlares {
		fmt.Println("Begin Time: ", sf.BeginTime)
		fmt.Println("Peak Time: ", sf.PeakTime)
		fmt.Println("End Time: ", sf.EndTime)
		fmt.Println("Note: ", sf.Note)
		fmt.Println("------------------------------------------")
	}
	return nil
}
