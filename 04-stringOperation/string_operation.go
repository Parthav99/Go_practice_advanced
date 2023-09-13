//Takes template and name as command line arguments.
//Replaces the name parameter in the template with the second argument
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*------------- Initializations -------------*/
	template := os.Args[1]
	name := os.Args[2]
	message, err := checkTemplate(template, name)

	if err == nil {
		fmt.Println(message)
	} else {
		fmt.Println(err)
	}
}

/*------------- Checks if template contains the name parameter. -------------*/
func checkTemplate(template, name string) (string, error) {

	if !strings.Contains(template, "name") {
		fmt.Println(errors.New(template + " Must contain name parameter"))
		os.Exit(1)
	}

	/*------------- Replaces instances of name in the template with the second argument -------------*/
	finalMessage := strings.Replace(template, "name", name, -1)
	return finalMessage, nil
}
