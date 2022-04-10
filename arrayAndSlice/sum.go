package arrayAndSlice

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}

	return sum
}

func SumAllTrails(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(numbers[1:]))
		}
	}

	return sum
}
