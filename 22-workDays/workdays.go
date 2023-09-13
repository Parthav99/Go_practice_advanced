//Same as (21) to accepts list of holidays and two dates, then calculates work days between two dates.
package main

import (
	"assignmentDependencies/inputs"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	holidayList := make(map[time.Time]bool)
	inputReader := bufio.NewReader(os.Stdin)

	//fetching userInput(holidays) into a map(holidayList) using a loop
	for {
		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
		holidayStr := inputs.FetchInput()

		if holidayStr == "stop" {
			inputDate1, inputDate2 := inputs.FetchAndValidateDates(inputReader) //fetches and parses the dates
			weekDays := countWorkday(inputDate1, inputDate2, holidayList)       //counts number of work days
			fmt.Println("Number of working days between two dates is", weekDays)
			break
		}

		holidayDate, err := time.Parse("2006-01-02", holidayStr)
		if err != nil {
			fmt.Println("Invalid date entered. Please follow the format")
			continue
		}
		holidayList[holidayDate] = true
	}
}

func countWorkday(inputDate1 time.Time, inputDate2 time.Time, holidayList map[time.Time]bool) int {

	countWeekDays, holidayCount := 0, 0
	hoursBetweenDates := inputDate2.Sub(inputDate1)
	totalDays := int(hoursBetweenDates.Hours() / 24)

	currentDate := inputDate1

	// if dates equal
	if inputDate1 == inputDate2 {
		if inputDate1.Weekday() != time.Saturday && inputDate1.Weekday() != time.Sunday &&
			!holidayList[inputDate1] {
			countWeekDays++
		}
		return countWeekDays
	}

	if holidayList[currentDate] || currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday {
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	for holidays := range holidayList {
		if inputDate1.Before(holidays) && inputDate2.After(holidays) || (holidays == inputDate1 || holidays == inputDate2) {
			if holidays.Weekday() != time.Saturday && holidays.Weekday() != time.Sunday {
				holidayCount++
			}
		}
	}

	weekendDays := (totalDays / 7) * 2
	remainingDays := totalDays % 7

	if remainingDays > 0 {
		for i := 0; i < remainingDays; i++ {
			if currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday {
				weekendDays++
			}
			currentDate = currentDate.AddDate(0, 0, 1)
		}
	}

	if inputDate2.Weekday() == time.Sunday {
		totalDays--
	}

	if inputDate2.Weekday() == time.Saturday {
		totalDays++
	}

	totalWeekDays := totalDays - (weekendDays + holidayCount)
	fmt.Println(totalDays)
	fmt.Println(weekendDays)
	fmt.Println(holidayCount)

	// adjust for weekends
	if totalWeekDays < 0 {
		totalWeekDays++
	}

	return totalWeekDays
}

/************************Logic Used**************************************/
//hoursBetweenDates := inputDate2.Sub(inputDate1) ---
// totalDays := int(hoursBetweenDates.Hours() / 24) ---
//if first date and second date
// holidays++ ----
//weekendDays := (totalDays / 7) * 2 ---
//remaining := totalDays % 7 ---
//check weekends in those remaining days
//loop over the holidayList as long as it is a holiday
//holiday++
//workDays := totalDays - (holidays + weekendDays)

/************************Normal Approach**************************************/

// //Same as (21) to accepts list of holidays and two dates, then calculates work days between two dates.
// package main

// import (
// 	"assignmentDependencies/inputs"
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"time"

// 	"golang.org/x/exp/slices"
// )

// func main() {

// 	holidayList := make([]time.Time, 0)
// 	inputReader := bufio.NewReader(os.Stdin)

// 	//fetching userInput(holidays) into a map(holidayList) using a loop
// 	for {
// 		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
// 		holidayStr := inputs.FetchInput()

// 		if holidayStr == "stop" {
// 			inputDate1, inputDate2 := inputs.FetchAndValidateDates(inputReader) //fetches and parses the dates
// 			weekDays := countWorkday(inputDate1, inputDate2, holidayList)       //counts number of work days
// 			fmt.Println("Number of working days between two dates is", weekDays)
// 			break
// 		}

// 		holidayDate, err := time.Parse("2006-01-02", holidayStr)
// 		if err != nil {
// 			fmt.Println("Invalid date entered. Please follow the format")
// 			continue
// 		}

// 		holidayList = append(holidayList, holidayDate)
// 	}
// }

// // Calculates number of workdays between two dates
// func countWorkday(inputDate1 time.Time, inputDate2 time.Time, holidayList []time.Time) int {
// 	countWeekDays := 0
// 	for currDate := inputDate1; currDate.Before(inputDate2) || currDate.Equal(inputDate2); currDate = currDate.AddDate(0, 0, 1) {
// 		if currDate.Weekday() != time.Saturday && currDate.Weekday() != time.Sunday &&
// 			!slices.Contains(holidayList, currDate) {
// 			countWeekDays++
// 		}
// 	}
// 	return countWeekDays
// }

//considering the first date
// if inputDate1.Weekday()!= 0 && inputDate1.Weekday()!= 6 && holidayList[inputDate1]{
// 	totalDays++
// }
