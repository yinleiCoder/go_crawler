package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"goCrawler/engine"
)

const postRe = `<a href="(https://www.zcool.com.cn/work/.*?)"[^>]*>[^<]*<img src="(.*?)" title="(.*?)"[^>]*>[^<]*</a>`
const nextRe = `<button class="active" data-page="(.*?)">[^<]*</button>`

//func ParsePost(contents []byte, currentCateIndex string) engine.ParseResult {
func ParsePost(contents []byte) engine.ParseResult {
	//re := regexp.MustCompile(postRe)
	//matches := re.FindAllSubmatch(contents, -1)
	//result := engine.ParseResult{}
	//for _, m := range matches {
	//	result.Items = append(result.Items, "Post "+string(m[3]))
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:        string(m[1]),
	//		ParserFunc: ParsePostDetail,
	//	})
	//}
	result := engine.ParseResult{}

	reader := bytes.NewReader(contents)
	doc, _ := goquery.NewDocumentFromReader(reader)
	var postDetailUrls []string
	doc.Find(".workList .cardImg").Each(func(i int, selection *goquery.Selection) {
		targetLink, _ := selection.Find("a").Attr("href")
		postDetailUrls = append(postDetailUrls, targetLink)
	})
	for _, postDetailUrl := range postDetailUrls {
		//result.Items = append(result.Items, "Post " + postDetailUrl)
		result.Requests = append(result.Requests, engine.Request{
			Url: postDetailUrl,
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParsePostDetail(contents, postDetailUrl)
			},
		})
	}

	//re = regexp.MustCompile(nextRe)
	//matches = re.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	currentPageIndex, _ := strconv.Atoi(string(m[1]))
	//	currentPageIndex++
	//	if currentPageIndex < 21 {
	//		result.Requests = append(result.Requests, engine.Request{
	//			Url: fmt.Sprintf("https://www.zcool.com.cn/discover?cate=1&subCate=%s&hasVideo=0&city=0&college=0&recommendLevel=2&sort=9&page=%s", currentCateIndex, strconv.Itoa(currentPageIndex)),
	//			ParserFunc: func(contents []byte) engine.ParseResult {
	//				return ParsePost(contents, currentCateIndex)
	//			},
	//		})
	//	}
	//}
	return result
}
