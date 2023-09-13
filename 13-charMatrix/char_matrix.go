// Implemented a rectangular matrix using nested for loops

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*------------- Initializations -------------*/
	var num3 string
	fmt.Print("Enter character to be printed: ")
	fmt.Scanln(&num3)

	if len(num3) > 1 {
		fmt.Println("More than one characters found. Enter a single character.")
		os.Exit(1)
	}

	num1 := fetchInput()
	num2 := fetchInput()

	/*------------- Looping through num1 and num2 to create a matrix -------------*/
	for i := 0; i < int(num1); i++ {
		for j := 0; j < int(num2); j++ {
			fmt.Print(num3 + " ")
		}
		fmt.Println()
	}
}

//Fetches the input and validates it
func fetchInput() int64 {
	var num int64
	fmt.Print("Enter input: ")
	readInput_1 := bufio.NewReader(os.Stdin)
	line, _ := readInput_1.ReadString('\n')
	line = line[:len(line)-1]

	num, err := strconv.ParseInt(line, 10, 64)
	if err != nil || num <= 0 {
		fmt.Println(line + " is not a valid input for representing a matrix.")
		os.Exit(1)
	}
	return num
}
