package task05

import (
	"container/heap"
	"strings"
)

func topKFrequent(words []string, k int) []string {
	if len(words) == 0 || k <= 0 {
		return nil
	}

	var (
		i        int
		wordFreq = wordUniq(words)
		hp       = new(myMinHeap)
	)
	for word, freq := range wordFreq {
		if i < k {
			heap.Push(hp, &element{word: word, freq: freq})
			i++
		} else {
			if freq > (*hp)[0].freq ||
				(freq == (*hp)[0].freq && strings.Compare(word, (*hp)[0].word) == -1) {
				heap.Pop(hp)
				heap.Push(hp, &element{word: word, freq: freq})
			}
		}
	}

	var res []string
	for hp.Len() != 0 {
		res = append(res, heap.Pop(hp).(*element).word)
	}
	reverse(res)

	return res
}

// wordUniq returns a map which mapping word to frequent.
func wordUniq(words []string) map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word] += 1
	}
	return wordFreq
}

func reverse(s []string) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

type element struct {
	word string
	freq int
}

type myMinHeap []*element

func (h myMinHeap) Len() int {
	return len(h)
}

func (h myMinHeap) Less(i, j int) bool {
	if h[i].freq < h[j].freq {
		return true
	}

	if h[i].freq == h[j].freq && strings.Compare(h[i].word, h[j].word) == 1 {
		return true
	}

	return false
}

func (h myMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myMinHeap) Push(v interface{}) {
	*h = append(*h, v.(*element))
}

func (h *myMinHeap) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
