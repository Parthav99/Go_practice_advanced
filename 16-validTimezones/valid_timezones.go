// Prints supported timezones and accepts valid timezone prints time as per the time zone selected
package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"time"
)

func main() {

	/*------------- Prompting user for timezone input, validating it and initializing it to dateTime variable -------------*/
	validTimezones := []string{"Asia/Calcutta", "Africa/Cairo", "America/Chicago"}
	var timezoneInput string
	var location *time.Location

	fmt.Print("Valid timezones: ")
	fmt.Println(validTimezones)
	fmt.Print("Enter valid timezone: ")
	fmt.Scanln(&timezoneInput)

	if slices.Contains(validTimezones, timezoneInput) {
		location, _ = time.LoadLocation(timezoneInput)
	} else {
		fmt.Println("Invalid timezone entered. Valid timezone examples: Asia/Calcutta, Africa/Cairo, America/Chicago, etc")
		os.Exit(1)
	}

	dateTime := time.Now().In(location)

	/*------------- Time as per timezone selected -------------*/
	switch timezoneInput {
	case "Asia/Calcutta":
		fmt.Println("Time as per timezone selected is", dateTime.Format("03:04:05"))
	case "Africa/Cairo":
		fmt.Println("Time as per timezone selected is", dateTime.Format("03:04:05"))
	case "America/Chicago":
		fmt.Println("Time as per timezone selected is", dateTime.Format("03:04:05"))
	default:
		fmt.Println("Invalid timezone selected.")
		os.Exit(1)
	}
}
