// Package weather provides means to forecast weather for a given city.
package weather

// CurrentCondition holds value of current weather concitions.
var CurrentCondition string

// CurrentLocation represent location as a string, i.e. city name.
var CurrentLocation string

// Forecast takes city and contiion and produces a weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
