package parser

import (
	"fmt"
	"goCrawler/engine"
	"regexp"
	"strconv"
)

const postRe = `<a href="(https://www.zcool.com.cn/work/.*?)"[^>]*>[^<]*<img src="(.*?)" title="(.*?)"[^>]*>[^<]*</a>`
const nextRe = `<button class="active" data-page="(.*?)">[^<]*</button>`

func ParsePost(contents []byte, currentCateIndex string) engine.ParseResult {
	re := regexp.MustCompile(postRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "Post "+string(m[3]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParsePostDetail,
		})
	}

	re = regexp.MustCompile(nextRe)
	matches = re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		currentPageIndex, _ := strconv.Atoi(string(m[1]))
		currentPageIndex++
		if currentPageIndex < 21 {
			result.Requests = append(result.Requests, engine.Request{
				Url: fmt.Sprintf("https://www.zcool.com.cn/discover?cate=1&subCate=%s&hasVideo=0&city=0&college=0&recommendLevel=2&sort=9&page=%s", currentCateIndex, strconv.Itoa(currentPageIndex)),
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParsePost(contents, currentCateIndex)
				},
			})
		}
	}
	return result
}
