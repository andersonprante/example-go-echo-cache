package main

import (
	"net/http"
	"time"

	"github.com/coocood/freecache"
	cache "github.com/gitsight/go-echo-cache"
	"github.com/labstack/echo/v4"
)

func main() {
	c := freecache.NewCache(1024 * 1024)

	e := echo.New()
	e.Use(cache.New(&cache.Config{
		TTL:        10 * time.Second,
		Methods:    []string{"GET"},
		StatusCode: []int{200, 204},
		Refresh: func(r *http.Request) bool {
			cache := r.Header.Get("cache")
			return cache == "false"
		},
	}, c))

	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, time.Now().String())
		return nil
	})

	e.Start(":8080")
}
