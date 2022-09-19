package spankbang

import (
	"Tools/BS4/soup"
	"Tools/util/file"
	"log"
	"strings"
	"time"
)

func FindByChannel(url string) {
	links := make([]string, 0)

	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Printf("网络问题%s\n重试", err)
		time.Sleep(3 * time.Second)
		FindByChannel(url)
	}
	doc := soup.HTMLParse(resp)
	//fmt.Printf("%v\n%v\n", doc, resp)
	body := doc.Find("main", "id", "container")
	div := body.Find("div", "id", "browse_new")
	//fmt.Println(body, div)
	top := div.Find("ul", "class", "top")
	t_pic := top.Find("li", "class", "p")
	t_img := t_pic.Find("img")
	t_src := t_img.Attrs()["src"]
	channel_icon := strings.Join([]string{"https:", t_src}, "")
	//fmt.Println(channel_icon)
	t_name := top.Find("li", "class", "i")
	t_h1 := t_name.Find("h1")
	channel_name := t_h1.Find("em").Text()
	log.Printf("图片地址%s\n频道名%s\n", channel_icon, channel_name)

	as := div.FindAll("a", "class", "n")
	for _, a := range as {
		href := a.Attrs()["href"]
		link := strings.Join([]string{"https://spankbang.com", href}, "")
		title := a.Text()
		log.Printf("视频名%s\t视频地址%s\n", title, link)
		links = append(links, link)
	}
	log.Printf("共获取到%d条视频\n", len(links))
	fname := strings.Join([]string{channel_name, "txt"}, ".")
	file.WriteLines(fname, links)
}
