package painter_calculator

import (
	"fmt"
)

func Scenario() {
	println()
	fmt.Print("Welcome to the paint calculator! I'll ask a few questions to help you.\n")

	walls := ReadObjectNumber("wall")
	wall_sizes := ReadRectangularObjectInfo(walls, "wall")

	fmt.Print("Thanks! How many windows do you have?.\n")
	windows := ReadObjectNumber("window")
	windows_sizes := ReadRectangularObjectInfo(windows, "window")

	fmt.Print("Thanks! How many doors do you have?.\n")
	doors := ReadObjectNumber("door")
	doors_sizes := ReadRectangularObjectInfo(doors, "door")

	square := CalculateSquareToPaint(wall_sizes, windows_sizes, doors_sizes)

	if square <= 0 {
		fmt.Print("I can't calculate the costs because the area of surfaces to be painted is less than 0. Please, restart the program.\n")
	} else {
		fmt.Print("Great! Now tell me about the paint. What is paint consumption (kg/m2)?\n")
		consumption := ReadFloatObject("consumption")

		fmt.Print("What price (kg)?\n")
		price := ReadFloatObject("price")

		layer := ReadObjectNumber("layer")

		kg, costs := CalculateCosts(square, consumption, price, layer)
		fmt.Printf("To paint %d walls with %d windows and %d doors you will need %.f kg of paint, with total costs of %.2f", walls, windows, doors, kg, costs)
	}

}
