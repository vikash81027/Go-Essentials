package fn

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	transformNumbers(&numbers, double)
	transformNumbers(&numbers, triple)
	fmt.Println(numbers)

	doubleFunc := getTransformFn()
	res := doubleFunc(2)
	fmt.Println(res)
}

// func doubleNumbers(numbers *[]int) {
// 	for i := 0; i < len(*numbers); i++ {
// 		(*numbers)[i] = double((*numbers)[i])
// 	}
// }

type transformFm func(int) int
// func receiving func in argument
func transformNumbers(numbers *[]int, transform transformFm) {
	for i := 0; i < len(*numbers); i++ {
		(*numbers)[i] = transform((*numbers)[i])
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
// func returning func
func getTransformFn() func(int) int {
	return double
}