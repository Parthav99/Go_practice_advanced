// Prints the number of times each word occurs in a file. 

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	//accept filename input
	filename := os.Args[1]

	//open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read file
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	//Inserting clean words, from the file into the map
	wordMap := make(map[string]int)
	for fileScanner.Scan() {
		word := strings.ToLower(fileScanner.Text())
		cleanedWord := wordCleaner(word)

		if strings.Contains(cleanedWord, " ") {
			for _, cleanWord := range strings.Fields(cleanedWord) {
				wordMap[cleanWord]++
			}
		} else {
			wordMap[cleanedWord]++
		}

	}

	// Prints word occurrence
	for word, wordFreq := range wordMap {
		if len(word) < 2 {
			continue
		} else {
			fmt.Printf("%s: %d\n", word, wordFreq)
		}
	}
}

//Cleans
func wordCleaner(word string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	replacedString := reg.ReplaceAllString(word, " ")
	replacedString = strings.TrimSpace(replacedString)

	return replacedString
}

//why?
