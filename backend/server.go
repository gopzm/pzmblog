package main

import (
	"fmt"
	"log"
	"net/http"

	srpb "pzm/blog/proto/summary_renderer"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	summary := &srpb.PzmBlogSummaryRenderer{
		Title:  "My first Blog",
		Author: "pzm",
	}
	fmt.Fprintf(w, "Hi there, I love %s!", summary)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
