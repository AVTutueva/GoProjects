package painter_calculator

import "testing"

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
