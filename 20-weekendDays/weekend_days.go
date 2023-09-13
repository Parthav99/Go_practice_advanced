//Accepts 2 dates and calculates the number of weekend days between them

package main

import (
	"assignmentDependencies/inputs"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	date1, date2 := inputs.FetchAndValidateDates(inputReader) //fetches and parses the dates

	/*------------- Counts the number of weekend days -------------*/
	countWeekendDays := 0
	for startDate := date1; startDate.Before(date2) || startDate.Equal(date2); startDate = startDate.AddDate(0, 0, 1) {
		if startDate.Weekday() == time.Saturday {
			if !startDate.Equal(date2) {
				countWeekendDays = countWeekendDays + 2
			} else {
				countWeekendDays = countWeekendDays + 1
			}
			startDate = startDate.AddDate(0, 0, 1)
		} else if startDate.Weekday() == time.Sunday {
			countWeekendDays = countWeekendDays + 1
		}
	}
	fmt.Printf("There are %d weekend days from the first date to the second\n", countWeekendDays)
}
