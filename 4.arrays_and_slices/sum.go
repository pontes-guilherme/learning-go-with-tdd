package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(sliceOfNumbers ...[]int) []int {
	var results []int

	for _, numbers := range sliceOfNumbers {
		results = append(results, Sum(numbers))
	}

	return results
}

func SumAllTails(sliceOfNumbers ...[]int) []int {
	var results []int

	for _, numbers := range sliceOfNumbers {
		if len(numbers) > 0 {
			tail := numbers[1:]
			results = append(results, Sum(tail))
		} else {
			results = append(results, 0)
		}
	}

	return results
}
