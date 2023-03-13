package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzzNominal(t *testing.T) {
	expect := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	actual, err := fizzbuzz(15, 3, 5, "Fizz", "Buzz")
	if err != nil {
		log.Fatal("should not return an error: ", err)
	}
	assert.Equal(t, expect, actual)
}

func TestFizzbuzzNegative(t *testing.T) {
	expect := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	actual, err := fizzbuzz(15, -3, -5, "Fizz", "Buzz")
	if err != nil {
		log.Fatal("should not return an error: ", err)
	}
	assert.Equal(t, expect, actual)
}

func TestFizzbuzzFailureZero(t *testing.T) {
	actual, err := fizzbuzz(15, 0, 5, "Fizz", "Buzz")
	assert.Equal(t, []string{}, actual)
	if err == nil {
		log.Fatal("should not accept zero parameter")
	}
}

func TestFizzbuzzFailureNegativeLimit(t *testing.T) {
	actual, err := fizzbuzz(-15, 3, 5, "Fizz", "Buzz")
	assert.Equal(t, []string{}, actual)
	if err == nil {
		log.Fatal("should not accept a negative limit")
	}
}

func TestFizzbuzzEmptyStrings(t *testing.T) {
	expect := []string{"1", "2", "", "4", "", "", "7", "8", "", "", "11", "", "13", "14", ""}
	actual, err := fizzbuzz(15, 3, 5, "", "")
	if err != nil {
		log.Fatal("should not return an error: ", err)
	}
	assert.Equal(t, expect, actual)
}

func BenchmarkFizzBuzzNominal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fizzbuzz(100, 3, 5, "Fizz", "Buzz")
	}
}
