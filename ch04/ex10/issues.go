package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	month := make([]*github.Issue, 0)
	year := make([]*github.Issue, 0)
	overYear := make([]*github.Issue, 0)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(now.AddDate(0, -1, 0)):
			month = append(month, item)
		case item.CreatedAt.After(now.AddDate(-1, 0, 0)):
			year = append(year, item)
		default:
			overYear = append(overYear, item)
		}
	}

	fmt.Printf("1ヶ月未満")
	printIssues(month)
	fmt.Printf("1年未満")
	printIssues(year)
	fmt.Printf("1年以上")
	printIssues(overYear)
}

func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
