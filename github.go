package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHub struct {
	Owner   string
	Repo    string
	Commit  string
	Version string
}

var (
	client *github.Client
	ctx    context.Context
)

func init() {
	ctx = context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}
	client.UserAgent = "github.com/zhsj/git-watch"
}

func WatchGitHub(owner, repo string) (result *GitHub, err error) {
	result = &GitHub{Owner: owner, Repo: repo}

	headSHA1, _, err := client.Repositories.GetCommitSHA1(ctx, owner, repo, "HEAD", "")
	if err != nil {
		return
	}
	result.Commit = headSHA1

	headCommit, _, err := client.Repositories.GetCommit(ctx, owner, repo, headSHA1)
	headDate := headCommit.Commit.Committer.GetDate()

	tags, _, err := client.Repositories.ListTags(ctx, owner, repo, nil)
	if err != nil {
		return
	}
	if len(tags) == 0 {
		var allDayCommits []*github.RepositoryCommit
		opt := &github.CommitsListOptions{
			Since: headDate.Truncate(24 * time.Hour),
			Until: headDate,
		}
		for {
			dayCommits, resp, e := client.Repositories.ListCommits(ctx, owner, repo, opt)
			if e != nil {
				return nil, e
			}
			allDayCommits = append(allDayCommits, dayCommits...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
		result.Version = fmt.Sprintf("0.0~git%s.%d.%7.7s", headDate.Format("20060102"), len(allDayCommits), headSHA1)
		return
	}

	latestTag := tags[0]
	compare, _, err := client.Repositories.CompareCommits(ctx, owner, repo, latestTag.GetName(), headSHA1)
	result.Version = fmt.Sprintf("%s+git%s.%d.%7.7s", latestTag.GetName(), headDate.Format("20060102"),
		compare.GetTotalCommits(), headSHA1)
	return
}
