# movie-booking

The Movie Booking System is a simple Go application that allows customers to book seats for different movie shows in a theater. The system displays available seats for each show categorized as Platinum, Gold, and Silver. Customers can select seats, view the total cost including taxes, and complete their bookings.

## Features

* Displays available seats for three different shows.
* Supports booking of seats in Platinum, Gold, and Silver categories.
* Calculates total cost including service tax, Swachh Bharat Cess, and Krishi Kalyan Cess.
* Provides a user-friendly interface for customers to select and book seats.
* Generates a summary of total revenue and taxes collected by the theater.


## Prerequisites
* Go installed on your system. You can download it from https://golang.org/dl/.

## How to Run
* Clone this repository to your local machine:
```
git clone https://github.com/akshay-singla/movie-booking.git
```

* Navigate to the project directory:
```
cd movie-booking
```

* Run the Go application:
```
go run main.go
```

* Follow the on-screen instructions to select the show and book seats.

## Usage
* Enter the show number (1, 2, or 3) to view available seats.
* Seats are displayed as follows:
    * Platinum seats: A1 to A9
    * Gold seats: B1 to B6
    * Silver seats: C1 to C9
* Enter the desired seats (comma-separated) to book them.
* The system calculates the subtotal, service tax, Swachh Bharat Cess, Krishi Kalyan Cess, and total cost for the booking.
* You can continue booking seats for multiple shows until you choose to stop.
* The system will display the total revenue, service tax, Swachh Bharat Cess, and Krishi Kalyan Cess collected after all bookings.

## Unit Tests
The project includes unit tests for various functions to ensure the correctness of the logic. You can run the tests using the following command:
```
go test ./... 
```