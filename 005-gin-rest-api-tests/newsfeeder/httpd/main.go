package main

import (
	"fmt"

	"newsfeeder/platform/newsfeed"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", handler.PingGet())
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	feed := newsfeed.New()
	fmt.Println(feed)
	feed.Add(newsfeed.Item{"Hello", " How ya doing mate?"})
	fmt.Println(feed)

}
