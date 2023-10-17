### Movie Theater Seating Management System

Three shows are running in a movie theater, offering seats categorized as Platinum, Gold, and Silver. A, B, and C rows correspond to Platinum, Gold, and Silver categories, with prices set at Rs. 320, Rs. 280, and Rs. 240, respectively. Customers are shown available seats before making selections. Bookings are subject to a 14% service tax, along with Swachh Bharat Cess and Krishi Kalyan Cess, each at a 0.5% rate.

#### Seating Arrangement:

**Show 1 Running in Audi 1:**
- All Seats: A1 A2 A3 A4 A5 A6 A7 A8 A9, B1 B2 B3 B4 B5 B6, C2 C3 C4 C5 C6 C7

**Show 2 Running in Audi 2:**
- All Seats: A1 A2 A3 A4 A5 A6 A7, B2 B3 B4 B5 B6, C1 C2 C3 C4 C5 C6 C7

**Show 3 Running in Audi 3:**
- All Seats: A1 A2 A3 A4 A5 A6 A7, B1 B2 B3 B4 B5 B6 B7 B8, C1 C2 C3 C4 C5 C6 C7 C8 C9

---
---

#### Booking Process:

**Input:**
- Enter Show no: 1
- Available Seats: A1 A2 A3 A4 A5 A6 A7 A8 A9, B1 B2 B3 B4 B5 B6, C2 C3 C4 C5 C6 C7
- Enter seats: B1, B4

**Output:**
- Print: Successfully Booked - Show 1
- Subtotal: Rs. 560
- Service Tax @14%: Rs. 78.40
- Swachh Bharat Cess @0.5%: Rs. 2.80
- Krishi Kalyan Cess @0.5%: Rs. 2.80
- Total: Rs. 644
- Would you like to continue (Yes/No): Yes

**Input:**
- Enter Show no: 1
- Available Seats: A1 A2 A3 A4 A5 A6 A7 A8 A9, B2 B3 B5 B6, C2 C3 C4 C5 C6 C7
- Enter seats: B1, B3

**Output:**
- Print: B1 Not available, Please select different seats
- Would you like to continue (Yes/No): Yes

**Input:**
- Enter Show no: 2
- Available seats: A1 A2 A3 A4 A5 A6 A7, B2 B3 B4 B5 B6, C1 C2 C3 C4 C5 C6 C7
- Enter seats: A1, A2, A3

**Output:**
- Print: Successfully Booked - Show 2
- Subtotal: Rs. 960
- Service Tax @14%: Rs. 134.40
- Swachh Bharat Cess @0.5%: Rs. 4.80
- Krishi Kalyan Cess @0.5%: Rs. 4.80
- Total: Rs. 1104
- Would you like to continue (Yes/No): No

---

**Total Sales:**
- Revenue: Rs. 1520
- Service Tax: Rs. 212.80
- Swachh Bharat Cess: Rs. 7.60
- Krishi Kalyan Cess: Rs. 7.60
