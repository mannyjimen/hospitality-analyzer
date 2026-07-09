package main

import (
	"fmt"

	"github.com/mannyjimen/hospitality-analyzer/reviewfilter"
)

func main() {
	business_ids := reviewfilter.GetUnfairBusinessIDs()
	fmt.Printf("We gathered %d unfair businesses", len(business_ids))
}
