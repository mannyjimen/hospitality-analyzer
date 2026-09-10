package main

import (
	"fmt"
	"os"

	"github.com/mannyjimen/hospitality-analyzer/reviewfilter"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Provide program arguments: main.go <config_directory_path> <yelp_data_path>")
		return
	}

	configDir := os.Args[1]
	yelpDir := os.Args[2]

	business_ids := reviewfilter.GetUnfairBusinessIDs(configDir, yelpDir)
	_ = business_ids

	fmt.Println(len(business_ids))

	for _, id := range business_ids {
		fmt.Println(id)
	}
}
