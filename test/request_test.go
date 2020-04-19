package main

import (
	"fmt"
	"testing"

	crawl "github.com/drag/src/crawl"
	"github.com/drag/src/util/logger"
)

func scrapeOnXPage() {
}

func TestRequest(t *testing.T) {
	content, err := crawl.SeedRequestForXPage("black magic")
	if err != nil {
		t.Error(err)
	}
	rslt, _ := crawl.RetrieveInfoOnXPage(content)
	logger.Printf(logger.InfoLevel, "%#v\n\n", rslt)
	fmt.Println("links size:", len(rslt.Links))
}
