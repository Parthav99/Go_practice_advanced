//Prompts user for a valid input until "proceed" is entered as an input.
//Then it calculates the sum of all the numbers entered and produces an output

package sumNumbers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Sum() float64 {

	/*------------- Initializations -------------*/
	var input string
	var sum float64

	/*------------- Fetch input until break condition is reached -------------*/
	for {
		fmt.Print("Type proceed to exit or Enter number to be added: ")
		fmt.Scan(&input)
		numberInput, err := strconv.ParseFloat(input, 64)

		/*------------- Break condition -------------*/
		if input == "proceed" {
			break
		}

		if err != nil {
			fmt.Println(errors.New(input + " is not a valid number"))
			os.Exit(1)
		}

		/*------------- Calculates Summation -------------*/
		sum += numberInput
	}

	return sum
}
