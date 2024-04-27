package nasaclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/drldavis/nasa-cli/models"
)

type NasaClient struct {
	apiKey string
}

func NewNasaClient(apiKey string) NasaClient {
	return NasaClient{apiKey: apiKey}
}

func (c NasaClient) GetMarsWeather() (models.MarsWeather, error) {
	query := fmt.Sprintf("https://api.nasa.gov/insight_weather/?api_key=%s&feedtype=json&ver=1.0", c.apiKey)

	responseBody, err := makeRequest(query)
	if err != nil {
		return models.MarsWeather{}, fmt.Errorf("error occured while getting mars data: %w", err)
	}

	var responseMap map[string]json.RawMessage
	if err = json.Unmarshal(responseBody, &responseMap); err != nil {
		return models.MarsWeather{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	var solKeys []string
	if err = json.Unmarshal(responseMap["sol_keys"], &solKeys); err != nil {
		return models.MarsWeather{}, fmt.Errorf("error unmarshaling response")
	}

	if len(solKeys) == 0 {
		return models.MarsWeather{}, fmt.Errorf("nasa does not currently have enough data to provide an accurate weather report on mars")
	}

	latestKey := solKeys[len(solKeys)-1]

	var sol models.Sol
	if err = json.Unmarshal(responseMap[latestKey], &sol); err != nil {
		return models.MarsWeather{}, fmt.Errorf("error unmarshaling response")
	}

	return models.MarsWeather{
		Temperature: sol.Temperature.Average,
		WindSpeed:   sol.WindSpeed.Average,
		Pressure:    sol.Pressure.Average,
		Season:      sol.Season,
	}, nil
}

func (c NasaClient) GetSolarFlares(startDate time.Time, endDate time.Time) ([]models.SolarFlare, error) {
	query := fmt.Sprintf("https://api.nasa.gov/DONKI/FLR?startDate=%s&endDate=%s&api_key=%s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), c.apiKey)
	responseBody, err := makeRequest(query)
	if err != nil {
		return nil, fmt.Errorf("error getting solar flare data: %w", err)
	}

	var solarFlares []models.SolarFlare
	if err = json.Unmarshal(responseBody, &solarFlares); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return solarFlares, nil
}

func makeRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error occured while querying api: %w", err)
	}

	if response.StatusCode == 429 {
		return nil, fmt.Errorf("you have reached your maximum amount of queries for the hour")
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error occured while reading response body: %w", err)
	}

	return responseBody, nil
}
