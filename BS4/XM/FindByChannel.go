package XM

import (
	"Tools/BS4/soup"
	util "Tools/util/file"
	"fmt"
	"log"
	"strings"
)

// 适用于https://xhamster.com/channels/${channleName}
func FindByChannel(url string) {
	//thumb-list__item video-thumb role-pop impression-thumb viewed-thumb
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Panicf("获取网页出现错误:%s\n", err.Error())
	}
	//SaveHtml(resp)
	var (
		lines []string
		links []string
	)

	//resp := ReadTemporary()
	doc := soup.HTMLParse(resp)
	title := doc.Find("title").Text()
	title = Replace(title)
	title = GetTitle(title)
	log.Printf("title is %v", title)
	thumbLinks := doc.FindAll("a", "data-role", "thumb-link")
	duplicate := make(map[string]bool)
	for _, thumbLink := range thumbLinks {
		href := thumbLink.Attrs()["href"]
		img := thumbLink.Attrs()["data-sprite"]
		view := thumbLink.Attrs()["data-previewvideo"]
		if _, ok := duplicate[href]; ok {
			log.Printf("跳过重复的项目%s\n", href)
			continue
		} else {
			duplicate[href] = true
		}

		links = append(links, href)

		video := fmt.Sprintf("<video id=\"%s\"> controls=\"\" preload=\"none\" poster=\"%s\"><source id=\"webm\" src=\"webm格式视频\" type=\"video/webm\"></videos>", href, view)
		log.Println(video)

		iframe := fmt.Sprintf("<iframe src=\"%s\" scrolling=\"no\" border=\"0\" frameborder=\"no\" framespacing=\"0\" allowfullscreen=\"true\" height=600 width=800> </iframe>", view)
		/*
			<iframe src="视频或者网页路径" scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true" height=600 width=800> </iframe>
		*/
		log.Println(iframe)

		one := strings.Join([]string{"|", href, "|![webp](", img, ")|"}, "")
		lines = append(lines, one)
		log.Printf("href = %v\nimg = %v\n view = %v\n", href, img, view)
	}
	txtName := strings.Join([]string{title, "txt"}, ".")
	markdownName := strings.Join([]string{title, "md"}, ".")
	util.WriteLines(txtName, links)
	util.WriteLines(markdownName, lines)
}
