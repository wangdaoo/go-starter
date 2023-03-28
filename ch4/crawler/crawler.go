package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// 网站的根URL
	url := "https://golang.org"

	// 使用一个map来存储已经下载过的页面，避免重复下载
	pages := make(map[string]bool)

	// 使用一个通道来传递需要下载的页面URL
	workList := make(chan []string)

	// 开始时，我们从根URL开始爬取
	go func() { workList <- []string{url} }()

	// 使用协程并发地下载和解析页面
	for i := 0; i < 20; i++ {
		go func() {
			for urls := range workList {
				for _, u := range urls {
					if !pages[u] {
						pages[u] = true
						links := crawl(u)
						go func() { workList <- links }()
					}
				}
			}
		}()
	}
}

func crawl(url string) []string {
	fmt.Println("crawling", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	// 使用html.Parse函数解析页面中的链接
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	links := visit(nil, doc)
	return links
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link := a.Val
				if strings.HasPrefix(link, "http") {
					links = append(links, link)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
