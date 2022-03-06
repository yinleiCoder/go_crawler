package parser

import (
	"goCrawler/engine"
	"regexp"
)

const postRe = `<a href="(https://www.zcool.com.cn/work/.*?)"[^>]*>[^<]*<img src="(.*)" title="(.*)"[^>]*>[^<]*</a>`

func ParsePost(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(postRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "Post "+string(m[3]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParsePostDetail,
		})
	}
	return result
}
