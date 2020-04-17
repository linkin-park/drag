package helper

import (
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// ParseHTMLByTag you parse html by tags given in struct
func ParseHTMLByTag(content string, parseParam *ParseParam) {
	tokenizer := html.NewTokenizer(strings.NewReader(content))
	var wg = &sync.WaitGroup{}

	// wait for all the go funcs to be completed
	defer wg.Wait()

	// loop over every token available
	for {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			// End of document
			return
		case tokenType == html.StartTagToken:
			// no nested loop
			func(t html.Token) {
				for _, tagDetail := range parseParam.TagDetails {

					if tagExist :=
						strings.ToLower(t.Data) == strings.ToLower(tagDetail.Tag); tagExist {

						wg.Add(1)
						go tagDetail.CallFunc(t, wg)
					}
				}
			}(tokenizer.Token())

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
	CallFunc func(html.Token, *sync.WaitGroup)
}
