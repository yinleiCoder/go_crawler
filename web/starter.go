package main

import (
	"goCrawler/web/controller"
	"log"
	"net/http"
)

func main() {

	// listen index.html
	http.Handle("/", http.FileServer(http.Dir("/home/yinlei/Desktop/go_crawler/web/view")))

	// listen search result.html
	http.Handle("/search", controller.CreateSearchResultHandler(
		"/home/yinlei/Desktop/go_crawler/web/view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println(err)
	}
}