package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {
	links, err := Extract2("http://c.biancheng.net/view/5712.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.Join(links, "\n"))
}

func Extract2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s failed:%s", url, err)
	}

	doc, err := html.Parse(resp.Body)

	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parse %s failed:%s", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	foreachNode(doc, visitNode)

	return links, nil
}

func foreachNode(n *html.Node, pre func(n *html.Node)) {
	if n != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = n.NextSibling {
		foreachNode(c, pre)
	}
}
