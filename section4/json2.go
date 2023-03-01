package main

import (
	"fmt"
	"log"
	"os"
	"pratice/section4/github"
)

func main() {
	data, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:", data.TotalCount)
	for _, item := range data.Items {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
