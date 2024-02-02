package lasagna

const (
	NoodlesPerLayer   = 50
	SaucePerLayer     = 0.2
	PortionsPerRecipe = 2
)

func PreparationTime(layers []string, avgTime int) int {
	t := avgTime
	if t < 1 {
		t = 2
	}
	return len(layers) * t
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, l := range layers {
		if l == "noodles" {
			noodles += NoodlesPerLayer
		} else if l == "sauce" {
			sauce += SaucePerLayer
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	q := []float64{}
	for i, _ := range quantities {
		q = append(q, quantities[i]*float64(portions)/PortionsPerRecipe)
	}
	return q
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
