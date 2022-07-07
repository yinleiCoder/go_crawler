package fetcher

import (
	"bufio"
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	var bodyReader *bufio.Reader
	if strings.HasPrefix(url, "https://www.zcool.com.cn/work/") {
		reader, err := fetchHtmlByChromedp(url, "body")
		if err != nil {
			return nil, err
		}
		bodyReader = bufio.NewReader(reader)
	} else {
		//resp, err := http.Get("https://www.zcool.com.cn/discover?cate=1&page=1")
		//if err != nil {
		//	panic(err)
		//}
		//defer resp.Body.Close()
		//if resp.StatusCode != http.StatusOK {
		//	fmt.Println("Error: status code", resp.StatusCode)
		//	return
		//}
		//reader, err := fetchHtmlByChromedp(url, "#body .all-work-list  #all-card-content .work-list-box")
		reader, err := fetchHtmlByChromedp(url, "body .workList")
		if err != nil {
			return nil, err
		}
		bodyReader = bufio.NewReader(reader)
	}
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

/**
通过Chromedp获取Html源码
https://github.com/chromedp/examples
*/
func fetchHtmlByChromedp(urlStr, waitVisible string) (r io.Reader, err error) {
	var htmlContent string
	// 自定义User-Agent、禁用图片加载
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // false 可以显示chrome进行调试
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		//chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"),
		chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"),
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
	if err = chromedp.Run(timeoutCtx,
		chromedp.Navigate(urlStr),
		chromedp.WaitVisible(waitVisible),
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
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1*time.Second),
		//chromedp.WaitVisible("#z-pages"),
		//chromedp.SendKeys(`div.spage-number`, kb.End, chromedp.ByQuery),
		//chromedp.Evaluate("document.querySelector('#footer').scrollIntoViewIfNeeded(true)", nil),
		chromedp.OuterHTML(`document.querySelector("html")`, &htmlContent, chromedp.ByJSPath),
	); err != nil {
		log.Fatal("timeout error： ", err)
		return
	}
	r = strings.NewReader(htmlContent)
	return
}

/**
根据网页编码检测并匹配该编码
*/
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//bytes, err := bufio.NewReader(r).Peek(1024)
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
