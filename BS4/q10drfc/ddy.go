package q10drfc

import (
	BS4 "Tools/BS4/activeVideo"
	"Tools/BS4/soup"
	"Tools/util/file"
	"log"
	"strings"
	"time"
)

func FindByKeyword(url string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("PANIC! 重试")
			time.Sleep(3 * time.Second)
			FindByKeyword(url)
		}
	}()
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Printf("网络问题%s\n重试", err)
		time.Sleep(3 * time.Second)
		FindByKeyword(url)
	}
	doc := soup.HTMLParse(resp)
	//fmt.Printf("%v,%v", doc, resp)

	title := doc.Find("title").Text()
	title = BS4.Replace(title)
	log.Printf("当前下载的是:%s\n", title)
	if strings.Contains(title, "-") {
		title = strings.Split(title, "-")[0]
	}
	p := doc.Find("p")
	srcs := make([]string, 0)
	imgs := p.FindAll("img")
	for _, img := range imgs {
		src := img.Attrs()["data-original"]
		srcs = append(srcs, src)
		log.Printf("解析到的图片链接:%s\n", src)
	}
	file.WriteLines(title, srcs)
}
func FindBySearch(url string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("PANIC! 重试")
			time.Sleep(3 * time.Second)
			FindBySearch(url)
		}
	}()
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Printf("网络问题%s\n重试", err)
		time.Sleep(3 * time.Second)
		FindBySearch(url)
	}
	doc := soup.HTMLParse(resp)
	//fmt.Printf("%v,%v", doc, resp)

	title := doc.Find("h1", "class", "entry-title").Text()
	title = BS4.Replace(title)
	log.Printf("当前下载的是:%s\n", title)
	if strings.Contains(title, "-") {
		title = strings.Split(title, "-")[0]
	}

	//boxes := doc.Find("div", "class", "entry-content clearfix")
	as := doc.FindAll("a", "title", "点击图片查看下一张")

	srcs := make([]string, 0)
	for _, a := range as {
		img := a.Find("img")
		src := img.Attrs()["data-original"]
		srcs = append(srcs, src)
		log.Printf("解析到的图片链接:%s\n", src)
	}
	file.WriteLines(title, srcs)
}
