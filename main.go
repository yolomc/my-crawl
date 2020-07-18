package main

import (
	"fmt"

	"github.com/yolomc/my-crawl/fetcher"
)

func main() {

	content, err := fetcher.Fetch("https://www.bilibili.com/ranking/all/129/0/3")
	if err != nil {
		fmt.Println(err)
	}

	//fetcher.ParseContent(content)
	fmt.Printf("%s", content)

}
