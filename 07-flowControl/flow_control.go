//Implemented switch cases to perform statistical operations.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*------------- Initializations -------------*/
	num1 := fetchInput()
	num2 := fetchInput()

	var performedOperation string
	var result int64
	var operation string

	fmt.Print("Enter operation to be performed: ")
	fmt.Scan(&operation)

	/*------------- Operations to be performed, based on user input. -------------*/
	switch operation {
	case "+":
		performedOperation = "sum"
		result = num1 + num2
	case "-":
		performedOperation = "difference"
		result = num1 - num2
	case "*":
		performedOperation = "multiply"
		result = num1 * num2
	case "/":
		performedOperation = "divide"
		if num2 == 0 {
			fmt.Println(errors.New(operation + " causes zero division error"))
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println(errors.New(operation + " is an invalid Operation"))
		os.Exit(1)

	}
	fmt.Printf("num1=%d num2=%d %s=%d\n", num1, num2, performedOperation, result)
}

// Fetches the input and validates it to be a whole number
func fetchInput() int64 {
	var num int64
	fmt.Print("Enter input: ")
	readInput_1 := bufio.NewReader(os.Stdin)
	line, _ := readInput_1.ReadString('\n')
	line = line[:len(line)-1]

	num, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println(errors.New(line + " Is not a whole number"))
		os.Exit(1)
	}
	return num
}
