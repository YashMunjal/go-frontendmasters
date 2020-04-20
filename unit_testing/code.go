package unit_testing

import "fmt"

func printNum(num int) {
	fmt.Println("Current No.", num)
}

func addNums(nums ...int) int {
	total := 0
	for _, v := range nums {
		printNum(v)
		total = +v
	}
	return total
}
