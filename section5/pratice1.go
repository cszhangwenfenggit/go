package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	fmt.Println(Max())

	arr := []string{"Hello", "World"}
	fmt.Println(StringJoin(arr, "=>"))

	links, err := ElementsByTageName("http://c.biancheng.net/view/5712.html", "a")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(links)
}

func Max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	var max int

	for i := 0; i < len(vals); i++ {
		if max < vals[i] {
			max = vals[i]
		}
	}

	return max
}

func StringJoin(arr []string, join string) string {
	var result string
	for k, s := range arr {
		if k == 0 {
			result += s
		} else {
			result += join + s
		}
	}
	return result
}

func ElementsByTageName(url string, name ...string) ([]*html.Node, error) {

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
		return nil, fmt.Errorf("prase %s Html:%s", url, err)
	}

	var result []*html.Node

	targetNode := func(n *html.Node) {
		for _, tag := range name {
			if n.Type == html.ElementNode && n.Data == tag {
				result = append(result, n)
			}
		}
	}

	foreachNodes(doc, targetNode)

	return result, nil
}

func foreachNodes(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		foreachNodes(c, pre)
	}
}
