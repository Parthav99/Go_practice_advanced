//Input using command line argument
package main

import (
	"fmt"
	"os"
)

func main() {

	/*------------- Checks if an argument is provided -------------*/
	if len(os.Args) < 2 {
		return
	}

	/*------------- Intialize name as argument input and print it with a greeting -------------*/
	name := os.Args[1]
	fmt.Println("Hello", name)
}
