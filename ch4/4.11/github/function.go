package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GET /repos/{owner}/{repo}/issues
func (p Params) GetIssues() ([]Issue, error) {
	url := IssuesURL + p.Owner + "/" + p.Repo + "/issues"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}

	return issues, nil
}

// GET /repos/{owner}/{repo}/issues/{issue_number}
func (p Params) GetIssue() (Issue, error) {
	url := IssuesURL + p.Owner + "/" + p.Repo + "/issues/" + p.Number

	resp, err := http.Get(url)
	if err != nil {
		return Issue{}, err
	}
	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return Issue{}, err
	}

	return issue, nil
}

// POST /repos/{owner}/{repo}/issues
func (p Params) CreateIssue() bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(p.Issue); err != nil {
		return false
	}

	url := IssuesURL + p.Owner + "/" + p.Repo + "/issues" +
		"?access_token=" + p.Token

	_, err := http.Post(url, "application/json", &buf)

	return err == nil
}

// PATCH /repos/{owner}/{repo}/issues/{issue_number}
func (p Params) EditIssue() bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(p.Issue); err != nil {
		return false
	}

	url := IssuesURL + p.Owner + "/" + p.Repo + "/issues/" + p.Number + "?access_token=" + p.Token

	request, err := http.NewRequest("PATCH", url, &buf)
	if err != nil {
		return false
	}

	request.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(request)

	return err == nil
}

// PATCH /repos/{owner}/{repo}/issues/{issue_number}
func (p Params) CloseIssue() bool {
	url := IssuesURL + p.Owner + "/" + p.Repo + "/issues/" + p.Number + "?access_token=" + p.Token

	request, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		return false
	}

	request.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(request)

	return err == nil
}
