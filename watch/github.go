package watch

import (
	"fmt"
	"strings"
	"time"
)

type GitHub struct {
	Owner   string
	Repo    string
	Commit  string
	Version string
}

func getGitHubCommit(owner, repo string) (string, *time.Time, error) {
	data, err := atomFetch(strings.Join([]string{"https://github.com", owner, repo, "commits.atom"}, "/"))
	if err != nil {
		return "", nil, err
	}
	entry, err := atomGetLatest(data)
	if err != nil {
		return "", nil, err
	}
	commit := strings.Split(entry.Id, "/")[1]
	date, err := time.Parse("2006-01-02T15:04:05Z", entry.Updated)
	if err != nil {
		return "", nil, err
	}
	return commit, &date, nil
}

func getGitHubTag(owner, repo string) (string, *time.Time, error) {
	data, err := atomFetch(strings.Join([]string{"https://github.com", owner, repo, "tags.atom"}, "/"))
	if err != nil {
		return "", nil, err
	}
	entry, err := atomGetLatest(data)
	if err != nil {
		return "", nil, err
	}
	tag := strings.Split(entry.Id, "/")[2]
	date, err := time.Parse("2006-01-02T15:04:05Z", entry.Updated)
	if err != nil {
		return "", nil, err
	}
	return tag, &date, nil
}

func WatchGitHub(owner, repo string) (*GitHub, error) {
	result := &GitHub{Owner: owner, Repo: repo}

	commit, date, err := getGitHubCommit(owner, repo)
	if err != nil {
		return result, err
	}
	result.Commit = commit

	tag, _, err := getGitHubTag(owner, repo)
	if err == nil {
		result.Version = fmt.Sprintf("%s+git%s.%7.7s", tag, date.Format("20060102"), commit)
	} else if err.Error() == "No Entries" {
		result.Version = fmt.Sprintf("0.0~git%s.%7.7s", date.Format("20060102"), commit)
	} else {
		return result, err
	}

	return result, nil
}
