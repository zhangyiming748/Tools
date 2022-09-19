package spankbang

import (
	"Tools/BS4/soup"
	"Tools/util/file"
	"log"
	"strings"
	"time"
)

func FindBySearch(url string) {
	links := make([]string, 0)

	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Printf("网络问题%s\n重试", err)
		time.Sleep(3 * time.Second)
		FindByChannel(url)
	}
	doc := soup.HTMLParse(resp)
	div := doc.Find("div", "class", "main_results")
	as := div.FindAll("a", "class", "n")
	for _, a := range as {
		href := a.Attrs()["href"]
		link := strings.Join([]string{"https://spankbang.com", href}, "")
		title := a.Text()
		log.Printf("视频名%s\t视频地址%s\n", title, link)
		links = append(links, link)
	}
	log.Printf("共获取到%d条视频\n", len(links))
	keyword := strings.Split(url, "/s/")[1]
	keyword = strings.Replace(keyword, "/", "", -1)

	fname := strings.Join([]string{keyword, "txt"}, ".")
	file.WriteLines(fname, links)
}
