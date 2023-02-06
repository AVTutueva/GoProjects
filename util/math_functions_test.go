package util

import "testing"

func TestAbsNegative(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d, wanted 1", got)
	}
}

func TestAbsPositive(t *testing.T) {
	got := Abs(1)
	if got != 1 {
		t.Errorf("Abs(1) = %d, wanted 1", got)
	}
}

func TestAbsZero(t *testing.T) {
	got := Abs(0)
	if got != 0 {
		t.Errorf("Abs(0) = %d, wanted 0", got)
	}
}

func TestMax_1_2(t *testing.T) {
	got := Max(1, 2)
	if got != 2 {
		t.Errorf("Max(1, 2) = %d, wanted 2", got)
	}
}

func TestMax_2_1(t *testing.T) {
	got := Max(2, 1)
	if got != 2 {
		t.Errorf("Max(2, 1) = %d, wanted 2", got)
	}
}

func TestMax_2_2(t *testing.T) {
	got := Max(2, 2)
	if got != 2 {
		t.Errorf("Max(2, 2) = %d, wanted 2", got)
	}
}

func TestMaxZeros(t *testing.T) {
	got := Max(0, 0)
	if got != 0 {
		t.Errorf("Max(0, 0) = %d, wanted 0", got)
	}
}

func TestMin_1_2(t *testing.T) {
	got := Min(1, 2)
	if got != 1 {
		t.Errorf("Min(1, 2) = %d, wanted 1", got)
	}
}

func TestMin_2_1(t *testing.T) {
	got := Min(2, 1)
	if got != 1 {
		t.Errorf("Min(2, 1) = %d, wanted 1", got)
	}
}

func TestMin_2_2(t *testing.T) {
	got := Min(2, 2)
	if got != 2 {
		t.Errorf("Min(2, 2) = %d, wanted 2", got)
	}
}

func TestMinZeros(t *testing.T) {
	got := Min(0, 0)
	if got != 0 {
		t.Errorf("Min(0, 0) = %d, wanted 0", got)
	}
}
