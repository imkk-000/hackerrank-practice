package main_test

import (
	dojo "dojo"
	"testing"
)

func TestCalculateTotalCostInputMealCost12Point00TipPercent20TaxPercent8ShouldReturn16(t *testing.T) {
	expectedMealCost := "16"

	actualMealCost := dojo.CalculateTotalCost(12.00, 25, 8)

	if expectedMealCost != actualMealCost {
		t.Errorf("expect '%s' but it got '%s'", expectedMealCost, actualMealCost)
	}
}

func TestCalculateTotalCostInputMealCost12Point00TipPercent20TaxPercent8ShouldReturn15(t *testing.T) {
	expectedMealCost := "15"

	actualMealCost := dojo.CalculateTotalCost(12.00, 20, 8)

	if expectedMealCost != actualMealCost {
		t.Errorf("expect '%s' but it got '%s'", expectedMealCost, actualMealCost)
	}
}
