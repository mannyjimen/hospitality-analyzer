package similarity

import (
	"container/heap"

	"github.com/mannyjimen/hospitality-analyzer/filter"
)

type CategoryFreq struct {
	category  string
	frequency int
}

type CategoryHeap []CategoryFreq

func (c *CategoryHeap) Push(cf any) {
	(*c) = append((*c), cf.(CategoryFreq))
}

func (c *CategoryHeap) Pop() any {
	last := (*c)[len(*c)-1]
	(*c) = (*c)[:len(*c)-1]

	return last
}

func (c *CategoryHeap) Len() int {
	return len(*c)
}

func (c *CategoryHeap) Less(i, j int) bool {
	return (*c)[i].frequency > (*c)[j].frequency
}

func (c *CategoryHeap) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func FindSimilarities(businesses []filter.Business) []CategoryFreq {
	categoryFreqs := make(map[string]int)

	for _, b := range businesses {
		for _, c := range b.Categories {
			categoryFreqs[c]++
		}
	}

	var h CategoryHeap

	for cat, freq := range categoryFreqs {
		cf := CategoryFreq{cat, freq}
		heap.Push(&h, cf)
	}

	frequentCategories := []CategoryFreq{}

	for len(h) != 0 && len(frequentCategories) < 10 {
		cat := heap.Pop(&h).(CategoryFreq)
		frequentCategories = append(frequentCategories, cat)
	}

	return frequentCategories
}
