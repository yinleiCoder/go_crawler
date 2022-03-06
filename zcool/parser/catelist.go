package parser

import (
	"fmt"
	"goCrawler/engine"
	"regexp"
)

const cateListRe = `<span .* data-id="([0-9]+)">[^<]*<span>([^<]+)</span>[^<]*</span>`

func ParseCateList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cateListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "Cate "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        fmt.Sprintf("https://www.zcool.com.cn/discover?cate=1&subCate=%s&hasVideo=0&city=0&college=0&recommendLevel=2&sort=9&page=1", m[1]),
			ParserFunc: ParsePost,
		})
		//fmt.Printf("zcool category url_id: %s, name: %s\n", m[1], m[2])
	}
	return result
}
