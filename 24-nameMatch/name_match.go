//Checks if name exists in a list of names stored in a Hash table
package main

import (
	"assignmentDependencies/inputs"
	"fmt"
	"os"
	"strings"
)

func main() {
	nameList := make(map[string]bool)
	//writes input into the map(nameList)
	for {
		fmt.Print("Enter names of candidates or done: ")
		inputStr := inputs.FetchInput()

		if inputStr == "done" {
			break
		}
		nameList[strings.ToLower(inputStr)] = true
	}

	if len(nameList) == 0 {
		fmt.Println("No names have been entered.")
		os.Exit(1)
	}

	//Checks if name exists in the map
	fmt.Print("Enter name to be searched: ")
	inputStr := inputs.FetchInput()

	if !nameList[strings.ToLower(inputStr)] {
		fmt.Println(inputStr, "does not exist")
	} else {
		fmt.Println(inputStr, "exists")
	}
}
