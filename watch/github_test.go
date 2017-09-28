package watch

import "testing"

func TestGetGitHubCommit(t *testing.T) {
	commit, time, err := getGitHubCommit("elves", "elvish")
	t.Logf("%s, %s, %s", commit, time, err)
}

func TestGetGitHubTag(t *testing.T) {
	tag, time, err := getGitHubTag("elves", "elvish")
	t.Logf("%s, %s, %s", tag, time, err)
}
