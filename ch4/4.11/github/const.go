// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

const IssuesURL = "https://api.github.com/search/repos/"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Params struct {
	Owner  string
	Repo   string
	Number string
	Token  string
	Issue
}

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
