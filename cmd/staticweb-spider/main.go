package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/golang/groupcache"
	"strings"
)

func main() {
	// 指定是否异步多线程、线程数
	// 解析主域名，并且可以参数指定主域名
	// 下载主域名所有页面，以及直接依赖外部域名，可以指定需要全部下载的域名（默认包括主域名）
	// 代理模式：none、all、otherDomain

	// 责任链，依次对url进行判断是否需要处理

	c := colly.NewCollector()
	// c.SetProxy("http://localhost:1081")

	counter := groupcache.AtomicInt(0)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if counter.Get() > 10 {
			return
		}
		href := e.Attr("href")
		url := e.Request.AbsoluteURL(href)
		if strings.Contains(url, "pdf") {
			return
		}
		fmt.Println(href, url, strings.Contains(url, "专栏"))
		e.Request.Visit(url)
		counter.Add(1)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Response", string(r.Body))
	})

	c.Visit("https://learn.lianglianglee.com/")
}
