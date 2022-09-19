package format

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

func isUTF8(src string) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	file, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err)
	}
	return utf8.Valid(file)
}

func readGB18030(src string) string {
	file, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err)
	}

	decoder := mahonia.NewDecoder("gb18030")
	if decoder == nil {
		fmt.Println("编码不存在")
	}
	return decoder.ConvertString(string(file))
}
func WriteUTF8(dst, s string) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	writeString, err := f.WriteString(s)
	if err != nil {
		panic(err)
	}
	return writeString
}

func GB18030ToUtf8(src string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("转换编码产生错误:%v\n", err)
		}
	}()
	if isUTF8(src) {
		log.Printf("跳过编码已经是UTF8的文件:%s\n", src)
	} else {
		dst := readGB18030(src)
		err := os.Remove(src)
		if err != nil {
			log.Printf("删除旧文件编码出现问题,出现问题的文件是%v\n", src)
		}
		nums := WriteUTF8(src, dst)
		log.Printf("文件\"%s\"写入%d个字符\n", src, nums)
	}
}
