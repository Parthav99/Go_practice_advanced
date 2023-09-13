// Extend (26) to accept URL as a command line argument instead of a hardcoded URL within
// the program.

package main

import (
	"compress/gzip"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"
)

func main() {

	start := time.Now()
	//Zero argument panic
	defer func() { //Immediately Invoked Function Expression
		if r := recover(); r != nil {
			var ok bool
			_, ok = r.(error)
			if ok {
				fmt.Println("No arguments entered, please enter a valid url.")
			}
		}
	}()

	//Fetching a url
	fileUrl := os.Args[1]
	if len(os.Args) != 2 {
		handleError(errors.New("incorrect number of arguments provided"))
	}

	//Skipping ssl verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	//Validating url
	parsedUrl, err := url.Parse(fileUrl)
	handleError(err)

	//creating context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) //what are contexts?
	defer cancel()

	//creating request with context
	request, err := http.NewRequest("GET", parsedUrl.String(), nil)
	handleError(err)
	request = request.WithContext(ctx)

	//bind to client to get response
	response, err := client.Do(request)
	handleError(err)
	defer response.Body.Close()

	//time taken to process the get request
	elapsedTime := time.Since(start).Seconds()
	fmt.Printf("%s took %v seconds \n", parsedUrl, elapsedTime)

	//Reading html content from url
	urlHtmlContent, err := ioutil.ReadAll(response.Body)
	handleError(err)

	//Creating zip file
	zipFile, err := os.Create("./merce-homepage.html" + ".zip")
	handleError(err)
	defer zipFile.Close()

	//Creating zip writer
	gzipWrite := gzip.NewWriter(zipFile)
	gzipWrite.Write(urlHtmlContent)
	fmt.Println("File successfully compressed")
	gzipWrite.Close()

	zipInfo, err := os.Stat(zipFile.Name())
	handleError(err)

	//Printing size of files in bytes
	fmt.Println("Size of original file in bytes:", len(urlHtmlContent))
	fmt.Println("Size of compressed file in bytes:", zipInfo.Size())
}

//Handles errors
func handleError(err error) {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	_, _, line, _ := runtime.Caller(1)
	if err != nil {
		log.Fatalf("\nError: %v\nError on line: %d", err, line)
	}
}
