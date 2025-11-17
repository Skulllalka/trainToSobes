package mymap

import (
	//"fmt"
	"github.com/k0kubun/pp"
)

type WordCount struct {
	order  []string
	limit  int
	counts map[string]int
}

func NewWordCounter(limit int) *WordCount {
	wc := WordCount{
		limit:  limit,
		counts: make(map[string]int),
	}
	return &wc
}

func (wc *WordCount) CountWord(word string) {
	if _, ok := wc.counts[word]; !ok{
		wc.order = append(wc.order, word)
	}
	wc.counts[word]++
	if len(wc.counts) > wc.limit {
		oldest := wc.order[0]
		wc.order = wc.order[1:]
		delete(wc.counts, oldest)
	}
}

func Task3() {
	wc := NewWordCounter(3)

	words := []string{"apple", "banana", "apple", "orange", "grape", "banana", "kiwi"}

	for _, word := range words {
		wc.CountWord(word)
	}
	pp.Print(wc.counts)
	
}
