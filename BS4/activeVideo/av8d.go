package BS4

import (
	"Tools/BS4/soup"
	util "Tools/util/file"
	"github.com/gookit/goutil/dump"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	uid   string
	title string
	thumb string
	pic   string
}

func FindByActress(url string) {
	lines := make([]string, 0)
	imgs := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Panicf("err = %v\n", err)
	}
	doc := soup.HTMLParse(resp)
	boxes := doc.FindAll("a", "class", "box")
	for index, box := range boxes {
		log.Printf("正在处理 %d/%d 项\n", index+1, len(boxes))
		if strings.Contains(url, "page") {
			page := strings.Split(url, "=")[1]
			log.Printf("正在处理第%v页\n", page)
		}
		var i Item
		img := box.Find("img").Attrs()["data-src"]
		log.Println(img)
		uid := box.Find("div", "class", "uid").Text()
		log.Println(uid)
		title := box.Find("div", "class", "video-title").Text()
		log.Println(title)
		detiallink := strings.Join([]string{"https://javdb30.com", box.Attrs()["href"]}, "")
		pic := FindDetialMainPic(detiallink)
		log.Println(pic)
		fname := strings.Join([]string{uid, "jpg"}, ".")
		downlink := strings.Join([]string{"wget", pic, "-O", fname}, " ")
		imgs = append(imgs, downlink)
		line := strings.Join([]string{"|", title, "|", uid, "|", "![", uid, "](", img, ")|![", uid, "](", pic, ")|"}, "")
		lines = append(lines, line)
		i.uid = uid
		i.title = title
		i.thumb = img
		i.pic = pic
		dump.P(i)
		log.Printf("处理完成 %d/%d 项\n", index+1, len(boxes))
		time.Sleep(1 * time.Second)
	}
	util.WriteLines("君岛美绪.sh", imgs)
	util.WriteLines("君岛美绪.md", lines)
}
func FindByActress2(url string) {
	lines := make([]string, 0)
	imgs := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Panicf("err = %v\n", err)
	}
	SaveHtml(resp)
	//resp := ReadTemporary()
	doc := soup.HTMLParse(resp)
	actress := doc.Find("title").Text()
	actress = Replace(actress)
	boxes := doc.FindAll("div", "class", "item")
	log.Printf("actress=%v\nboxes=%v\n", actress, boxes)
	for i, box := range boxes {
		a := box.Find("a")
		href := a.Attrs()["href"]
		hyper := strings.Join([]string{"https://javdb.com", href}, "")
		title := a.Attrs()["title"]
		title = Replace(title)
		v_title := a.Find("div", "class", "video-title")
		code := v_title.Find("strong").Text()

		pics := FindDetialMainPics(hyper)
		for i, pic := range pics {
			fname := strings.Join([]string{strconv.Itoa(i + 1), ".", code, ".", "jpg"}, "")
			downlink := strings.Join([]string{"wget", pic, "-O", fname}, " ")
			imgs = append(imgs, downlink)
		}

		thumb := a.Find("img").Attrs()["src"]
		log.Println(thumb)

		log.Printf("%d. %s is %s\tcode is %v\n", i+1, title, hyper, code)
		var finnalPic string
		if len(pics) >= 1 {
			finnalPic = pics[0]
		} else {
			finnalPic = thumb
		}
		var i = &Item{
			uid:   code,
			title: title,
			thumb: thumb,
			pic:   finnalPic,
		}
		dump.P(i)

		one := strings.Join([]string{"|", code, "|", title, "|", "![", code, "](", finnalPic, ")|"}, "")
		log.Println(one)
		lines = append(lines, one)
	}
	util.WriteLines("松島楓.sh", imgs)
	util.WriteLines("松島楓.md", lines)
}
func FindDetialMainPics(url string) []string {
	var links []string
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Panicf("err = %v\n", err)
	}
	//SaveHtml(resp)
	doc := soup.HTMLParse(resp)
	body := doc.Find("div", "class", "message-body")
	pics := body.FindAll("a", "class", "tile-item")
	for i, pic := range pics {
		href := pic.Attrs()["href"]
		links = append(links, href)
		for j := i; j > 0; j-- {
			log.Printf("查询完当前大图,冷却%d秒后进入下一步\n", j)
			time.Sleep(time.Second)
		}
	}
	return links
}
func FindDetialMainPic(url string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("查询大图时产生错误:%v\n", err)
		}
	}()
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	img := doc.Find("img", "class", "video-cover").Attrs()["src"]
	return img
}
func Replace(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "《", "", -1)
	str = strings.Replace(str, "》", "", -1)
	str = strings.Replace(str, "【", "", -1)
	str = strings.Replace(str, "】", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	return str
}
func SaveHtml(s string) {
	content := []byte(s)
	err := ioutil.WriteFile("temporary.html", content, 0766)
	if err != nil {
		log.Panicf("写html文件发生错误:%s\n", err.Error())
	}
}
func ReadTemporary() string {
	filepath := "temporary.html"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panicf("读html缓存文件发生错误:%s\n", err.Error())
	}
	return string(content)
}
