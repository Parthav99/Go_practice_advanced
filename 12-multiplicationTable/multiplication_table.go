// Multiplication table with two positive whole numbers using for loops.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*------------- Initializations -------------*/
	num1 := fetchInput()
	num2 := fetchInput()

	/*------------- Looping over num2 to create multiplication table -------------*/
	if num1 != 0 && num2 != 0 {
		for i := 1; i <= int(num2); i++ {
			fmt.Printf("%d * %d = %d\n", num1, i, int(num1)*i)
		}
	} else {
		fmt.Println("Multiplication with zero, results in zero")
	}

}

//Fetches the input and validates it to be a positive whole number
func fetchInput() int64 {
	var num int64
	fmt.Print("Enter input: ")
	readInput_1 := bufio.NewReader(os.Stdin)
	line, _ := readInput_1.ReadString('\n')
	line = line[:len(line)-1]

	num, err := strconv.ParseInt(line, 10, 64)
	if err != nil || num < 0 {
		fmt.Println(line + " is not a positive whole number")
		os.Exit(1)
	}
	return num
}
