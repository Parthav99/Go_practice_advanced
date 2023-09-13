package main

import (
	"assignmentDependencies/inputs"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	holidayMap := make(map[time.Time]bool)
	inputReader := bufio.NewReader(os.Stdin)

	// inputs are to be stored inside a holidayMap for better lookup
	for {
		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
		holidayStr := inputs.FetchInput()

		if holidayStr == "stop" {
			inputDate, businessDays := inputs.FetchAndValidateInput(inputReader)
			relativeDate := calculateEndDate(inputDate, businessDays, holidayMap)
			fmt.Println("Relative Date is", relativeDate.Format("2006-01-02"))
			break
		}

		holidayDate, err := time.Parse("2006-01-02", holidayStr)
		if err != nil {
			fmt.Println("Invalid date entered. Please follow the format")
			continue
		}

		holidayMap[holidayDate] = true
	}
}

// excludes holidays and weekends in order to calculate an appropriate end date
func calculateEndDate(inputDate time.Time, businessDays int, holidayMap map[time.Time]bool) time.Time {

	currentDate := inputDate
	businessDaysEntered := businessDays

	addDay, parseDirection := 1, 1
	holidayCount, counter := 0, 0

	if businessDays == 0 {
		fmt.Println("No BusinessDays were entered")
		return inputDate
	}

	if businessDaysEntered < 0 {
		businessDays = -businessDays
		parseDirection = -1
	}

	weeks := businessDays / 5
	remainingDays := businessDays % 5
	weeksWithHolidays := (len(holidayMap) + weeks) / 5

	if (businessDays > 0 && weeks != 0) || (businessDays < 0 && remainingDays == 0) {
		for weeks <= weeksWithHolidays {
			counter++
			weeks = weeks + 1
		}
	}
	weekendDaysPerWeek := weeks * 2

	//iterates over remaining days and increments the weekend days
	if remainingDays > 0 {
		for i := 0; i < remainingDays; i++ {
			counter++
			if (currentDate.Weekday() == time.Saturday ||
				currentDate.Weekday() == time.Sunday) && currentDate != inputDate {
				weekendDaysPerWeek++
			}
			currentDate = currentDate.AddDate(0, 0, parseDirection)
		}
	}

	var checkHolidayTill time.Time
	addDays := weekendDaysPerWeek + businessDays + len(holidayMap)

	if businessDaysEntered < 0 {
		checkHolidayTill = inputDate.AddDate(0, 0, -addDays)
	} else {
		checkHolidayTill = inputDate.AddDate(0, 0, addDays)
	}

	// holidays are accounted here, only valid ones are considered
	for holiday := range holidayMap {
		counter++
		if (inputDate.Before(holiday) && checkHolidayTill.After(holiday)) ||
			holiday == checkHolidayTill {
			if holiday.Weekday() != time.Saturday && holiday.Weekday() != time.Sunday {
				holidayCount++
			}
		} else if holiday.Before(inputDate) && holiday.After(checkHolidayTill) {
			if holiday.Weekday() != time.Saturday && holiday.Weekday() != time.Sunday {
				holidayCount++
			}
		}
	}

	fmt.Println("WeekendDays", weekendDaysPerWeek)
	fmt.Println("Holidays", holidayCount)

	// to find the end date, holidays and weekends have to be accounted for
	totalHolidays := weekendDaysPerWeek + holidayCount
	totalDaysToAdd := businessDays + totalHolidays

	// if negative businessDays are entered, the days have to be decremented
	if businessDaysEntered < 0 {
		totalDaysToAdd = -totalDaysToAdd
		addDay = -1
	}

	// end date can be a holiday, hence has to be checked
	endDate := inputDate.AddDate(0, 0, totalDaysToAdd)
	notSunday := endDate

	for holidayMap[endDate] ||
		endDate.Weekday() == time.Saturday ||
		endDate.Weekday() == time.Sunday {
		counter++
		if endDate.Weekday() == time.Sunday && notSunday.Weekday() != time.Saturday && inputDate.Weekday() != time.Saturday { //condition -> Sunday, but if it is saturday then it comes to sunday and increments
			endDate = endDate.AddDate(0, 0, addDay)
		}
		endDate = endDate.AddDate(0, 0, addDay)
	}

	fmt.Println("Iterations", counter)
	return endDate
}

//-------------------------------------------------------------------------------//
// // With slices
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

// 	holidayMap := make([]time.Time, 0)
// 	inputReader := bufio.NewReader(os.Stdin)

// 	//feeds input into the map
// 	for {
// 		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
// 		holidayStr := inputs.FetchInput()

// 		//break condition
// 		if holidayStr == "stop" {
// 			inputDate, businessDays := inputs.FetchAndValidateInput(inputReader)
// 			relativeDate := calculateEndDate(inputDate, businessDays, holidayMap)
// 			fmt.Println("Relative Date is", relativeDate.Format("2006-01-02"))
// 			break
// 		}

// 		holidayDate, err := time.Parse("2006-01-02", holidayStr)
// 		if err != nil {
// 			fmt.Println("Invalid date entered. Please follow the format")
// 			continue
// 		}

// 		holidayMap = append(holidayMap, holidayDate)
// 	}
// }

// func calculateEndDate(inputDate time.Time, businessDays int64, holidayMap []time.Time) time.Time {

// 	//initializations
// 	currentDate := inputDate
// 	businessDaysEntered := businessDays
// 	addDay := 1
// 	checkDay := 1
// 	holiday := 0
// 	counter := 0 //track iterations

// 	//Check if businessDays are entered
// 	if businessDays == 0 {
// 		fmt.Println("No BusinessDays were entered")
// 		return inputDate
// 	}

// 	//check if negative business days are entered
// 	if businessDaysEntered < 0 {
// 		businessDays = -businessDays
// 		checkDay = -1
// 	}

// 	weeks := businessDays / 5         //converts business days into business weeks
// 	remainingDays := businessDays % 5 //days which remain after converting to weeks
// 	weeksWithHolidays := (len(holidayMap) + int(weeks)) / 5

// 	if (businessDays > 0 && weeks != 0) || (businessDays < 0 && remainingDays == 0) { //check?
// 		for int(weeks) <= weeksWithHolidays {
// 			counter++
// 			weeks = weeks + 1 //do this conditionally
// 		}
// 	}

// 	weekendDaysPerWeek := weeks * 2 // weekends per week
// 	//iterates over remaining days and increments the weekend days
// 	if remainingDays > 0 {
// 		for i := 0; i < int(remainingDays); i++ {
// 			counter++
// 			if (currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday) && currentDate != inputDate {
// 				weekendDaysPerWeek++
// 			}
// 			currentDate = currentDate.AddDate(0, 0, checkDay)
// 		}
// 	}

// 	var checkEndDate time.Time
// 	addDays := weekendDaysPerWeek + businessDays + int64(len(holidayMap))
// 	if businessDaysEntered < 0 {
// 		checkEndDate = inputDate.AddDate(0, 0, -int(addDays))
// 	} else {
// 		checkEndDate = inputDate.AddDate(0, 0, int(addDays))
// 	}

// 	for _, holidays := range holidayMap {
// 		counter++
// 		if (inputDate.Before(holidays) && checkEndDate.After(holidays)) || holidays == checkEndDate {
// 			if holidays.Weekday() != time.Saturday && holidays.Weekday() != time.Sunday {
// 				holiday++
// 			}
// 		} else if holidays.Before(inputDate) && holidays.After(checkEndDate) {
// 			if holidays.Weekday() != time.Saturday && holidays.Weekday() != time.Sunday {
// 				holiday++
// 			}
// 		}
// 	}

// 	fmt.Println("WeekendDays", weekendDaysPerWeek)
// 	fmt.Println("Holidays", holiday)

// 	//Calculating days to be added in order to get the end date
// 	totalHolidays := weekendDaysPerWeek + int64(holiday)
// 	totalDaysToAdd := businessDays + totalHolidays

// 	//if businessDays are negative, decrement days
// 	if businessDaysEntered < 0 {
// 		totalDaysToAdd = -totalDaysToAdd
// 		addDay = -1
// 	}

// 	//add days to the end date and iterate as long as the end date is a holiday
// 	endDate := inputDate.AddDate(0, 0, int(totalDaysToAdd))
// 	fmt.Println("Input date:", inputDate)
// 	fmt.Println("End date:", endDate)
// 	notSunday := endDate

// 	for slices.Contains(holidayMap, endDate) || endDate.Weekday() == time.Saturday || endDate.Weekday() == time.Sunday {
// 		counter++
// 		if endDate.Weekday() == time.Sunday && notSunday.Weekday() != time.Saturday && inputDate.Weekday() != time.Saturday { //condition -> Sunday, but if it is saturday then it comes to sunday and increments
// 			endDate = endDate.AddDate(0, 0, addDay)
// 		}
// 		endDate = endDate.AddDate(0, 0, addDay)
// 	}
// 	fmt.Println("Iterations", counter)
// 	return endDate
// }
