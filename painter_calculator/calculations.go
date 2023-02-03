package painter_calculator

import (
	"math"
)

func CalculateSquareToPaint(walls_size []float64, windows_size []float64, doors_size []float64) float64 {

	return CalculateRectangularSquares(walls_size) - CalculateRectangularSquares(windows_size) - CalculateRectangularSquares(doors_size)
}

func CalculateCosts(square float64, consumption float64, price float64, layer int64) (float64, float64) {
	kg := math.Ceil(square * float64(layer) * consumption)
	return kg, kg * price
}

func CalculateRectangularSquares(object_dimensions []float64) float64 {
	var square float64
	for i := 0; i < len(object_dimensions); i = i + 2 {
		square = square + object_dimensions[i]*object_dimensions[i+1]
	}
	return square
}
