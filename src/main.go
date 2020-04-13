package main

import (
	"fmt"
	"net/http"

	"github.com/drag/src/util"
)

var seedURL = []string{
	"https://duckduckgo.com/?q=",
}

func main() {
	// req for X [via ddg]
	// get the links from X page , break to chunk and spin goroutines.
	// add to queue [if not in visitedMap]
	// if visited add to visitedMap link , bool(visited)

}

func request(input string) (string, error) {
	defer func() {
		err := recover()
		util.DebugLog.Printf("Error on Request %#v", err)
		util.ErrorLog.Printf("Error on Request %#v", err)
	}()

	// sanitize input , random choose of seedURL TODO()
	httpResp, err := http.Get(seedURL[0] + input)
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
			if err != nil {
				return "", err
			}

			finalByts = append(finalByts, tmp[:rdN]...)
			rdN, err = httpResp.Body.Read(tmp)
		}

		return string(finalByts), nil
	}()
}
