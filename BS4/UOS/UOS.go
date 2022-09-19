package UOS

import (
	"Tools/BS4/soup"
	util "Tools/util/file"
	"log"
	"os"
	"strings"
)

func UOS(url string) {
	links := make([]string, 0)
	links = append(links, "|iso|")
	links = append(links, "|:---:|")
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	isos := doc.FindAll("a")
	for _, iso := range isos {
		href := iso.Attrs()["href"]
		link := strings.Join([]string{url, href}, "/")
		if strings.Contains(link, "iso") {
			log.Println(link)
			line := strings.Join([]string{"|[", href, "](", link, ")|"}, "")
			links = append(links, line)
		} else {
			log.Printf("link %v\n", link)
		}
	}
	util.WriteLines("UOS.md", links)
}
