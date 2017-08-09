package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New("tmpl").Parse(
	`<a href="/github-download/{{ .Owner }}/{{ .Repo }}/{{ .Commit }}/{{ .Version }}.tar.gz">{{ .Version }}</a>
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, result)
}

func githubDownloader(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	paths := strings.Split(path[len("/github-download/"):len(path)-len(".tar.gz")], "/")
	fmt.Printf("%+v", paths[1])
	http.Redirect(w, r, fmt.Sprintf("https://github.com/%s/%s/archive/%s.tar.gz", paths[0], paths[1], paths[2]),
		http.StatusFound)
	return
}
