package main

import (
	"fmt"
	"sort"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	fmt.Println(strings.Join(topoSort(prereqs), "\n"))
}

func topoSortOld(m map[string][]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Println(keys)

	return keys
}

func topoSort(m map[string][]string) []string {

	isSend := make(map[string]bool)

	var order []string

	var visitAll = func(items []string) {}
	visitAll = func(items []string) {
		for _, k := range items {
			if !isSend[k] {
				isSend[k] = true
				visitAll(m[k])
				order = append(order, k)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	visitAll(keys)

	return order
}
