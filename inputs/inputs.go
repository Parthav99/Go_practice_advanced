package inputs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//Fetches userInput
func FetchInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	line, _ := inputReader.ReadString('\n')
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		log.Fatal("No input entered")
	}
	return line
}

//Validates dates
func FetchAndValidateDates(inputReader *bufio.Reader) (time.Time, time.Time) {
	layoutFormat := "2006-01-02"
	//loop until valid dates are entered
	for {
		fmt.Print("Enter first date(YYYY-MM-DD): ")
		inputDateStr1 := FetchInput()
		fmt.Print("Enter second date(YYYY-MM-DD): ")
		inputDateStr2 := FetchInput()

		inputDate1, err1 := time.Parse(layoutFormat, inputDateStr1)
		inputDate2, err2 := time.Parse(layoutFormat, inputDateStr2)

		if err1 != nil || err2 != nil || inputDate1.After(inputDate2) {
			fmt.Println("Invalid dates entered. Please follow the format and ensure the first date is before the second.")
			continue
		}

		return inputDate1, inputDate2
	}
}

func FetchAndValidateInput(inputReader *bufio.Reader) (time.Time, int) {
	fmt.Print("Enter Date(YYYY-MM-DD): ")
	inputDateStr1 := FetchInput()
	fmt.Print("Enter business days: ")
	daysStr := FetchInput()

	inputDate, errDate1 := time.Parse("2006-01-02", inputDateStr1)
	businessDays, errDays := strconv.Atoi(daysStr)

	if errDate1 != nil || errDays != nil {
		log.Fatal("Invalid input. Please follow the format")
	}

	return inputDate, businessDays
}
