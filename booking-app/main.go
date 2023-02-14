package main

import "fmt"

func main() {
	bookingType := "Studio Time"
	const timeSlots = 10
	remainingSlots := 10

	fmt.Printf("Welcome to the %v booking application!\n", bookingType)
	fmt.Printf("We have a total of %v slots, and %v slots still available.\n", timeSlots, remainingSlots)
	fmt.Printf("Reserve your time slot here.\n ")

	var userName string
	var timeBlocks int
	// ask the user for their name

	userName = "Solange"
	timeBlocks = 3
	fmt.Printf("User %v booked %v timeslots.\n", userName, timeBlocks)

}
