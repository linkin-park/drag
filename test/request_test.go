package main

import (
	"fmt"
	"testing"

	crawl "github.com/drag/src/crawl"
)

func scrapeOnXPage() {
}

func TestRequest(t *testing.T) {
	content, err := crawl.Request("golang")
	if err != nil {
		t.Error(err)
	}
	rslt, _ := crawl.RetrieveInfoOnXPage(content)
	fmt.Printf("%#v", rslt)
}
