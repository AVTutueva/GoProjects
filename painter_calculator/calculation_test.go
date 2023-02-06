package painter_calculator

import (
	"fmt"
	"testing"
)

func TestCalculateCostsNormalCase(t *testing.T) {
	got_kg, got_price := CalculateCosts(0.5, 0.5, 0.5, 1)
	if (got_kg != 1) || (got_price != 0.5) {
		t.Errorf("CalculateCosts(0.5, 0.5, 0.5, 1)) = %f, %.2f, wanted 1, 0.50", got_kg, got_price)
	}
}

func TestCalculateCostsZero(t *testing.T) {
	got_kg, got_price := CalculateCosts(0.5, 0.5, 0.5, 0)
	if (got_kg != 0) || (got_price != 0) {
		t.Errorf("CalculateCosts(0.5, 0.5, 0.5, 0)) = %f, %.2f, wanted 0, 0", got_kg, got_price)
	}
}

func TestCalculateRectangularSquares(t *testing.T) {
	input := []float64{3, 3, 6, 5}
	got := CalculateRectangularSquares(input)
	fmt.Println(got)

	if !((got - 39) <= 0.000001) {
		t.Errorf("CalculateRectangularSquares(3, 3, 6, 5) = %.2f, wanted 39", got)
	}
}

func TestCalculateRectangularSquaresZeros(t *testing.T) {
	input := []float64{0, 0, 0, 0}
	got := CalculateRectangularSquares(input)

	if !((0 - got) <= 0.000001) {
		t.Errorf("CalculateRectangularSquares(0, 0, 0, 0) = %.2f, wanted 0", got)
	}
}

func TestCalculateRectangularSquaresSeveralZeros(t *testing.T) {
	input := []float64{1, 2, 0, 0}
	got := CalculateRectangularSquares(input)

	if !((2 - got) <= 0.000001) {
		t.Errorf("CalculateRectangularSquares(1, 2, 0, 0) = %.2f, wanted 2", got)
	}
}
