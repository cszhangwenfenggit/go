package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issues
}
type Issues struct {
	ID        int
	Number    int
	Title     string
	State     string
	Body      string
	HtmlURL   string `json:"html_url"`
	CreatedAt string `json:"created_at"`
	User      *User
}

type User struct {
	Login   string
	HtmlURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesResult, error) {
	//todo 拼接url请求参数
	q := IssuesURL + "?q=" + url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(q)
	if err != nil {
		return nil, err
	}
	//todo 根据接口返回结果，判断是否成功
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed:%s", resp.Status)
	}
	//todo 请求接口，获取数据
	var result IssuesResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	//todo
	resp.Body.Close()

	return &result, err
}
