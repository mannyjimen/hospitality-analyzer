package main

import "github.com/mannyjimen/hospitality-analyzer/reviewfilter"

func main() {
	business_ids := reviewfilter.GetNegativeBusinessIDs()
	_ = business_ids
}
