package main

import (
	"testing"
	"time"

	"github.com/drag/src/crawl"
	"github.com/drag/src/util/logger"
)

func TestRequest(t *testing.T) {
	var w crawl.WebSpider
	w.Start(crawl.Config{
		Timeout: 1 * time.Minute,
		Search:  "covid",
	})

	sum := 0
	for _, r := range w.Results {
		sum += len(r.Links)
	}
	logger.Println(logger.InfoLevel,
		"Look Up", len(w.Results),
		",",
		"Links", sum)
	// logger.Printf(logger.DebugLevel, "%#v", w.Results)
}
