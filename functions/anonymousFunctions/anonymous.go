package anonymouseFunctions

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(3)(2)
	fmt.Println(double)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}