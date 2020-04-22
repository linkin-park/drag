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
	var done = make(chan bool, 3)
	var rcvrchan = make(chan bool)

	go c.appendResult(rschan, done, rcvrchan)

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

	go c.spin(rschan, done)
	go c.spin(rschan, done)
	go c.spin(rschan, done)

	clearSignal(
		time.Duration(option.TimeoutMinute),
		time.Minute,
		done,
		rcvrchan,
		rschan)

	return nil
}

// spin over links in queue
func (c *WebSpider) spin(
	rschan chan<- Result,
	done <-chan bool,
) {

	for c.queue.Size() > 0 {
		select {
		case <-done:
			return
		default:
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
}

func clearSignal(
	x time.Duration,
	d time.Duration,
	done chan<- bool,
	rcvrdone chan<- bool,
	rschan chan<- Result) {

	select {
	case <-time.After(time.Duration(x) * d):
		done <- true
		done <- true
		done <- true
		close(done)

		rcvrdone <- true
		close(rcvrdone)

		close(rschan)
		return
	}
}

func (c *WebSpider) appendResult(
	rschan <-chan Result,
	done <-chan bool,
	rcvrdone <-chan bool) {

	for {
		select {
		case <-rcvrdone:
			return
		default:
			if rs, ok := <-rschan; ok {
				c.Results = append(c.Results, rs)
			}
		}
	}
}
