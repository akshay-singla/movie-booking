package booking

import "testing"

func TestIsSeatBooked(t *testing.T) {
	// Set up test data
	bookedSeats[1] = []string{"A1"}
	bookedSeats[2] = []string{"B2"}
	// Test cases
	tests := []struct {
		showNo int
		seat   string
		want   bool
	}{
		{1, "A1", true},  // A1 is booked in show 1
		{1, "B1", false}, // B1 is not booked in show 1
		{2, "A2", false}, // A2 is not booked in show 2
		{2, "B2", true},  // B2 is booked in show 2
	}

	for _, test := range tests {
		got := isSeatBooked(test.showNo, test.seat)
		if got != test.want {
			t.Errorf("isSeatBooked(%d, %s) = %t; want %t", test.showNo, test.seat, got, test.want)
		}
	}
}

func TestAreSeatsAvailable(t *testing.T) {
	// Set up test data
	bookedSeats[1] = []string{"A1", "B2"}
	bookedSeats[2] = []string{"A1", "B2"}

	// Test cases
	tests := []struct {
		showNo        int
		selectedSeats []string
		want          bool
	}{
		{1, []string{"A2", "B1"}, true},       // A2 and B1 are available in show 1
		{1, []string{"A1", "B2"}, false},      // A1 is booked in show 1
		{2, []string{"A2", "B1"}, true},       // A2 and B1 are available in show 2
		{2, []string{"A1", "B2"}, false},      // B2 is booked in show 2
		{3, []string{"A1", "B1", "C1"}, true}, // A1, B1, and C1 are available in show 3
	}

	for _, test := range tests {
		got := areSeatsAvailable(test.showNo, test.selectedSeats)
		if got != test.want {
			t.Errorf("areSeatsAvailable(%d, %v) = %t; want %t", test.showNo, test.selectedSeats, got, test.want)
		}
	}
}

func TestParseSeats(t *testing.T) {
	// Test cases
	tests := []struct {
		input string
		want  []string
	}{
		{"A1,A2,A3", []string{"A1", "A2", "A3"}},
		{"A1, A2 ,A3 ", []string{"A1", "A2", "A3"}},
		{"A1,A2,A3", []string{"A1", "A2", "A3"}},
		{"B1,B2,B3", []string{"B1", "B2", "B3"}},
	}

	for _, test := range tests {
		got := parseSeats(test.input)
		if len(got) != len(test.want) {
			t.Errorf("parseSeats(%s) = %v; want %v", test.input, got, test.want)
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("parseSeats(%s) = %v; want %v", test.input, got, test.want)
			}
		}
	}
}
