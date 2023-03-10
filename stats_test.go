package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStat(t *testing.T) {
	t.Run("nominal", func(t *testing.T) {
		s := NewStat(3)
		for i := 0; i < 15; i++ {
			s.save("a")
		}
		s.save("b")
		assert.Equal(t, map[string]int{"a": 15, "b": 1}, s.counter)
	})

	t.Run("rotation", func(t *testing.T) {
		s := NewStat(3)
		s.save("a")
		for i := 0; i < 3; i++ {
			s.save("b")
			s.save("c")
		}
		s.save("d")
		assert.Equal(t, map[string]int{"b": 3, "c": 3, "d": 1}, s.counter)
	})

	t.Run("0 maxSize", func(t *testing.T) {
		s := NewStat(0)
		s.save("a")
		s.save("b")
		assert.Equal(t, map[string]int{"b": 1}, s.counter)
	})
}
