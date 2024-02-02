package space

import "strings"

type Planet string

var Orbit = map[string]float64{
	"mercury": 0.2408467,
	"venus":   0.61519726,
	"earth":   1,
	"mars":    1.8808158,
	"jupiter": 11.862615,
	"saturn":  29.447498,
	"uranus":  84.016846,
	"neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	p := strings.ToLower(string(planet))
	if Orbit[p] == 0 {
		return -1
	}
	return float64(seconds) / (Orbit[p] * 31557600)
}
