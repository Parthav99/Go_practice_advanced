//Implemented Sort operation using in-build function to sort elements of an array
//Extend (8)
package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	/*------------- Initializations -------------*/
	var input string
	var arrayOfNumbers []float64
	var sum float64
	/*------------- A loop which appends input to an array -------------*/
	for {
		fmt.Print("Type proceed to exit or Enter number to be added: ")
		fmt.Scan(&input)
		numberInput, err := strconv.ParseFloat(input, 64)

		/*------------- Break condition -------------*/
		if input == "proceed" {
			break
		}

		/*------------- Checks if the input is a valid number before appending to an array -------------*/
		if err != nil {
			fmt.Println(errors.New(input + " is not a valid number"))
		} else {
			arrayOfNumbers = append(arrayOfNumbers, numberInput)
		}
		fmt.Println("Array before sorting", arrayOfNumbers)

		/*------------- sort in-built function (Ascending Order) -------------*/
		sort.SliceStable(arrayOfNumbers, func(i, j int) bool {
			return arrayOfNumbers[i] < arrayOfNumbers[j]
		})
		sum += numberInput
		fmt.Println("Array after sorting", arrayOfNumbers)
	}
	fmt.Printf("Summation is %.2f\n", sum)
}
