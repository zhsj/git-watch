package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	service = flag.String("service", "github", "Service Type")
	owner   = flag.String("owner", "", "GitHub repository owner")
	repo    = flag.String("repo", "", "GitHub repository name")
	web     = flag.Bool("web", false, "Run as web service")
	listen  = flag.String("listen", ":9000", "Web server listen address")
)

func init() {
	flag.Parse()
}

func main() {
	if *web {
		http.HandleFunc("/github/", githubHandler)
		http.HandleFunc("/github-download/", githubDownloader)
		http.ListenAndServe(*listen, nil)
		return
	}
	switch *service {
	case "github":
		if *owner == "" || *repo == "" {
			fmt.Println("Error: unknown repository")
			flag.Usage()
			return
		}
		result, err := WatchGitHub(*owner, *repo)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result.Version)
		}
	default:
		fmt.Println("Error: unknown service type")
		flag.Usage()
	}
}
