package parser

import (
	"goCrawler/engine"
	"goCrawler/model"
	"log"
	"regexp"
)

var imgRe = regexp.MustCompile(`<div class="photo-information-content">\s*<img src="(https://img.zcool.cn/community/.*?)"`)
var authorRe = regexp.MustCompile(`<a href="(https://www.zcool.com.cn/u/.*?)" title="(.*?)"\s*class="title-content" target="_blank">[^<]*</a>`)

func ParsePostDetail(contents []byte) engine.ParseResult {
	post := model.Post{}
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
	log.Printf("%v", post)
	return result
}
