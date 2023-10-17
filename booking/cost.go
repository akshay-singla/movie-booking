package booking

const (
	PlatinumCost = 320
	GoldCost     = 280
	SilverCost   = 240
	ServiceTax   = 0.14
	SwachhCess   = 0.005
	KrishiCess   = 0.005
)

func calculateSubtotal(showNo int, selectedSeats []string) float64 {
	var subtotal float64
	for _, seat := range selectedSeats {
		seat = seat[:1]
		switch seat {
		case "A":
			subtotal += PlatinumCost
		case "B":
			subtotal += GoldCost
		case "C":
			subtotal += SilverCost
		}
	}
	return subtotal
}
