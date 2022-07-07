package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"goCrawler/engine"
	"goCrawler/model"
)

func ParsePostDetail(contents []byte, postDetailUrl string) engine.ParseResult {
	post := model.Post{}
	reader := bytes.NewReader(contents)
	doc, _ := goquery.NewDocumentFromReader(reader)
	post.Name = doc.Find(".contentTitle").ChildrenFiltered("h1").Text()
	post.Home = postDetailUrl

	doc.Find(".detailContentBox img").Each(func(i int, selection *goquery.Selection) {
		imgSrc, _ := selection.Attr("src")
		post.Imgs = append(post.Imgs, imgSrc)
	})
	result := engine.ParseResult{
		Items: []interface{}{post},
	}
	//log.Printf("%v", post)
	return result
}
