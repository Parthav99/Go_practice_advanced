//Takes 2 whole numbers or floating point values as input and calculates their sum

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*------------- Takes first input, also accepts floating point numbers -------------*/
	fmt.Print("Enter first input: ")
	readInput_1 := bufio.NewReader(os.Stdin)
	line, _ := readInput_1.ReadString('\n')
	line = line[:len(line)-1]

	num1, err := strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Println(errors.New(line + " Is not a whole number"))
		os.Exit(1)
	}

	/*------------- Takes second input, also accepts floating point numbers -------------*/
	fmt.Print("Enter second input: ")
	readInput_2 := bufio.NewReader(os.Stdin)
	line2, _ := readInput_2.ReadString('\n')
	line2 = line2[:len(line2)-1]

	num2, err := strconv.ParseFloat(line2, 64)
	if err != nil {
		fmt.Println(errors.New(line2 + " Is not a whole number"))
		os.Exit(1)
	}

	/*------------- Calculates and prints the summation -------------*/
	fmt.Printf("num1=%0.2f num2=%0.2f sum=%.2f\n", num1, num2, num1+num2)
}
