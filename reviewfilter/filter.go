package reviewfilter

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Review struct {
	Business_id string `json:"business_id"`
	Text        string `json:"text"`
}

type ReviewStreamer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func GetNegativeBusinessIDs() []string {

	preprocess()
	streamer := getReviewStreamer("YelpJSON/yelp_academic_dataset_review.json")
	ids := getNegativeBusinessIDs(streamer)

	return ids
}

func getReviewStreamer(filePath string) ReviewStreamer {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	streamer := ReviewStreamer{
		file:    file,
		scanner: scanner}

	return streamer
}

//go through every review
//first determine whether the review is of a business with a bussiness_id in businesses map
//then determine whether review text has any negative keyword
//if so, we need to store/remember the business_id of this "unfair treatement" review

func getNegativeBusinessIDs(streamer ReviewStreamer) []string {
	var unfairBusinessIDs = make(map[string]struct{})
	for streamer.scanner.Scan() {
		review := streamer.getNextReview()

		if isSelectedBusiness(review.Business_id) && isUnfairReview(review.Text) {
			unfairBusinessIDs[review.Business_id] = struct{}{}
		}
	}

	var idList = []string{}
	for id := range unfairBusinessIDs {
		idList = append(idList, id)
	}
	return idList
}

func (r ReviewStreamer) getNextReview() Review {
	var review Review

	err := json.Unmarshal([]byte(r.scanner.Text()), &review)

	if err != nil {
		log.Fatal(err)
	}

	return review
}

// returns whether the business_id is in businesses map (of chosen cities)
func isSelectedBusiness(business_id string) bool {
	_, ok := businesses[business_id]
	return ok
}

// returns whether a negative keyword in the review text
func isUnfairReview(review_text string) bool {
	for keyword := range keywords {
		if strings.Contains(review_text, keyword) {
			return true
		}
	}
	return false
}
