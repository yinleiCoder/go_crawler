package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"goCrawler/engine"
)

/**
https://github.com/PuerkitoBio/goquery
*/
func ParseCateList(contents []byte) engine.ParseResult {
	var cateHrefList []string
	var cateTextList []string
	reader := bytes.NewReader(contents)
	doc, _ := goquery.NewDocumentFromReader(reader)
	doc.Find(".subCateBox").Children().Each(func(i int, selection *goquery.Selection) {
		hrefStr, _ := selection.Attr("href")
		cateHrefList = append(cateHrefList, ZCOOL_PAGE+hrefStr)
		cateTextList = append(cateTextList, selection.Find("h1").Text())
	})

	result := engine.ParseResult{}
	for _, href := range cateHrefList {
		//result.Items = append(result.Items, cateTextList[index])
		result.Requests = append(result.Requests, engine.Request{
			Url: href,
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParsePost(contents, href)
			},
		})
	}
	return result
}
