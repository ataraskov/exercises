package purchase

const choice_pharse = " is clearly the better choice."

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return option2 + choice_pharse
	}
	return option1 + choice_pharse
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	value := 0.7
	if age < 3 {
		value = 0.8
	} else if age >= 10 {
		value = 0.5
	}

	return originalPrice * value
}
