// Pacakge allergies provides a solution for exercise Allergies on Exercism's Go Track.
package allergies

// Bitmasks for allergies
// Do not change order
const (
	eggs uint = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

// adic holds a map of descriptions to bit masks
var adic = map[string]uint{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

// Allergies converts allergies score to a string array
func Allergies(allergies uint) []string {
	res := []string{}
	for name := range adic {
		if AllergicTo(allergies, name) {
			res = append(res, name)
		}
	}
	return res
}

// AllergicTo checks if allergies score contains given allergen
func AllergicTo(allergies uint, allergen string) bool {
	return allergies&adic[allergen] > 0
}
