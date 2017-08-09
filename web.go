package main

import (
	"net/http"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New("tmpl").Parse(
	`<a href="https://github.com/{{ .Owner }}/{{ .Repo }}/archive/{{ .Commit }}.tar.gz">{{ .Version }}</a>
`))

func githubHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path[len("/github/"):], "/")
	if len(paths) != 2 {
		http.Error(w, "Wrong URL format", http.StatusBadRequest)
		return
	}
	owner := paths[0]
	repo := paths[1]
	result, err := WatchGitHub(owner, repo)
	if err != nil {
		http.Error(w, "Fetch latest version failed", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, result)
}
