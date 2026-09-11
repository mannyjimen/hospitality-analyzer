package main

import (
	"fmt"
	"os"

	"github.com/mannyjimen/hospitality-analyzer/filter"
	"github.com/mannyjimen/hospitality-analyzer/similarity"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Provide program arguments: main.go <config_directory_path> <yelp_data_path>")
		return
	}

	configDir := os.Args[1]
	yelpDir := os.Args[2]

	businesses := filter.GetUnfairBusinesses(configDir, yelpDir)
	similarities := similarity.FindSimilarities(businesses)

	// fmt.Println(len(businesses))

	// for _, b := range businesses {
	// 	fmt.Println(b)
	// }

	for _, s := range similarities {
		fmt.Println(s)
	}
}
