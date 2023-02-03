package exercise1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculator() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter two decimal numbers separated by a space\n")

	int_numbers := []int{} // array to save numbers

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")

		for i := 0; i < len(numbers); i++ { // check numbers

			number, err := strconv.ParseInt(numbers[i], 10, 64)
			if err != nil { // if the output wrong, read again
				fmt.Print("The numbers are incorrect. Try again, please\n")
				break
			} else { // else save number to array
				int_numbers = append(int_numbers, int(number))
			}
		}

		if len(int_numbers) == 2 { // if there are exactly two numbers, then read the operator
			break
		} else {
			fmt.Print("The numbers are incorrect. Try again, please\n")
			int_numbers = []int{}
		}
	}

	fmt.Print("Enter the operator from the list: [+, -, *, /] \n")

	for scanner.Scan() {
		operator := scanner.Text()
		if !strings.Contains("+-*/", operator) {
			fmt.Print("The operator is incorrect. Try again, please\n")
		} else {
			var result float64
			switch {
			case operator == "+":
				result = float64(int_numbers[0]) + float64(int_numbers[1])
			case operator == "-":
				result = float64(int_numbers[0]) - float64(int_numbers[1])
			case operator == "*":
				result = float64(int_numbers[0]) * float64(int_numbers[1])
			case operator == "/":
				if int_numbers[1] == 0 {
					fmt.Print("Cannot divide by 0. Enter another operator, please\n")
					continue
				}
				result = float64(int_numbers[0]) / float64(int_numbers[1])

			}
			fmt.Printf("Result: %.2f", result)
			break
		}
	}

}
