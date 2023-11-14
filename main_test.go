package main

import "testing"

func TestCalculateNumberOfPacks(t *testing.T) {
	testCases := []struct {
		orderQuantity int
		expected      int
	}{
		{1, 1},
		{250, 1},
		{251, 1},
		{501, 2},
		{12001, 4},
	}

	for _, tc := range testCases {
		numberOfPacks := calculateNumberOfPacks(tc.orderQuantity)

		if numberOfPacks != tc.expected {
			t.Errorf("Expected %d packs for order quantity %d, but got %d", tc.expected, tc.orderQuantity, numberOfPacks)
		}
	}
}