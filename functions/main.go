package main

import "fmt"

func main() {
	numbers := []int{1, 2, 4, 5}
	argSlice := []int{1, 2, 4, 5, 9}
	sum := sliceSum(numbers)
	fmt.Println(sum)
	
	anotherSum := variadicSum(1, 2, 3, 4)
	newSum := variadicSum(argSlice...) // unpack slice into comma separated values

	fmt.Println(anotherSum, newSum)
}

func sliceSum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// variadic function takes comma separated args and packs them into a slice
func variadicSum(numbers... int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}