package main

import (
	"github.com/mannyjimen/hospitality-analyzer/reviewfilter"
)

func main() {
	business_ids := reviewfilter.GetUnfairBusinessIDs()
	_ = business_ids

	// fmt.Println(len(business_ids))

	// for _, id := range business_ids {
	// 	fmt.Println(id)
	// }
}
