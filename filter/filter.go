package filter

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
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

func GetUnfairBusinesses(configDir, yelpDir string) []Business {
	preprocess(configDir, yelpDir)
	streamer := getReviewStreamer(filepath.Join(yelpDir, "yelp_academic_dataset_review.json"))
	ids := getUnfairBusinesses(streamer)

	return ids
}

func getReviewStreamer(filePath string) *ReviewStreamer {
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
func getUnfairBusinesses(streamer *ReviewStreamer) []Business {
	defer helper.TrackTime(time.Now(), "Business Extraction")

	var unfairBusinessIDs = make(map[string]struct{})
	unfairBusinesses := []Business{}

	unfairReviewCount := 0

	for streamer.scanner.Scan() && unfairReviewCount < 5000 {

		review, err := streamer.getReview()
		if err != nil {
			continue
		}

		//business already deemed unfair and stored
		if _, ok := unfairBusinessIDs[review.Business_id]; ok {
			continue
		}

		if isTargetCityBusiness(review.Business_id) && isUnfairReview(review.Text) {
			unfairBusinessIDs[review.Business_id] = struct{}{}

			b := businesses[review.Business_id]

			//filling in Categories
			if b.RawCategories != "" {
				b.Categories = strings.Split(b.RawCategories, ", ")
			}

			unfairBusinesses = append(unfairBusinesses, b)
			unfairReviewCount++
		}
	}

	return unfairBusinesses
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
func isTargetCityBusiness(business_id string) bool {
	_, ok := businesses[business_id]
	return ok
}

func isTargetCity(city string) bool {
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

// // ARCHIVED
// func convMapToSlice(m map[string]struct{}) []string {
// 	s := []string{}
// 	for str := range m {
// 		s = append(s, str)
// 	}
// 	return s
// }
