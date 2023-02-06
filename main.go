package main

import (
	"fmt"
	"time"
)

func sum(nums []int, ch chan<- int) {
	result := 0
	for _, num := range nums {
		result += num
		time.Sleep(time.Second)
		fmt.Println(num, result)
	}
	ch <- result
}

func main() {
	numbers := []int{1, -1, 2, 3, 10, 5, 4, -4, 12, 0, 7}
	ch := make(chan int, 2)

	go sum(numbers[:len(numbers)/2], ch)
	go sum(numbers[len(numbers)/2:], ch)

	fmt.Println(<-ch + <-ch)

}
