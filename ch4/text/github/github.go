package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}
type Issue struct {
	Number int
	HTMLURL string `json: "html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"create_at"`
	Body string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	//fmt.Println(terms)
	//fmt.Println(q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	//fmt.Println(resp.Body)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.) if resp.StatusCode != http.StatusOK {
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}