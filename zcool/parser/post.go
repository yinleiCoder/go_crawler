package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"goCrawler/engine"
	"strconv"
	"strings"
)

func ParsePost(contents []byte, belongOfCateHref string) engine.ParseResult {
	result := engine.ParseResult{}

	reader := bytes.NewReader(contents)
	doc, _ := goquery.NewDocumentFromReader(reader)
	var postDetailUrls []string
	doc.Find(".workList .cardImg").Each(func(i int, selection *goquery.Selection) {
		targetLink, _ := selection.Find("a").Attr("href")
		postDetailUrls = append(postDetailUrls, targetLink)
	})
	for _, postDetailUrl := range postDetailUrls {
		tempPostDetailUrl := postDetailUrl
		//result.Items = append(result.Items, "Post " + postDetailUrl)
		result.Requests = append(result.Requests, engine.Request{
			Url: postDetailUrl,
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParsePostDetail(contents, tempPostDetailUrl)
			},
		})
	}

	belongOfCateHref = belongOfCateHref[:strings.LastIndex(belongOfCateHref, "=")+1]
	maxPaginationNum, _ := strconv.Atoi(doc.Find(".pagination_wrap").ChildrenFiltered("span").Last().Text())
	currentPaginationNum := 2
	for {
		if currentPaginationNum > maxPaginationNum {
			break
		}
		result.Requests = append(result.Requests, engine.Request{
			Url: fmt.Sprintf(belongOfCateHref+"%d", currentPaginationNum),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParsePost(contents, fmt.Sprintf(belongOfCateHref+"%d", currentPaginationNum))
			},
		})
		currentPaginationNum++
	}
	return result
}
