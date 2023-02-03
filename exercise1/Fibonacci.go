package exercise1

import "fmt"

func Fib(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		result = append(result, FibRec(i))
	}
	return result
}

func FibRec(n int) int {
	if n < 2 {
		return n
	}
	return FibRec(n-1) + FibRec(n-2)
}

func FibMain() {
	fmt.Println(Fib(10))

}
