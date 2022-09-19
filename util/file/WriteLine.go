package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func WriteLines(fname string, s []string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
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
func WriteAll(fname, content string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	n, err := f.WriteString(content)
	if err != nil {
		log.Println("写文件出错")
	} else {
		log.Printf("写入%d个字节", n)
	}
}
func ReadAll(fname string) string {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(content)
}
func Writeline(fname, content string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	_, _ = f.WriteString("\n")
	if err != nil {
		log.Println("写文件出错")
	} else {
		//log.Printf("写入%d个字节", n)
	}
}
