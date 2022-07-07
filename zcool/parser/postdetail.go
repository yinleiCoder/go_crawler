package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"goCrawler/engine"
	"goCrawler/model"
	"regexp"
)

var imgRe = regexp.MustCompile(`<div class="photo-information-content">\s*<img src="(https://img.zcool.cn/community/.*?)"`)
var authorRe = regexp.MustCompile(`<a href="(.*?)" title="(.*?)"\s*class="title-content" target="_blank">[^<]*</a>`)

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
	matches := authorRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		post.Name = string(m[2])
		post.Home = string(m[1])
	}
	matches = imgRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		post.Imgs = append(post.Imgs, string(m[1]))
	}
	result := engine.ParseResult{
		Items: []interface{}{post},
	}
	//log.Printf("%v", post)
	return result
}
