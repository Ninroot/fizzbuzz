package main

import (
	"fmt"
	"strconv"
)

// Returns a list of strings with numbers from 1 to limit, where: all multiples
// of int1 are replaced by str1, all multiples of int2 are replaced by str2, all
// multiples of int1 and int2 are replaced by str1str2.
func fizzbuzz(limit, int1, int2 int, str1, str2 string) ([]string, error) {
	if limit < 0 {
		return []string{}, fmt.Errorf("limit must be greater than 0")
	}
	if int1 == 0 || int2 == 0 {
		return []string{}, fmt.Errorf("cannot use 0 as fizzbuzz parameter")
	}
	s := make([]string, limit)
	concat := str1 + str2
	for i := 1; i <= limit; i++ {
		var fb string
		mod1 := i%int1 == 0
		mod2 := i%int2 == 0
		if mod1 && mod2 {
			fb = concat
		} else if mod1 {
			fb = str1
		} else if mod2 {
			fb = str2
		} else {
			fb = strconv.Itoa(i)
		}
		s[i-1] = fb
	}
	return s, nil
}
