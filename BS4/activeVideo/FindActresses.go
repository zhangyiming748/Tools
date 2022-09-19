package BS4

import (
	"Tools/BS4/soup"
	util "Tools/util/file"
	"github.com/gookit/goutil/dump"
	"log"
	"strings"
)

type Actresses struct {
	Name  string
	Alias string
	Image string
}

func FindActresses(url string) {
	lines := make([]string, 0)
	imgs := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Panicf("err = %v\n", err)
	}
	SaveHtml(resp)
	doc := soup.HTMLParse(resp)
	boxes := doc.Find("div", "class", "actors")
	log.Println(boxes)
	divs := boxes.FindAll("div")
	log.Println(divs)
	for i, div := range divs {
		a := div.Find("a")
		alias := a.Attrs()["title"]
		alias = Replace(alias)

		figure := a.Find("figure", "class", "image")
		img := figure.Find("img", "class", "avatar")
		src := img.Attrs()["src"]

		strong := a.Find("strong")
		name := strong.Text()
		name = Replace(name)

		log.Printf("第%d位老师%s的别名是%s\t对应的头像是%s\n", i+1, name, alias, src)
		var act = &Actresses{
			Name:  name,
			Alias: alias,
			Image: src,
		}
		dump.P(act)
		fname := strings.Join([]string{act.Name, "jpg"}, ".")
		if strings.Contains(src, "unknow") {
			fname = strings.Join([]string{act.Name, "png"}, ".")
			act.Image = "https://javdb.com/assets/actor_unknow-15f7d779b3d93db42c62be9460b45b79e51f8a944796eee30ed87bbb04de0a37.png"
		}
		log.Printf("文件名:%s\n", fname)
		downlink := strings.Join([]string{"wget", act.Image, "-O", fname}, " ")
		imgs = append(imgs, downlink)
		line := strings.Join([]string{"|", act.Name, "|", act.Alias, "|", "![", act.Name, "](", act.Image, ")|"}, "")
		lines = append(lines, line)
	}
	util.WriteLines("actresses.sh", imgs)
	util.WriteLines("actresses.md", lines)
}
