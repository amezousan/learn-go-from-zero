package main

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSum {
		lastItemIndex := 0
		if len(numbers) != 0 {
			lastItemIndex = len(numbers) - 1
		}
		// fmt.Printf("Check: %v, %v, %v\n", numbersToSum, numbers, lastItemNum)
		tail := numbers[lastItemIndex:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
