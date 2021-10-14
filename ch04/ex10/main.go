package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/roronya/go-programming-language/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var month []*github.Issue
	var year []*github.Issue
	var other []*github.Issue
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)
	for _, item := range result.Items {
		if item.CreatedAt.After(monthAgo) {
			month = append(month, item)
		} else if item.CreatedAt.Before(yearAgo) {
			other = append(other, item)
		} else {
			year = append(year, item)
		}
	}
	fmt.Println("一ヶ月未満")
	for _, item := range month {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("一年未満")
	for _, item := range year {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("一年以上")
	for _, item := range year {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
