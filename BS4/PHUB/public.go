package PHUB

import (
	"Tools/util/file"
	"io/ioutil"
	"log"
	"strings"
)

const PREFIX = "https://www.pornhub.com"

func writeContent2File(channelname string, lines, links []string) {
	txt := strings.Join([]string{channelname, "txt"}, ".")
	md := strings.Join([]string{channelname, "md"}, ".")
	file.WriteLines(txt, links)
	file.WriteLines(md, lines)
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
