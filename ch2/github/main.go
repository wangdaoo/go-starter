package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Issue struct {
	Number int
	Title  string
	User   struct {
		Login string
	}
}

type Milestone struct {
	Number int
	Title  string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.github.com/repos/wangdaoo/el-filter/issues?state=closed")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var issues []Issue
		if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
			log.Fatal(err)
		}

		var milestones []Milestone
		if err := json.NewDecoder(resp.Body).Decode(&milestones); err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "Closed Issues:\n")
		for _, issue := range issues {
			fmt.Fprintf(w, "#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
		}

		fmt.Fprintf(w, "\nMilestones:\n")
		for _, milestone := range milestones {
			fmt.Fprintf(w, "#%-5d %.55s\n", milestone.Number, milestone.Title)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
