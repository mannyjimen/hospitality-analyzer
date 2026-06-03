package reviewfilter

import (
	"bufio"
	"os"
)

type ReviewStreamer struct {
	file    *os.File
	scanner *bufio.Scanner
}
