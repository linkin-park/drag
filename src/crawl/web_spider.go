package crawl

import (
	"sync"
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
	mux     sync.Mutex
}

// Start let the lord be with you
func (c *WebSpider) Start(
	option Config,
) error {
	defer logger.Println(logger.InfoLevel, "Start", "Complete")
	logger.Println(logger.InfoLevel, "Start", "Initiated")

	// intialize
	c.queue.IntializeDefultValues()
	var done = make(chan bool, 3)
	var wgProducer = new(sync.WaitGroup)

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

	wgProducer.Add(3)
	go c.spin(done, wgProducer)
	go c.spin(done, wgProducer)
	go c.spin(done, wgProducer)

	clearSignal(
		time.Duration(option.TimeoutMinute),
		time.Minute,
		done,
		wgProducer,
	)

	return nil
}

// spin over links in queue
func (c *WebSpider) spin(
	done <-chan bool,
	wgProducer *sync.WaitGroup,
) {

	defer logger.Println(
		logger.DebugLevel,
		"spin",
		"complete",
	)
	defer wgProducer.Done()

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
				c.appendResult(result)
			}
		}
	}
}

func clearSignal(
	x time.Duration,
	d time.Duration,
	done chan<- bool,
	wgProducer *sync.WaitGroup,
) {
	defer func() {
		logger.Println(logger.DebugLevel, "clearSignal", "Complete")
	}()

	select {
	case <-time.After(time.Duration(x) * d):
		done <- true
		done <- true
		done <- true
		close(done)
		wgProducer.Wait()

		return
	}
}

func (c *WebSpider) appendResult(
	rs Result,
) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.Results = append(c.Results, rs)
}
