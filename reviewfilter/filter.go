package reviewfilter

import (
	"bufio"
	"os"
)

type Review struct {
	business_id string
	text        string
}

type ReviewStreamer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func GetNegativeBusinessIDs() []string {

	preprocess()
	// scanner := bufio.NewScanner(file)
	// streamer := ReviewStreamer{
	// 	file:    file,
	// 	scanner: scanner}

	return []string{}
}
