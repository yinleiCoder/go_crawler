package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"goCrawler/engine"
)

const cateListRe = `<span .* data-id="([0-9]+)">[^<]*<span>([^<]+)</span>[^<]*</span>`

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
				return ParsePost(contents)
			},
		})
	}
	//re := regexp.MustCompile(cateListRe)
	//matches := re.FindAllSubmatch(contents, -1)
	//result := engine.ParseResult{}
	//for _, m := range matches {
	//	currentCateIndex := string(m[1])
	//result.Items = append(result.Items, "Cate "+string(m[2]))
	//result.Requests = append(result.Requests, engine.Request{
	//	Url: fmt.Sprintf("https://www.zcool.com.cn/discover?cate=1&subCate=%s&hasVideo=0&city=0&college=0&recommendLevel=2&sort=9&page=1", currentCateIndex),
	//	ParserFunc: func(contents []byte) engine.ParseResult {
	//		return ParsePost(contents, currentCateIndex)
	//	},
	//})
	//fmt.Printf("zcool category url_id: %s, name: %s\n", m[1], m[2])
	//}
	return result
}
