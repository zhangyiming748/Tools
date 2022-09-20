package XM

import (
	"Tools/util/file"
	"io/ioutil"
	"log"
	"strings"
)

const PREFIX = "https://xhamster.com/"

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
	str = strings.Replace(str, "\u0000", "", -1)
	return str
}
func GetTitle(s string) string {
	//title is ArtofGlossPornVideos:artofgloss.net|xHamster
	if strings.Contains(s, ":") {
		suffix := strings.Split(s, ":")[0]
		if strings.Contains(suffix, "PornVideo") {
			prefix := strings.Split(suffix, "PornVideo")[0]
			return prefix
		}
	}
	return s
}
