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

//ARCHIVED
// type Review struct {
// 	Business_id string `json:"business_id"`
// 	Text        string `json:"text"`
// }

type Business_id struct {
	Business_id string `json:"business_id"`
}

type ReviewText struct {
	Text string `json:"text"`
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
		business_id, err := streamer.getBusinessID()
		if err != nil {
			continue
		}

		if isSelectedBusiness(business_id) {
			reviewText, err := streamer.getReviewText()
			if err != nil {
				continue
			}

			if isUnfairReview(reviewText) {
				unfairBusinessIDs[business_id] = struct{}{}
			}
		}
	}

	return convMapToSlice(unfairBusinessIDs)
}

//ARCHIVED
// func (r *ReviewStreamer) getNextReview() (Review, error) {
// 	var review Review

// 	err := json.Unmarshal(r.scanner.Bytes(), &review)

// 	if err != nil {
// 		return review, err
// 	}

// 	return review, nil
// }

func (r *ReviewStreamer) getBusinessID() (string, error) {
	var b Business_id

	err := json.Unmarshal(r.scanner.Bytes(), &b)

	if err != nil {
		return "", err
	}

	return b.Business_id, nil
}

func (r *ReviewStreamer) getReviewText() (string, error) {
	var rt ReviewText

	err := json.Unmarshal(r.scanner.Bytes(), &rt)

	if err != nil {
		return "", err
	}

	return rt.Text, nil
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
func isUnfairReview(reviewText string) bool {
	reviewText = strings.ToLower(reviewText)
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
