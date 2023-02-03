package painter_calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadObjectNumber(object_title string) int64 {

	switch {
	case object_title == "layer":
		fmt.Printf("How many %ss are you going to paint?: ", object_title)
	default:
		fmt.Printf("How many %ss do you have? Please enter an integer: ", object_title)
	}

	var objects int64

	error_message := "The number of " + object_title + "s is incorrect. Try again, please\n"

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		temp, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			fmt.Print(error_message)
		} else {
			if temp > 0 {
				objects = temp
				break
			} else if (temp == 0) && (object_title != "wall") && (object_title != "layer") {
				objects = temp
				break
			} else {
				fmt.Print(error_message)
			}

		}
	}

	fmt.Print("\nWonderful, thank you!\n")
	println()
	return objects
}

func ReadRectangularObjectInfo(objects_number int64, object_title string) []float64 {
	if objects_number == 0 {
		return []float64{}
	}

	object_size := []float64{} // odd indices - length, even indices - height

	fmt.Printf("Please enter width and height of each %s by space on a separate line.\n", object_title)

	error_message := "The sizes are incorrect. Try again, please\n"

	for i := int64(1); i <= objects_number; i++ { // loop by objects
		fmt.Printf("Width and height of %s %d in m: ", object_title, i)

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			line := scanner.Text()
			sizes := strings.Split(line, " ")

			if len(sizes) != 2 {
				fmt.Print(error_message)
				fmt.Printf("Width and height of %s %d in m: ", object_title, i)
			} else {
				width, err1 := strconv.ParseFloat(sizes[0], 64)
				height, err2 := strconv.ParseFloat(sizes[1], 64)

				if (err1 != nil) || (err2 != nil) || (width == 0) || (height == 0) {
					fmt.Print(error_message)
					fmt.Printf("Width and height of %s %d in m: ", object_title, i)
				} else {
					object_size = append(object_size, width, height)
					break
				}
			}

		}
	}
	println()
	return object_size
}

func ReadFloatObject(object_title string) float64 {

	// one more comment
	var float_value float64
	fmt.Printf("Please enter the value of %s: ", object_title)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if (err != nil) || (value <= 0) {
			fmt.Print("The value is incorrect. Try again, please\n")
			fmt.Printf("Please enter the value of %s: \n", object_title)
		} else {
			float_value = value
			break
		}
	}
	println()
	return float_value
}
