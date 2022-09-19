package public

import (
	"Tools/util/file"
	"log"
	"strings"
)

func Md2Links(src string) {
	lines := file.ReadLink(src)
	var links []string
	for _, line := range lines {
		if !strings.Contains(line, "|") {
			continue
		}
		values := strings.Split(line, "|")
		for _, value := range values {
			if strings.Contains(value, "https://www.pornhub.com/view_video.php") {
				links = append(links, value)
			}
		}
	}
	log.Printf("提取出的下载链接是: %v", links)

	file.WriteLines("out.txt", links)
}
