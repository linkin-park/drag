package main

import (
	"testing"

	"github.com/drag/src/crawl"
	"github.com/drag/src/util/logger"
)

func TestRequest(t *testing.T) {
	var w crawl.WebSpider
	w.Start(crawl.Config{
		TimeoutMinute: 1,
		Search:        "covid",
	})

	sum := 0
	for _, r := range w.Results {
		sum += len(r.Links)
	}
	logger.Println(logger.InfoLevel, len(w.Results), sum)
}
