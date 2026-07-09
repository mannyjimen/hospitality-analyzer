package reviewfilter

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mannyjimen/hospitality-analyzer/helper"
)

type Review struct {
	Business_id string          `json:"business_id"`
	Text        json.RawMessage `json:"text"`
}

type ReviewStreamer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func GetUnfairBusinessIDs() []string {
	preprocess()
	streamer := getReviewStreamer("YelpJSON/yelp_academic_dataset_review.json")
	ids := getUnfairBusinessIDs(streamer)

	return ids
}

func getReviewStreamer(filePath string) *ReviewStreamer {
	defer helper.TrackTime(time.Now(), "getReviewStreamer")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	streamer := ReviewStreamer{
		file:    file,
		scanner: scanner}

	return &streamer
}

/*
Go through every review and determine whether the associated business lies inside
the selected cities, and if so, determine whether the review is unfair.
Return a list of all business_ids that pass these checks.
*/
func getUnfairBusinessIDs(streamer *ReviewStreamer) []string {
	defer helper.TrackTime(time.Now(), "getUnfairBusinessIDs")

	var unfairBusinessIDs = make(map[string]struct{})
	for streamer.scanner.Scan() {
		review, err := streamer.getReview()
		if err != nil {
			continue
		}

		if isSelectedBusiness(review.Business_id) && isUnfairReview(review.Text) {
			unfairBusinessIDs[review.Business_id] = struct{}{}
		}
	}

	return convMapToSlice(unfairBusinessIDs)
}

func (r *ReviewStreamer) getReview() (Review, error) {
	var review Review

	err := json.Unmarshal(r.scanner.Bytes(), &review)

	if err != nil {
		return review, err
	}

	return review, nil
}

// returns whether the business_id is in businesses map (of chosen cities)
func isSelectedBusiness(business_id string) bool {
	_, ok := businesses[business_id]
	return ok
}

func isSelectedCity(city string) bool {
	city = strings.ToLower(city)
	_, ok := cities[city]
	return ok
}

// returns whether a negative keyword in the review text
func isUnfairReview(rawText json.RawMessage) bool {
	reviewText := string(rawText)

	for keyword := range keywords {
		if strings.Contains(reviewText, keyword) {
			return true
		}
	}
	return false
}

// helper functions
func convMapToSlice(m map[string]struct{}) []string {
	defer helper.TrackTime(time.Now(), "convMapToSlice")

	s := []string{}
	for str := range m {
		s = append(s, str)
	}
	return s
}
