//Compares 2 Dates on the basis of whether they are equal, or which date appears earlier than the other

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

	/*------------- Comparing Date1 and Date2 -------------*/
	difference := date1.Sub(date2)

	if date1 == date2 {
		fmt.Println("Both dates are equal")
	}

	if difference < 0 {
		fmt.Println("Date 1 is earlier than Date 2")
	} else {
		fmt.Println("Date 2 is earlier than Date 1")
	}

}

// Fetches Input, outputs a string
func fetchInput() string {
	readInput := bufio.NewReader(os.Stdin)
	line, _ := readInput.ReadString('\n')
	line = line[:len(line)-1]
	return line
}
