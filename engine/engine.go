package engine

import (
	"fmt"

	"github.com/yolomc/my-crawl/fetcher"
)

func Run(seeds ...Request) {

	for _, r := range seeds {
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			panic(err)
		}

		// if r.ParseFunc == nil {
		// 	return
		// }
		parseResult := r.ParseFunc(body)
		for _, r := range parseResult.Requests {
			fmt.Printf("[%s]\n", r.Url)
		}

		fmt.Println(len(parseResult.Requests))
		if len(parseResult.Requests) > 1 {
			Run(parseResult.Requests[0:1]...)
		}
	}

}
