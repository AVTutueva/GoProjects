package painter_calculator

import (
	"fmt"
)

func Scenario() {

	fmt.Print("Welcome to the paint calculator! I'll ask a few questions to help you.\n")

	// ask about walls
	object_title := "wall"
	walls := ReadObjectNumber(object_title)
	wall_size := ReadRectangularObjectInfo(walls, object_title)

	// ask about windows
	object_title = "window"
	windows := ReadObjectNumber(object_title)
	windows_size := ReadRectangularObjectInfo(windows, object_title)

	// ask about doors
	object_title = "door"
	doors := ReadObjectNumber(object_title)
	doors_size := ReadRectangularObjectInfo(doors, object_title)

	square := CalculateSquareToPaint(wall_size, windows_size, doors_size)

	if square <= 0 {
		fmt.Print("I can't calculate the costs because the area of surfaces to be painted is less than 0. Please, restart the program and try again.\n")
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
