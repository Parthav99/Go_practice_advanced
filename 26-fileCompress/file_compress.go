//Opens an html file compresses it, returns its orginal and compressed sizes.

package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func main() {

	inputPath := "./merce-homepage.html"

	//Opening html file
	inputFile, err := os.Open(inputPath)
	handleError(err)
	defer inputFile.Close()

	//Reading html file
	htmlcontent, err := ioutil.ReadAll(inputFile)
	handleError(err)

	//Creating zip file
	zipFile, err := os.Create(inputPath + ".zip")
	handleError(err)
	defer zipFile.Close()

	//Creating zip writer
	gzipWrite := gzip.NewWriter(zipFile)
	gzipWrite.Write(htmlcontent)
	fmt.Println("File successfully compressed")
	gzipWrite.Close()

	zipInfo, err := os.Stat(zipFile.Name())
	handleError(err)

	//opening zip file
	openZip,_:= os.Open(zipFile.Name())
	defer openZip.Close()
	fmt.Println(openZip)
	
	//Printing size of files in bytes
	fmt.Println("Size of original file in bytes:", len(htmlcontent))
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
