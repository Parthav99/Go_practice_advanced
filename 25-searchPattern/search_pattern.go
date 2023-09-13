//Takes a list of names and a search pattern. Returns a list of names which match the seach pattern.
package main

import (
	"assignmentDependencies/inputs"
	"fmt"
	"log"
	"regexp"
)

func main() {
	nameList := make([]string, 0)

	// A loop which appends input to a list
	for {
		fmt.Print("Enter names of candidates or done: ")
		inputStr := inputs.FetchInput()

		// Break condition
		if inputStr == "done" {
			if len(nameList) == 0 {
				log.Fatal("List is empty, nothing to search.")
			}
			break
		}
		nameList = append(nameList, inputStr) //adds input to the list
	}

	// Prints the list of successful matches
	fmt.Print("Enter a pattern to be searched: ")
	searchInput := inputs.FetchInput()
	matchList := searchPattern(nameList, searchInput)
	fmt.Println(matchList)
}

//If the given pattern exists in the nameList elements, it appends that element into matchList
func searchPattern(nameList []string, searchInput string) []string {
	searchPattern := searchInput
	searchStr, _ := regexp.Compile("(?i)" + searchPattern) //parses or compiles the pattern into a regular expression

	var matchList []string
	for _, name := range nameList {
		if searchStr.MatchString(name) { //matches pattern with the name in the list
			matchList = append(matchList, name)
		}
	}

	if len(matchList) == 0 {
		fmt.Println("No matches found")
	}
	return matchList
}
