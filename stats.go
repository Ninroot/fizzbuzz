package main

import (
	"fmt"
	"math"
)

type stat struct {
	// size of the requests window to be saved
	window  int
	counter map[string]int
}

func NewStat(window int) stat {
	if window < 1 {
		window = 1
	}
	return stat{window: window, counter: make(map[string]int)}
}

// save saves a given request and increment its counter
func (s stat) save(req string) {
	if _, prs := s.counter[req]; prs {
		s.counter[req]++
		return
	}

	if len(s.counter) >= s.window {
		k, err := minKey(s.counter)
		if err == nil {
			delete(s.counter, k)
		}
	}
	s.counter[req]++
}

func minKey(m map[string]int) (string, error) {
	if len(m) == 0 {
		return "", fmt.Errorf("m is empty")
	}

	minK := ""
	minV := math.MaxInt
	for k, v := range m {
		if v <= minV {
			minV = v
			minK = k
		}
	}
	return minK, nil
}
