package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {

	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run()

	// feed := newsfeed.New()

	// fmt.Println(feed)

	// feed.Add(newsfeed.Item{"Hello", "How ya' doing mate?"})

	// fmt.Println(feed)
}
