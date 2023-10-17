package booking

import "fmt"

func Movie() {

	var totalRevenue, totalServiceTax, totalSwachhCess, totalKrishiCess float64

	for {
		fmt.Print("Enter Show no: ")
		var showNo int
		fmt.Scanln(&showNo)
		if showNo < 1 || showNo > 3 {
			fmt.Println("Invalid Show Number. Please enter a valid show number.")
			continue
		}

		fmt.Println("Available Seats:")
		fmt.Println("All Seats:")
		for _, seat := range shows[showNo] {
			if !isSeatBooked(showNo, seat) {
				fmt.Print(seat + " ")
			} else {
				fmt.Print("X ")
			}
		}
		fmt.Println()

		var selectedSeats []string
		fmt.Print("Enter seats (comma-separated): ")
		var seatsInput string
		fmt.Scanln(&seatsInput)
		selectedSeats = parseSeats(seatsInput)

		if !areSeatsAvailable(showNo, selectedSeats) {
			fmt.Println("Some of the seats are not available. Please select different seats.")
			continue
		}

		// Book the selected seats
		bookSeats(showNo, selectedSeats)

		// Calculate cost
		subtotal := calculateSubtotal(showNo, selectedSeats)
		serviceTax := subtotal * ServiceTax
		swachhCess := subtotal * SwachhCess
		krishiCess := subtotal * KrishiCess
		totalCost := subtotal + serviceTax + swachhCess + krishiCess

		// Print booking details
		fmt.Println("Successfully Booked - Show", showNo)
		fmt.Printf("Subtotal: Rs. %.2f\n", subtotal)
		fmt.Printf("Service Tax @14%%: Rs. %.2f\n", serviceTax)
		fmt.Printf("Swachh Bharat Cess @0.5%%: Rs. %.2f\n", swachhCess)
		fmt.Printf("Krishi Kalyan Cess @0.5%%: Rs. %.2f\n", krishiCess)
		fmt.Printf("Total: Rs. %.2f\n", totalCost)

		// Update total revenue and taxes
		totalRevenue += subtotal
		totalServiceTax += serviceTax
		totalSwachhCess += swachhCess
		totalKrishiCess += krishiCess

		fmt.Print("Would you like to continue (Yes/No): ")
		var continueBooking string
		fmt.Scanln(&continueBooking)
		if continueBooking != "Yes" {
			break
		}
	}

	// Print total sales
	fmt.Println("Total Sales:")
	fmt.Printf("Revenue: Rs. %.2f\n", totalRevenue)
	fmt.Printf("Service Tax: Rs. %.2f\n", totalServiceTax)
	fmt.Printf("Swachh Bharat Cess: Rs. %.2f\n", totalSwachhCess)
	fmt.Printf("Krishi Kalyan Cess: Rs. %.2f\n", totalKrishiCess)
}
