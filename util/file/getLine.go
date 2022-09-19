package file

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadLine(src string) []string {
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
	pswds := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		pswds = append(pswds, string(a))
		log.Printf("读取到的密码(%s)\n", string(a))
	}
	return pswds
}

//读取markdown文件获取图片超链接
func Read2Wget(src string) []string {
	fi, err := os.Open(src)
	if err != nil {
		log.Printf("打开md文件失败: %s\n", err)
		return []string{}
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Printf("关闭md文件失败: %s\n", err)
		}
	}()
	ulrs := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := string(a)
		prefix := strings.Split(s, "(")[1]
		suffix := strings.Split(prefix, ")")[0]
		if len(suffix) != 0 {
			ulrs = append(ulrs, suffix)
		}
		fmt.Printf("%s\n", suffix)
		//download.WGet(suffix,"/Users/zen/Github/Tools/test")

	}
	return ulrs
}
