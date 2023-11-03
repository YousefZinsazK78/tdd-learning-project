package array

func Sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	var sum = make([]int, lengthOfNumbers)

	for i, number := range numbers {
		sum[i] = Sum(number)
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
