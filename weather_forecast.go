package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/joho/godotenv"
)

func weatherForecast() (*[]Weather, error) {
	apiURL, err := getApiURL()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API HTTP error %d: %s", resp.StatusCode, string(body))
	}

	var weatherResponse WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	if weatherResponse.Cod != "200" {
		return nil, fmt.Errorf("API error: %v", weatherResponse.Message)
	}

	weathersPredictions, err := processWeatherResponseData(&weatherResponse)
	if err != nil {
		return nil, fmt.Errorf("error processing weather data: %v", err)
	}

	return weathersPredictions, nil
}


func getApiURL() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	OPEN_WEATHER_MAP_API_KEY:= os.Getenv("OPEN_WEATHER_MAP_API_KEY")

	if OPEN_WEATHER_MAP_API_KEY== "" {
		return "", fmt.Errorf("API key not found in .env file")
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/forecast?q=Jakarta&appid=%s&units=metric",
		OPEN_WEATHER_MAP_API_KEY,
	)

	return url, nil
}

func processWeatherResponseData(data *WeatherResponse) (*[]Weather, error) {
    dailyTemps := make(map[string][]float64)
    
    for _, forecast := range data.List {
        date := time.Unix(int64(forecast.Dt), 0)
        dateKey := date.Format("2006-01-02")
        dailyTemps[dateKey] = append(dailyTemps[dateKey], forecast.Main.Temp)
    }

    var weathers []Weather
    for dateKey, temps := range dailyTemps {
        date, err := time.Parse("2006-01-02", dateKey)
        if err != nil {
            return nil, fmt.Errorf("error parsing date key: %v", err)
        }

        var sum float64
        for _, temp := range temps {
            sum += temp
        }
        avgTemp := sum / float64(len(temps))
        
        weathers = append(weathers, Weather{
            Date: date,
            Temp: avgTemp,
        })
    }
    
	// For consistent output order
    sort.Slice(weathers, func(i, j int) bool {
        return weathers[i].Date.Before(weathers[j].Date)
    })
    
    return &weathers, nil
}