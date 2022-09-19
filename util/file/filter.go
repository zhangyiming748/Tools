package file

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func Filter(src, keyword string) {
	before := readInLine(src)
	after := make([]string, 0)
	var count int
	for _, lines := range before {
		attrs := strings.Split(lines, " ")
		for _, line := range attrs {
			if !strings.Contains(line, "href") {
				continue
			}
			if strings.Contains(line, keyword) {
				log.Println(line)
				log.Printf("替换之前%s\n", line)
				line = procrssGirlygirlpic(line)
				log.Printf("替换之后%s\n", line)
				after = append(after, line)
				log.Println(line)
				count++
			}
		}
	}
	s := strings.Split(src, ".")
	fname := s[0]
	extname := s[1]
	newname := strings.Join([]string{fname, "after.", extname}, "")

	writeInLine(newname, after)
	dup(newname, "final.txt")
}

//按行读文件
func readInLine(src string) []string {
	fi, err := os.Open(src)
	if err != nil {
		log.Printf("打开文件失败: %s\n", err)
		return []string{}
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Printf("关闭文件失败: %s\n", err)
		}
	}()
	links := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		links = append(links, string(a))
		//log.Printf("读取到的行(%s)\n", string(a))
	}
	return links
}

//写入符合过滤条件的行
func writeInLine(dst string, s []string) {
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0776)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for _, v := range s {
		_, err := f.WriteString(strings.Join([]string{v, "\n"}, ""))
		if err != nil {
			return
		}
	}
}
func procrssGirlygirlpic(s string) string {
	log.Printf("替换函数接收到的%s\n", s)
	s = strings.Replace(s, "data-src", "", -1)
	s = strings.Replace(s, "href", "", -1)
	s = strings.Replace(s, "src", "", -1)

	s = strings.Replace(s, "=", "", -1)
	s = strings.Replace(s, "\"", "", -1)
	log.Printf("替换函数返回的%s\n", s)
	return s
}
func dup(src, dst string) {
	var passwd = map[string]bool{}
	for _, v := range ReadLine(src) {
		if _, ok := passwd[v]; ok {
			continue
		} else {
			passwd[v] = true
		}
	}

	after := make([]string, 0)
	for k, _ := range passwd {
		after = append(after, k)
	}

	writeInLine(dst, after)
}
