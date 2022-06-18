package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (tailSums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tailSums = append(tailSums, Sum(numbers[1:]))
		}
	}
	return
}
