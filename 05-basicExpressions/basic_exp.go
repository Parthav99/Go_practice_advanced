// Takes 2 whole numbers as input and calculates their sum

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*------------- Takes first input, checks if input is a whole number -------------*/
	fmt.Print("Enter first input: ")
	readInput_1 := bufio.NewReader(os.Stdin)
	line, _ := readInput_1.ReadString('\n')
	line = line[:len(line)-1]

	num1, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println(errors.New(line + " Is not a whole number"))
		os.Exit(1)
	}

	/*------------- Takes second input, checks if input is a whole number -------------*/
	fmt.Print("Enter second input: ")
	readInput_2 := bufio.NewReader(os.Stdin)
	line2, _ := readInput_2.ReadString('\n')
	line2 = line2[:len(line2)-1]

	num2, err := strconv.ParseInt(line2, 10, 64)
	if err != nil {
		fmt.Println(errors.New(line2 + " Is not a whole number"))
		os.Exit(1)
	}

	/*------------- Calculates and prints the summation -------------*/
	fmt.Printf("num1=%d num2=%d sum=%d", num1, num2, num1+num2)
}
