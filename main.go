package main

import (
	"fmt"
	"os"

	"github.com/mannyjimen/hospitality-analyzer/filter"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Provide program arguments: main.go <config_directory_path> <yelp_data_path>")
		return
	}

	configDir := os.Args[1]
	yelpDir := os.Args[2]

	businesses := filter.GetUnfairBusinesses(configDir, yelpDir)
	_ = businesses

	// fmt.Println(len(businesses))

	// for _, b := range businesses {
	// 	fmt.Println(b)
	// }
}
