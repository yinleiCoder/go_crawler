package main

import (
	"goCrawler/engine"
	"goCrawler/zcool/parser"
)

/**
Go站酷爬虫
@author yinlei
@date 2022/3/3
*/
func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zcool.com.cn/discover?cate=1&subCate=0&hasVideo=0&city=0&college=0&recommendLevel=2&sort=9&page=1",
		ParserFunc: parser.ParseCateList,
	})
}
