package booking

import "strings"

var (
	shows = map[int][]string{
		1: {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "A9", "B1", "B2", "B3", "B4", "B5", "B6", "C2", "C3", "C4", "C5", "C6", "C7"},
		2: {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "B2", "B3", "B4", "B5", "B6", "C1", "C2", "C3", "C4", "C5", "C6", "C7"},
		3: {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9"},
	}
	bookedSeats = map[int][]string{
		1: {},
		2: {},
		3: {},
	}
)

func isSeatBooked(showNo int, seat string) bool {
	for _, bookedSeat := range bookedSeats[showNo] {
		if bookedSeat == seat {
			return true
		}
	}
	return false
}

func areSeatsAvailable(showNo int, selectedSeats []string) bool {
	for _, seat := range selectedSeats {
		if isSeatBooked(showNo, seat) {
			return false
		}
	}
	return true
}

func bookSeats(showNo int, selectedSeats []string) {
	bookedSeats[showNo] = append(bookedSeats[showNo], selectedSeats...)
}

func parseSeats(seatsInput string) []string {
	seats := make([]string, 0)
	seatsInput = strings.ReplaceAll(seatsInput, " ", "")
	seatsInputs := strings.Split(seatsInput, ",")
	for _, seat := range seatsInputs {
		seats = append(seats, string(seat))
	}
	return seats
}
