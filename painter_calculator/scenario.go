package painter_calculator

import (
	"fmt"
)

func Scenario() {

	fmt.Print("Welcome to the paint calculator! I'll ask a few questions to help you.\n")

	objects := []string{"wall", "window", "door"}

	object_quantity := make(map[string]int64)
	object_sizes := make(map[string][]float64)

	// read info about objects
	for _, object := range objects {
		object_quantity[object] = ReadObjectNumber(object)
		object_sizes[object] = ReadRectangularObjectInfo(object_quantity[object], object)
	}

	// calculate square to paint
	square := CalculateSquareToPaint(object_sizes)
	if square <= 0 {
		// if it is negative, print error
		fmt.Print("I can't calculate the costs because the area of surfaces to be painted is less than 0. Please, restart the program and try again.\n")
	} else {
		// ask details about paint
		fmt.Print("Great! Now tell me about the paint. What is paint consumption (kg/m2)? \n")
		consumption := ReadFloatObject("consumption")

		fmt.Print("What price (kg)? ")
		price := ReadFloatObject("price")

		layer := ReadObjectNumber("layer")

		kg, costs := CalculateCosts(square, consumption, price, layer)
		fmt.Printf("To paint %d walls with %d windows and %d doors you will need %.f kg of paint, with total costs of %.2f",
			int(object_quantity["wall"]), int(object_quantity["window"]), int(object_quantity["door"]), kg, costs)
	}

}
