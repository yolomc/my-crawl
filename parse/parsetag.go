package parse

import (
	"regexp"

	"github.com/yolomc/my-crawl/engine"
)

const regStr = `<a href="(/tag/.+)">([^<]+)</a>`

func ParseContent(content []byte) engine.ParseResult {
	reg := regexp.MustCompile(regStr)
	matches := reg.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://book.douban.com" + string(m[1]),
			ParseFunc: engine.NilParse,
		})
		result.Items = append(result.Items, m[2])
	}

	return result
}
