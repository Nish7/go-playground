package array

func Sum(numbers []int) int {
	sum := 0 

	for _, n := range numbers {
		sum += n 
	}  

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, n := numbersToSum {
		sum = append(sum, Sum(n))
	}

	return sum
}

func SumAllTail(numbersToSum ...[]int) []int {
	var sum []int

	for _, n := range numbersToSum {
		if len(n) == 0 {
			sum = append(sum, 0)
			continue
		}

		tail := n[1:]
		sum = append(sum, Sum(tail))
	}

	return sum
}
