package file

import (
	"path"
	"strings"
)

// 短文件名只获取文件名本名
func ShortNameGetFileName(fname string) string {
	ext := path.Ext(fname)
	justname := strings.Trim(fname, ext)
	return justname
}

// 短文件名只获取文件扩展名
func ShortNameGetExtNmae(fname string) string {
	dot := path.Ext(fname)
	ext := strings.Trim(dot, ".")
	return ext
}

// 文件绝对路径获取长文件名本名
func LongNameGetFileName(fname string) string {
	ext := path.Ext(fname)
	longname := strings.Replace(fname, ext, "", 1)
	return longname
}

// 文件绝对路径获取文件扩展名
func LongNameGetExtName(fname string) string {
	dot := path.Ext(fname)
	ext := strings.Trim(dot, ".")
	return ext
}
