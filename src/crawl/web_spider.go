package crawl

import (
	"time"

	"github.com/drag/src/util/dsa"
	"github.com/drag/src/util/logger"
)

//Config for WebSpider
type Config struct {
	TimeoutMinute int
	Search        string
}

// WebSpider everything stars here
type WebSpider struct {
	queue   dsa.CQueueString
	Results []Result
}

// Start let the lord be with you
func (c *WebSpider) Start(option Config) error {
	// intialize
	c.queue.IntializeDefultValues()
	var rschan = make(chan Result)
	go c.appendResult(rschan)

	// start
	content, seedErr := SeedRequestForXPage(option.Search)
	if seedErr != nil {
		return seedErr
	}
	rslt, retrieveErr := RetrieveInfoOnXPage(content)
	if retrieveErr != nil {
		return retrieveErr
	}
	c.queue.Enqueue(rslt.Links...)
	c.Results = append(c.Results, rslt)

	go c.spin(rschan)
	go c.spin(rschan)
	go c.spin(rschan)

	time.Sleep(time.Duration(option.TimeoutMinute) * time.Minute)
	return nil
}

// spin over links in queue
func (c *WebSpider) spin(rschan chan<- Result) {
	for c.queue.Size() > 0 {
		lnk, _ := c.queue.Dequeue()
		content, wbReqErr := WebRequest(lnk)
		if wbReqErr != nil {
			logger.Printf(
				logger.ErrorLevel,
				"on WebRequest\n\t %#v",
				wbReqErr)
		} else {
			result, resultErr := RetrieveInfoOnXPage(content)
			if resultErr != nil {
				logger.Printf(
					logger.ErrorLevel,
					"on retrieveInfoOnXpage for %s \n\t %#v",
					lnk,
					resultErr)
			}
			c.queue.Enqueue(result.Links...)
			rschan <- result
			// logger.Printf(logger.DebugLevel, "%#v", len(c.Results))
		}
	}
}

func (c *WebSpider) appendResult(rschan <-chan Result) {

	for rs := range rschan {
		c.Results = append(c.Results, rs)
	}

}
