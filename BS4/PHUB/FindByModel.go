package PHUB

import (
	"Tools/BS4/soup"
	"log"
	"os"
	"strings"
)

//适用于https://www.pornhub.com/model
func FindByModel(url string) (ChannelName string, lines, links []string) {
	//links := make([]string, 0)
	//lines := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	// log.Println(string(resp))
	namediv := doc.Find("div", "class", "name")
	ChannelName = namediv.Find("h1").Text()
	ChannelName = strings.Replace(ChannelName, " ", "", -1)
	ChannelName = strings.Replace(ChannelName, "?", "", -1)
	ChannelName = strings.Replace(ChannelName, "\n", "", -1)
	log.Printf("channel name is %v\n", ChannelName)
	lis := doc.FindAll("div", "class", "phimage")
	for _, li := range lis {
		a := li.Find("a")
		href := a.Attrs()["href"]
		Hyperlink := strings.Join([]string{PREFIX, href}, "")
		//log.Printf("href is %v\n", Hyperlink)
		img := a.Find("img")
		src := img.Attrs()["data-thumb_url"]
		//log.Printf("src is %v\n", src)
		title := img.Attrs()["alt"]
		//log.Printf("title is %v\n", title)
		line := strings.Join([]string{"|", title, "|", Hyperlink, "|", "![", title, "](", src, ")|"}, "")
		lines = append(lines, line)
		links = append(links, Hyperlink)
	}
	writeContent2File(ChannelName, lines, links)
	return
}
