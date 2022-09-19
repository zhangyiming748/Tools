package PHUB

import (
	"Tools/BS4/soup"
	"log"
	"os"
	"strings"
)

//适用于https://www.pornhub.com/channels
func FindByChannel(url string) (ChannelName string, lines, links []string) {
	//links := make([]string, 0)
	//lines := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Panicf("获取网页出现错误:%s\n", err.Error())
	}
	doc := soup.HTMLParse(resp)
	// log.Println(string(resp))
	ChannelName = doc.Find("h1").Text()
	log.Printf("channel name is %v\n", ChannelName)
	lis := doc.FindAll("div", "class", "phimage")
	for _, li := range lis {
		a := li.Find("a")
		href := a.Attrs()["href"]
		Hyperlink := strings.Join([]string{PREFIX, href}, "")
		log.Printf("href is %v\n", Hyperlink)
		img := a.Find("img")
		src := img.Attrs()["data-thumb_url"]
		log.Printf("src is %v\n", src)
		title := img.Attrs()["alt"]
		log.Printf("title is %v\n", title)
		line := strings.Join([]string{"|", title, "|", Hyperlink, "|", "![", title, "](", src, ")|"}, "")
		lines = append(lines, line)
		links = append(links, Hyperlink)
	}
	writeContent2File(ChannelName, lines, links)
	return
}
func FindByChannelPlus(url string) (ChannelName string, lines, links []string) {
	//links := make([]string, 0)
	//lines := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	// log.Println(string(resp))
	ChannelName = doc.Find("h1").Text()
	log.Printf("channel name is %v\n", ChannelName)
	items := doc.FindAll("a")
	for _, item := range items {
		if href := item.Attrs()["href"]; strings.Contains(href, "/view_video.php?viewkey") {
			title := item.Attrs()["data-title"]
			Hyperlink := strings.Join([]string{PREFIX, href}, "")
			log.Printf("href is %v\n", Hyperlink)
			log.Printf("title is %v\n", title)
			line := strings.Join([]string{"|", title, "|", Hyperlink, "|", "![", title, "](", ")|"}, "")
			lines = append(lines, line)
			links = append(links, Hyperlink)
		}
	}
	writeContent2File(ChannelName, lines, links)
	return
}
func FindByChannelAfter2(url string) {
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		os.Exit(1)
	}
	var lines []string
	var links []string
	doc := soup.HTMLParse(resp)
	ChannelName := doc.Find("h1").Text()
	log.Printf("channel name is %v\n", ChannelName)
	ul := doc.Find("ul", "id", "showAllChanelVideos")
	lis := ul.FindAll("li")
	for _, li := range lis {
		wrap := li.Find("div", "class", "wrap")
		log.Println(wrap)
		phimage := wrap.Find("div", "class", "phimage")
		log.Println(phimage)
		a := phimage.Find("a")
		href := a.Attrs()["href"]
		url := strings.Join([]string{"https://www.pornhub.com", href}, "")
		span := li.Find("span", "class", "title")
		title := span.Find("a").Attrs()["title"]
		links = append(links, url)
		log.Printf("获取到的url:%s\n获取到的标题:%s\n", url, title)
		one := strings.Join([]string{"|", title, "|", url, "|"}, "")
		lines = append(lines, one)
	}
	writeContent2File(ChannelName, lines, links)
}
