package main

import (
	"testing"

	crawl "github.com/drag/src/crawl"
)

func scrapeOnXPage() {
}

func TestRequest(t *testing.T) {
	_, err := crawl.Request("hi!")
	if err != nil {
		t.Error(err)
	}
}
