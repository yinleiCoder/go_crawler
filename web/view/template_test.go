package view

import (
	"goCrawler/model"
	"goCrawler/web/pagemodel"
	"html/template"
	"log"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	out, err := os.Create("template.test.html")
	if err != nil {
		log.Printf("err: %v", err)
	}
	templ := template.Must(
		template.ParseFiles("template.html"))
	page := pagemodel.SearchResult{}
	page.Hits = 1000
	page.Start = 20
	page.Items = append(page.Items, model.Post{
		Name: "yinlei",
		Home: "",
		Imgs: nil,
	})
	err = templ.Execute(out, page)
	if err != nil {
		log.Printf("err: %v", err)
	}
}