// Checks whether the year given as an input is a leap year.
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

	date1, err1 := time.Parse("2006", dateStr1)
	if err1 != nil {
		fmt.Println("Invalid format entered.")
		os.Exit(1)
	}

	/*------------- Checks whether given year is a leap year -------------*/
	if (date1.Year()%4 == 0 && date1.Year()%100 != 0) || date1.Year()%400 == 0 {
		fmt.Printf("%v is a leap year.\n", date1.Year())
	} else {
		fmt.Printf("%v is not a leap year.\n", date1.Year())
	}

}

// Fetches Input, outputs a string
func fetchInput() string {
	readInput := bufio.NewReader(os.Stdin)
	line, _ := readInput.ReadString('\n')
	line = line[:len(line)-1]
	return line
}
