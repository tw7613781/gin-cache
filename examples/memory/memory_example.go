package main

import (
	"time"

	"github.com/gin-gonic/gin"
	cache "github.com/tw7613781/gin-cache"
	"github.com/tw7613781/gin-cache/persist"
)

func main() {
	app := gin.New()

	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	app.GET("/hello",
		cache.CacheByRequestURI(memoryStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
