//Processing config/cities.txt, config/negative_keywords.txt, and yelp_academic_dataset_business.json

package reviewfilter

import (
	"bufio"
	"log"
	"os"
)

// map containing cities we want to use for research.
// All cities are newline separated in 'config/cities.txt'
var cities = make(map[string]struct{})

// map containing negative keywords, will be used for review filtering.
// All negative keywords are newline separated in 'config/negative_keywords.txt'
var keywords = make(map[string]struct{})

// map containing {business_id : city} pairs
// only businesses that reside in target cities will be stored.
var businesses = make(map[string]string)

func preprocess() {
	processCities()
	processKeywords()
	processBusinesses()
}

func processCities() {
	file, err := os.Open("config/cities.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		city := scanner.Text()
		cities[city] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processKeywords() {
	file, err := os.Open("config/negative_keywords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyword := scanner.Text()
		keywords[keyword] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processBusinesses() {}
