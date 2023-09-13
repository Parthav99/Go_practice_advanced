// Accepts two dates, validates it and calculates the difference between them
//Add changes
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	/*------------- Fetches input, valids it to be of a given format -------------*/
	fmt.Print("Enter the first date(YYYY-MM-DD hh:mm:ss): ")
	dateStr1 := fetchInput()
	fmt.Print("Enter the second date(YYYY-MM-DD hh:mm:ss): ")
	dateStr2 := fetchInput()

	date1, err1 := time.Parse("2006-01-02 15:04:05", dateStr1)
	date2, err2 := time.Parse("2006-01-02 15:04:05", dateStr2)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid format entered.")
		os.Exit(1)
	}

	/*------------- Calulating the Difference -------------*/
	if date2.After(date1) {
		date1, date2 = date2, date1
	}
	yearDifference := date1.Year() - date2.Year()
	dayDifference := date1.Day() - date2.Day()
	minuteDifference := date1.Minute() - date2.Minute()

	hourDifference := date1.Hour() - date2.Hour()
	monthDifference := date1.Month() - date2.Month()

	if minuteDifference < 0 {
		minuteDifference += 60
		hourDifference--
	}
	if hourDifference < 0 {
		hourDifference += 24
		dayDifference--
	}
	if dayDifference < 0 {
		dayDifference += 31
		monthDifference--
	}
	if monthDifference < 0 {
		monthDifference += 12
		yearDifference--
	}
	fmt.Printf("Difference is of %d years %d months %d days %d minutes\n", yearDifference, monthDifference, dayDifference, minuteDifference)
}

// Fetches Input, outputs a string
func fetchInput() string {
	readInput := bufio.NewReader(os.Stdin)
	line, _ := readInput.ReadString('\n')
	line = line[:len(line)-1]
	return line
}
