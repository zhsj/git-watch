package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/zhsj/git-watch/watch"
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
		log.Printf("About to listen on %s", *listen)
		log.Fatal(http.ListenAndServe(*listen, nil))
	}
	switch *service {
	case "github":
		if *owner == "" || *repo == "" {
			fmt.Println("Error: unknown repository")
			flag.Usage()
			return
		}
		result, err := watch.WatchGitHub(*owner, *repo)
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
