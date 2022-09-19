package PHUB

import (
	"Tools/BS4/soup"
	"os"
	"strings"
)

//适用于https://www.pornhub.com/pornstar
func FindByStarUpload(url string) (StarName string, lines, links []string) {
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	StarName = doc.Find("div", "class", "name").Find("h1").Text()
	StarName = strings.Replace(StarName, "\t", "", -1)
	StarName = strings.Replace(StarName, "\n", "", -1)
	items := doc.FindAll("div", "class", "phimage")
	for _, item := range items {
		a := item.Find("a")
		href := a.Attrs()["href"]
		Hyperlink := strings.Join([]string{PREFIX, href}, "")
		title := a.Attrs()["title"]
		img := a.Find("img")
		src := img.Attrs()["data-thumb_url"]
		line := strings.Join([]string{"|", title, "|", Hyperlink, "|", "![", title, "](", src, ")|"}, "")
		lines = append(lines, line)
		links = append(links, Hyperlink)
	}
	writeContent2File(StarName, lines, links)
	return
}
