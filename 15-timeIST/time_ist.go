// Prints time in IST timezone

package main

import (
	"fmt"
	"time"
)

func main() {

	/*------------- Using IST by setting location to Calcutta -------------*/
	location, _ := time.LoadLocation("Asia/Calcutta")
	dateTime := time.Now().In(location)

	/*------------- Printing current date and time in different formats  -------------*/
	//Specified Format: 16 Mar 2022
	fmt.Println(dateTime.Format("2 January 2006"))

	//Specified Format: Mar 16, 2022
	fmt.Println(dateTime.Format("Jan 02, 2006"))

	//Specified Format: 2022-03-16
	fmt.Println(dateTime.Format("2006-01-02"))

	//Specified Format: 2022-03-16T15:52:00Z
	fmt.Println(dateTime.Format(time.RFC3339))

	//Specified Format: Tuesday, 16 March 2022
	fmt.Println(dateTime.Format("Monday, 02 Jan 2006"))
}
