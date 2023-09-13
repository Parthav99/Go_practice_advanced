//Accept list of holidays, and check if a day is working or not

package main

import (
	"assignmentDependencies/inputs"
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	holidayList := make(map[time.Time]bool)

	//fetching userInput(holidays) into a map(holidayList) using a loop
	for {
		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
		holidayStr := inputs.FetchInput()

		//break condition
		if holidayStr == "stop" {
			fmt.Print("Check whether this date is a working day(YYYY-MM-DD): ")
			inputDateStr := inputs.FetchInput()
			inputDate, err := time.Parse("2006-01-02", inputDateStr) //parsing time into a particular format
			handleError(err)

			fmt.Println(checkWorkingday(inputDate, holidayList))
			break
		}
		holidayDate, err := time.Parse("2006-01-02", holidayStr)
		handleError(err)

		holidayList[holidayDate] = true //creates entries of input dates inside the map
	}

}

//checks whether date entered is a working day
func checkWorkingday(inputDate time.Time, holidayList map[time.Time]bool) string {
	dayStatus := "The date provided is a working day"
	if inputDate.Weekday() == time.Saturday ||
		inputDate.Weekday() == time.Sunday ||
		holidayList[inputDate] {
		dayStatus = "The date provided is a holiday"
	}
	return dayStatus
}

//Handles errors
func handleError(err error) {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	_, _, line, _ := runtime.Caller(1)
	if err != nil {
		log.Fatalf("\nError: %v\nError on line: %d", err, line)
	}
}











//------------------------------------------------------------------------------------------------//

// //Accepts list of holidays from the user. With the help of this list and weekends,
// // checks whether user input date is a working data or not.
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"golang.org/x/exp/slices"
// 	"os"
// 	"time"
// )

// func main() {
// 	var holidayList []time.Time
// 	//A loop which appends date input to an array
// 	for {

// 		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
// 		holidayStr := fetchInput()
// 		holidayDate, err := time.Parse("2006-01-02", holidayStr)

// 		if err != nil && holidayStr != "stop" {
// 			fmt.Println("Invalid date entered. Please follow the format")
// 		} else if slices.Contains(holidayList, holidayDate) {
// 			fmt.Println(holidayDate.String() + " already exists in holidayList.")
// 			// Break condition
// 		} else if holidayStr == "stop" {
// 			fmt.Print("Check whether this date is a working day(YYYY-MM-DD): ")
// 			inputDateStr := fetchInput()
// 			inputDate, err2 := time.Parse("2006-01-02", inputDateStr)

// 			if err2 != nil {
// 				fmt.Println("Invalid date entered. Please follow the format")
// 				os.Exit(1)
// 			}
// 			//Outputs whether the given date is working or not
// 			fmt.Println(checkWorkingday(inputDate, holidayList))
// 			break
// 		} else {
// 			holidayList = append(holidayList, holidayDate)
// 		}
// 	}
// }

// //fetches input
// func fetchInput() string {
// 	inputReader := bufio.NewReader(os.Stdin)
// 	line, _ := inputReader.ReadString('\n')
// 	line = line[:len(line)-1]
// 	return line
// }

// //Checks whether the date is a working day or a holiday
// func checkWorkingday(inputDate time.Time, holidayList []time.Time) string {
// 	isHoliday := "The date provided is a working day"
// 	if inputDate.Weekday().String() == "Saturday" ||
// 		inputDate.Weekday().String() == "Sunday" ||
// 		slices.Contains(holidayList, inputDate) {
// 		isHoliday = "The date provided is a holiday"
// 	}
// 	return isHoliday
// }
