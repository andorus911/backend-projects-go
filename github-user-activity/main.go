package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {

	username := os.Args[1]

	if len(username) == 0 {
		err := fmt.Errorf("no arguments")
		log.Fatalln(err)
		return
	}

	resp, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		log.Fatal(err)
	}

	history := make(map[string]map[string]int)
	for _, v := range events {
		if _, ok := history[v.Type]; !ok {
			history[v.Type] = make(map[string]int)
		}

		history[v.Type][v.Repo.Name]++
	}

	fmt.Println("Output:")
	for eventName, v := range history {
		for repoName, sum := range v {

			switch eventName {
			case "CommitCommentEvent":
				fmt.Printf(" - Wrote %d commit comments in %s", sum, repoName)
			case "CreateEvent":
				fmt.Printf(" - Created %d branches/tags at %s", sum, repoName)
			case "DeleteEvent":
				fmt.Printf(" - Deleted %d branches/tags at %s", sum, repoName)
			case "ForkEvent":
				fmt.Printf(" - Created a fork of %s", repoName)
			case "GollumEvent":
				fmt.Printf(" - Updated %d wiki pages in %s", sum, repoName)
			case "IssueCommentEvent":
				fmt.Printf(" - Updated %d comments on PR/issue at %s", sum, repoName)
			case "IssuesEvent":
				fmt.Printf(" - Created %d issue updates at %s", sum, repoName)
			case "MemberEvent":
				fmt.Printf(" - Added %d collaborators to %s", sum, repoName)
			case "PublicEvent":
				fmt.Printf(" - Published %d repositories in %s", sum, repoName)
			case "PullRequestEvent":
				fmt.Printf(" - Updated %d PRs to %s", sum, repoName)
			case "PullRequestReviewEvent":
				fmt.Printf(" - Created %d PRs to %s", sum, repoName)
			case "PullRequestReviewCommentEvent":
				fmt.Printf(" - Updated %d PR comments at %s", sum, repoName)
			case "PullRequestReviewThreadEvent":
				fmt.Printf(" - Updated %d PR review threads at %s", sum, repoName)
			case "PushEvent":
				fmt.Printf(" - Pushed %d commits to %s", sum, repoName)
			case "ReleaseEvent":
				fmt.Printf(" - Created/updated %d releases at %s", sum, repoName)
			case "SponsorshipEvent":
				fmt.Printf(" - Updated a sponsorship listing %d times %s", sum, repoName)
			case "WatchEvent":
				fmt.Printf(" - Starred %s", repoName)
			default:
				fmt.Printf(" ! NEW EVENT TYPE (%s) : %d time for %s", eventName, sum, repoName)
			}

			fmt.Println()
		}
	}
}
