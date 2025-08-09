package recursion

import "fmt"

func main() {
	res := factorial(5)
	fmt.Println(res)
}


// factorial of number with recursion
func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n - 1)
}