package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	if productionRate == 0 || successRate == 0 {
		return 0
	}
	return int(float64(productionRate)*successRate) / 60 / 100
}

// CalculateCost works out the cost of producing the given number of cars.
const costOfIndividual = 10000
const costOfBatch = 95000
const countOfBatch = 10

func CalculateCost(carsCount int) uint {
	return uint(costOfBatch*(carsCount/countOfBatch) + costOfIndividual*(carsCount%countOfBatch))
}
