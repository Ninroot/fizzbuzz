package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzRoute(t *testing.T) {
	router := setupRouter()

	t.Run("nominal", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := []byte(`{
			"int1":3,
			"int2":5,
			"limit":15,
			"str1":"Fizz",
			"str2":"Buzz"
		}`)

		req, err := http.NewRequest("POST", "/fizzbuzz", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("could not make request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, `["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]`, w.Body.String())
	})

	t.Run("bad request - missing parameters", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := []byte(`{
			"int1":3,
			"int2":5
		}`)
		req, err := http.NewRequest("POST", "/fizzbuzz", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("could not make request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	t.Run("unprocessable param - wrong limit", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := []byte(`{
			"int1":3,
			"int2":5,
			"limit":-15,
			"str1":"Fizz",
			"str2":"Buzz"
		}`)

		req, err := http.NewRequest("POST", "/fizzbuzz", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("could not make request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		router.ServeHTTP(w, req)

		assert.Equal(t, 422, w.Code)
		assert.Equal(t, `{"msg":"limit must be greater than 0"}`, w.Body.String())
	})
}

func TestStatsRoute(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		router := setupRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/stats", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{}", w.Body.String())
	})

	t.Run("nominal", func(t *testing.T) {
		router := setupRouter()
		for i := 0; i < 3; i++ {
			postFizzBuzz(t, router, []byte(`{"int1":3,"int2":5,"limit":15,"str1":"Fizz","str2":"Buzz"}`))
			postFizzBuzz(t, router, []byte(`{"int1":2,"int2":3,"limit":10,"str1":"Abc","str2":"Def"}`))
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/stats", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, `{"{10 2 3 Abc Def}":3,"{15 3 5 Fizz Buzz}":3}`, w.Body.String())
	})
}

func postFizzBuzz(t *testing.T, router *gin.Engine, body []byte) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/fizzbuzz", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("could not make request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
