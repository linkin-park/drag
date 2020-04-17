package crawl

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/drag/src/util/helper"
	"github.com/drag/src/util/logger"
	"golang.org/x/net/html"
)

var seedURL = []string{
	"https://www.google.com/search?source=hp&q=%s&oq=%s",
	"https://www.google.com/search?q=%s",
	"https://duckduckgo.com/?q=%s",
}

// Result holds the information
// often used with crawl
type Result struct {
	Title, Description string
	Keywords, Links    []string
}

var webClient = http.Client{
	Timeout: 30 * time.Second,
}

// RetrieveInfoOnXPage will crawl over the content given
// often used in conjunction with Result
func RetrieveInfoOnXPage(content string) (Result, error) {
	var result Result
	param := &helper.ParseParam{
		TagDetails: []helper.TagDetail{
			{
				Tag: "a",
				CallFunc: func(t html.Token, wg *sync.WaitGroup) {
					defer wg.Done()

					for _, attr := range t.Attr {
						if strings.ToLower(attr.Key) == "href" {
							if func() bool {
								return strings.HasPrefix(attr.Val, "http://") ||
									strings.HasPrefix(attr.Val, "https://")
							}() {
								result.Links = append(result.Links, attr.Val)
							}
						}
					}
				},
			},
			{
				Tag: "meta",
				CallFunc: func(t html.Token, wg *sync.WaitGroup) {
					defer wg.Done()

					var isDesc = false
					var isKeyword = false
					for _, attr := range t.Attr {
						// handling meta tag attr
						switch {
						case strings.ToLower(attr.Key) == "name" &&
							strings.ToLower(attr.Val) == "description":
							isDesc = true
						case isDesc &&
							strings.ToLower(attr.Key) == "content":
							result.Description = attr.Val
						case strings.ToLower(attr.Key) == "name" &&
							strings.ToLower(attr.Val) == "Keywords":
							isKeyword = true
						case isKeyword &&
							strings.ToLower(attr.Key) == "content":
							result.Keywords = strings.Split(attr.Val, ",")
						}
					}
				},
			},
		},
	}
	helper.ParseHTMLByTag(content, param)
	return result, nil
}

// Request a webRequest for the input given
func Request(input string) (string, error) {
	defer func() {
		err := recover()
		if err != nil {
			logger.DebugLog.Printf("Error on Request %#v", err)
			logger.ErrorLog.Printf("Error on Request %#v", err)
		}
	}()

	// sanitize input , random choose of seedURL TODO()
	httpResp, err := webClient.Get(fmt.Sprintf(seedURL[0], input))
	if err != nil {
		return "", err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Invalid Status %d", httpResp.StatusCode)
	}

	return func() (string, error) {
		tmp := make([]byte, 1024)
		finalByts := make([]byte, 0)

		for rdN, err := httpResp.Body.Read(tmp); rdN > 0; {
			if err != nil && err.Error() != "EOF" {
				return "", err
			}

			finalByts = append(finalByts, tmp[:rdN]...)
			rdN, err = httpResp.Body.Read(tmp)
		}

		return string(finalByts), nil
	}()
}
