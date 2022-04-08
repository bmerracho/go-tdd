package arrayAndSlice

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	sum := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers)
	}

	return sum
}
