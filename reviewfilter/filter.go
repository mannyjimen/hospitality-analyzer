package reviewfilter

import (
	"bufio"
	"log"
	"os"
)

// map containing negative keywords, will be used for review filtering.
// All negative keywords are newline separated and are in './negativekeywords.txt'
var keywords = make(map[string]bool)

type ReviewStreamer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func GetNegativeBusinessIDs() []string {
	file, err := os.Open("negativekeywords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	streamer := ReviewStreamer{
		file:    file,
		scanner: scanner}

	return []string{}
}
