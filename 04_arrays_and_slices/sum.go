package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, nums := range numbers {
		sums[i] = Sum(nums)
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sumTails []int
	for _, nums := range numbers {
		if len(nums) == 0 {
			sumTails = append(sumTails, 0)
		} else {
			sumTails = append(sumTails, Sum(nums[1:]))
		}
	}

	return sumTails
}
