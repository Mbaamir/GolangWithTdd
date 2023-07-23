package iterables

func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func ReturnSumArray(numberArrs ...[]int) []int {
	var sums []int
	for _, numbers := range numberArrs {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numberArrs ...[]int) []int {
	sums := []int{}
	for _, arr := range numberArrs {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			tail := arr[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
