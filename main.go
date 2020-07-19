package main

import (
	"github.com/yolomc/my-crawl/engine"
	"github.com/yolomc/my-crawl/parse"
)

func main() {

	// content, err := fetcher.Fetch("https://www.bilibili.com/ranking/all/129/0/3")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//fetcher.ParseContent(content)
	//fmt.Printf("%s", content)

	engine.Run(engine.Request{
		Url:       "https://book.douban.com/tag/",
		ParseFunc: parse.ParseContent,
	})

}
