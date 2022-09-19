package girlygirlpic

import (
	"Tools/BS4/soup"
	"Tools/util/file"
	"log"
)

func FindByKey(url string) {
	links := []string{}
	resp := file.ReadAll("exam.html")
	log.Println(resp)
	doc := soup.HTMLParse(resp)
	mediabody := doc.Find("div", "class", "post-media-body")
	imgs := mediabody.FindAll("img")
	for _, img := range imgs {
		src := img.Attrs()["src"]
		links = append(links, src)
	}
	file.WriteLines("exam.txt", links)
}
