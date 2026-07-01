package main

import (
	"github.com/mannyjimen/hospitality-analyzer/reviewfilter"
)

func main() {
	business_ids := reviewfilter.GetNegativeBusinessIDs()
	_ = business_ids

	// for _, id := range business_ids {
	// 	fmt.Println(id)
	// }
}
