package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"example.com/basic-go/ch04/ex11/github/issues"
)

func main() {
	// 4actions
	action := flag.String("action", "list", "action")

	switch *action {
	case "list":
		showIssues()
	default:
		os.Exit(1)
	}
}

func showIssues() {
	result, err := issues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
