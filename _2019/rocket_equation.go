package _2019

/*
Problem: Day 1: The Tyranny of the Rocket Equation
Link: https://adventofcode.com/2019/day/1
*/
func findTotalFuel(masses []int) int {
	fuels := make([]int, 0)
	for _, mass := range masses {
		fuels = append(fuels, findFuel(mass, false, 0))
	}

	return sumOfGivenArray(fuels)
}

func sumOfGivenArray(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum
}

func findFuel(mass int, additionalFuel bool, sumFuel int) int {
	currentFuel := (mass / 3) - 2

	if !additionalFuel {
		if mass <= 0 || currentFuel <= 0 {
			return 0
		}

		return currentFuel
	}

	if mass <= 0 || currentFuel <= 0 {
		return sumFuel
	}

	return findFuel(currentFuel, additionalFuel, sumFuel+currentFuel)

}
