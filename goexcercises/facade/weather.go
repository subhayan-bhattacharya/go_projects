package facade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Latitude float64
type Longitude float64

func (l Longitude) isValid() bool {
	return l >= -180.0 && l <= 180.0
}

func (l Latitude) isValid() bool {
	return l >= -90.0 && l <= 90.0
}

type Coordinates struct {
	Latitude  Latitude  `json:"latitude"`
	Longitude Longitude `json:"longitude"`
}

func (c Coordinates) isValid() bool {
	return c.Latitude.isValid() && c.Longitude.isValid()
}

type Weather struct {
	CityName    string      `json:"city_name"`
	Temperature float64     `json:"temperature"`
	Sunrise     string      `json:"sunrise"`
	Sunset      string      `json:"sunset"`
	Coordinates Coordinates `json:"coordinates"`
}

func formatOpenMeteoTime(apiTime string) (string, error) {
	t, err := time.Parse("2006-01-02T15:04", apiTime)

	if err != nil {
		return "", err
	}
	return t.Format("15:04"), nil
}

type OpenApiResponse struct {
	Current struct {
		Temperature2m float64 `json:"temperature_2m"`
	} `json:"current"`
	Daily struct {
		Sunrise []string `json:"sunrise"`
		Sunset  []string `json:"sunset"`
	} `json:"daily"`
}

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
}

type WeatherFacade struct {
	Client *http.Client
}

func (w WeatherFacade) GetByCityAndCountryCode(city, countryCode string) (Weather, error) {
	var weather Weather
	coordinates, err := w.getCoordinatesForCityAndCountry(city, countryCode)
	if err != nil {
		return weather, fmt.Errorf("latitude and longiture are not valid %w", err)
	}
	baseUrl := "https://api.open-meteo.com/v1/forecast"
	params := url.Values{}
	params.Set("latitude", fmt.Sprintf("%f", coordinates.Latitude))
	params.Set("longitude", fmt.Sprintf("%f", coordinates.Longitude))
	params.Add("daily", "sunset")
	params.Add("daily", "sunrise")
	params.Set("current", "temperature_2m")
	params.Set("timezone", "auto")
	fullURL := baseUrl + "?" + params.Encode()
	response, err := w.Client.Get(fullURL)
	defer response.Body.Close()
	errorString := fmt.Sprintf("could not get the latitude and longitude for city %s and country %s", city, countryCode)
	if err != nil {
		return weather, fmt.Errorf(errorString + fmt.Sprintf("%w", err))
	}
	var apiResponse OpenApiResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return weather, fmt.Errorf(errorString + fmt.Sprintf("%w", err))
	}
	sunrise, _ := formatOpenMeteoTime(apiResponse.Daily.Sunrise[0])
	sunset, _ := formatOpenMeteoTime(apiResponse.Daily.Sunset[0])
	weather = Weather{
		CityName:    city,
		Temperature: apiResponse.Current.Temperature2m,
		Sunrise:     sunrise,
		Sunset:      sunset,
		Coordinates: coordinates,
	}
	return weather, nil
}

type apiCoordinates struct {
	Results []Coordinates `json:"results"`
}

func (w WeatherFacade) getCoordinatesForCityAndCountry(city, countryCode string) (Coordinates, error) {
	var coordinates Coordinates
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&country_code=%s&count=1", city, countryCode)
	response, err := w.Client.Get(url)
	defer response.Body.Close()
	errorString := fmt.Sprintf("could not get the latitude and longitude for city %s and country %s", city, countryCode)
	if err != nil {
		return coordinates, fmt.Errorf(errorString + fmt.Sprintf("%w", err))
	}
	if response.StatusCode != http.StatusOK {
		return coordinates, fmt.Errorf(errorString + fmt.Sprintf("%w", err))
	}
	var apiCoordinates apiCoordinates
	err = json.NewDecoder(response.Body).Decode(&apiCoordinates)
	if err != nil {
		return coordinates, fmt.Errorf(errorString)
	}
	if len(apiCoordinates.Results) == 0 {
		return coordinates, fmt.Errorf(errorString + fmt.Sprintf("%w", err))
	}
	coordinates = apiCoordinates.Results[0]
	if !coordinates.isValid() {
		return Coordinates{}, fmt.Errorf(errorString + fmt.Sprintf("%w", err))
	}
	return coordinates, nil
}
