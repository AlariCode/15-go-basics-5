package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	result := weather.GetWeather(geoData, format)
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получение %v", expected, result)
	}
}
