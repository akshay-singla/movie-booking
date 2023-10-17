package booking

import "testing"

func TestCalculateSubtotal(t *testing.T) {
	// Test cases
	tests := []struct {
		showNo        int
		selectedSeats []string
		want          float64
	}{
		{1, []string{"A1", "A2"}, PlatinumCost * 2},        // 2 Platinum seats in show 1
		{2, []string{"A1", "A2"}, PlatinumCost * 2},        // 2 Gold seats in show 2
		{3, []string{"B1", "C2"}, GoldCost + SilverCost},   // 2 Silver seats in show 3
		{1, []string{"A1", "B1"}, PlatinumCost + GoldCost}, // 1 Platinum and 1 Gold seat in show 1
	}

	for _, test := range tests {
		got := calculateSubtotal(test.showNo, test.selectedSeats)
		if got != test.want {
			t.Errorf("calculateSubtotal(%d, %v) = %.2f; want %.2f", test.showNo, test.selectedSeats, got, test.want)
		}
	}
}
