package main

import "fmt"

func sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + sum(numbers[1:])
}

func main() {
	toSum := []int{2, 4, 6}

	fmt.Println(sum(toSum))
}
