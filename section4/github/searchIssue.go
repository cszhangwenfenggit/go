package github

import (
	"fmt"
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
	CreatedAt int    `json:"created_at"`
	User      *User
}

type User struct {
	Login   string
	HtmlURL string `json:"html_url"`
}

func SearchIssues(terms []string) {
	fmt.Println(strings.Join(terms, " "))
}
