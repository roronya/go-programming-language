package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var c, r, u, d bool
	flag.BoolVar(&c, "c", false, "create")
	flag.BoolVar(&r, "r", false, "read")
	flag.BoolVar(&u, "u", false, "update")
	flag.BoolVar(&d, "d", false, "delete")
	flag.Parse()

	if c {
		if len(flag.Args()) != 4 {
			log.Fatal()
		}
		create(flag.Arg(0), flag.Arg(1), flag.Arg(2), flag.Arg(3))
	} else if r {
		if len(flag.Args()) != 3 {
			log.Fatal()
		}
		read(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	} else if u {
		if len(flag.Args()) != 5 {
			log.Fatal()
		}
		update(flag.Arg(0), flag.Arg(1), flag.Arg(2), flag.Arg(3), flag.Arg(4))
	} else if d {
		delete_(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	}
}

func create(owner string, repo string, title string, body string) {
	result := CreateIssue(owner, repo, map[string]string{"title": title, "body": body})
	PrintIssue(result)
}

func read(owner string, repo string, issueNumber string) {
	result := ReadIssue(owner, repo, issueNumber)
	PrintIssue(result)
}

func update(owner string, repo string, issueNumber string, title string, body string) {
	result := EditIssue(owner, repo, issueNumber, map[string]string{"title": title, "body": body})
	PrintIssue(result)
}

func delete_(owner string, repo string, issueNumber string) {
	result := EditIssue(owner, repo, issueNumber, map[string]string{"state": "close"})
	PrintIssue(result)
}

func PrintIssue(issue Issue) {
	fmt.Printf("%d %s status:%s\n%s\n%s", issue.Number, issue.Title, issue.State, issue.Body, issue.CreatedAt.String())
}

const URL = "https://api.github.com/repos"

type Issue struct {
	Number    int
	Title     string
	Body      string
	State     string
	CreatedAt time.Time `json:"created_at"`
}

func EditIssue(owner string, repo string, issueNumber string, params map[string]string) Issue {
	json_, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("%s/%s/%s/issues/%s", URL, owner, repo, issueNumber)
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(json_))
	req.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("GITHUB_TOKEN")))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("failed %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
	}
	resp.Body.Close()

	return result
}

func ReadIssue(owner string, repo string, issueNumber string) Issue {
	url := fmt.Sprintf("%s/%s/%s/issues/%s", URL, owner, repo, issueNumber)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("failed %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
	}
	resp.Body.Close()

	return result
}

func CreateIssue(owner string, repo string, params map[string]string) Issue {
	json_, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("%s/%s/%s/issues", URL, owner, repo)
	resp, err := http.Post(url, "", bytes.NewBuffer(json_))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		log.Fatalf("failed %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
	}
	resp.Body.Close()

	return result
}
