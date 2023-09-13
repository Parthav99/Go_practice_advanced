// Prints the number of times each word occurs in a file. Also excludes common words

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

	// Prints word occurrence exluding common words
	commonWords := map[string]bool{
		"an": true, "the": true, "of": true, "if": true,
		"and": true, "for": true, "then": true, "to": true,
		"it": true, "or": true, "did": true, "are": true, "in": true,
		"this": true, "is": true, "what": true}

	for word, wordFreq := range wordMap {
		if len(word) < 2 || commonWords[word] {
			continue
		} else {
			fmt.Printf("%s: %d\n", word, wordFreq)
		}
	}
}


func wordCleaner(word string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	replacedString := reg.ReplaceAllString(word, " ")
	replacedString = strings.TrimSpace(replacedString)

	return replacedString
}

//why?
