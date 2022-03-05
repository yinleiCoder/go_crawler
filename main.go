package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

/**
Go站酷爬虫
@author yinlei
@date 2022/3/3
*/
func main() {
	//resp, err := http.Get("https://www.zcool.com.cn/discover?cate=1&page=1")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", resp.StatusCode)
	//	return
	//}
	_, reader := fetchHtmlByChromedp("https://www.zcool.com.cn/discover?cate=1&page=1")
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func fetchHtmlByChromedp(urlStr string) (htmlContent string, r io.Reader) {
	// 自定义User-Agent、禁用图片加载
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	chromeCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	if err := chromedp.Run(chromeCtx); err != nil {
		log.Fatal("run error:", err)
	}
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 60*time.Second)
	defer cancel()
	log.Printf("chrome visit page %s \n", urlStr)
	if err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(urlStr),
		chromedp.WaitVisible("#body .all-work-list  #all-card-content .work-list-box"),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		//chromedp.WaitVisible("#z-pages"),
		//chromedp.SendKeys(`div.spage-number`, kb.End, chromedp.ByQuery),
		//chromedp.Evaluate("document.querySelector('#footer').scrollIntoViewIfNeeded(true)", nil),
		chromedp.OuterHTML(`document.querySelector("html")`, &htmlContent, chromedp.ByJSPath),
	); err != nil {
		log.Fatal("timeout error", err)
		return
	}
	r = strings.NewReader(htmlContent)
	return
}

/**
根据网页编码检测并匹配该编码
*/
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
