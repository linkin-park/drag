package helper

import (
	"strings"
	"sync"

	"github.com/drag/src/util/logger"
	"golang.org/x/net/html"
)

// generate an AST? // not good on crawling!
// AST query !! // not required!

// ParseHTMLByTag you parse html by tags given in struct
func ParseHTMLByTag(content string, parseParam *ParseParam) {
	tokenizer := html.NewTokenizer(strings.NewReader(content))
	var wg = &sync.WaitGroup{}

	// wait for all the go funcs to be completed
	defer wg.Wait()

	// loop over every token available
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			// End of document
			return
		case html.StartTagToken:
			// no nested loop
			func(t html.Token, toknzr *html.Tokenizer) {
				defer func() {
					if err := recover(); err != nil {
						logger.Printf(
							logger.ErrorLevel,
							"\t %#v",
							err)
					}
				}()

				for _, tagDetail := range parseParam.TagDetails {

					if tagExist :=
						strings.ToLower(t.Data) == strings.ToLower(tagDetail.Tag); tagExist {

						tagDetail.CallFunc(t, toknzr)
					}
				}
			}(tokenizer.Token(), tokenizer)

		}
	}
}

// ParseParam a wrapper over
// slice of TagDetail
type ParseParam struct {
	TagDetails []TagDetail
}

// TagDetail of every token to be crawled
type TagDetail struct {
	Tag      string
	CallFunc func(html.Token, *html.Tokenizer)
}
