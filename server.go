package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const defaultStatsWindow = 10000

type fizzbuzzRequest struct {
	Limit int    `json:"limit" binding:"required"`
	Int1  int    `json:"int1" binding:"required"`
	Int2  int    `json:"int2" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}

var stats stat

func defaultEnv(env string, def int) int {
	if env == "" {
		return def
	}
	ret, err := strconv.Atoi(env)
	if err != nil {
		return def
	}
	return ret
}

func setupRouter() *gin.Engine {
	stats = NewStat(defaultEnv(os.Getenv("STATS_WINDOW"), defaultStatsWindow))
	gin.DisableConsoleColor()

	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/stats", func(c *gin.Context) { c.JSON(200, stats.counter) })
	r.POST("/fizzbuzz", fizzBuzzHandler)

	return r
}

func fizzBuzzHandler(c *gin.Context) {
	var req fizzbuzzRequest
	if c.Bind(&req) == nil {
		res, err := fizzbuzz(req.Limit, req.Int1, req.Int2, req.Str1, req.Str2)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": err.Error()})
			return
		}
		stats.save(fmt.Sprint(req))
		c.JSON(http.StatusOK, res)
	}
}

func main() {
	r := setupRouter()
	r.Run()
}
