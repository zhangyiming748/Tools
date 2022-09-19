package file

import "strings"

//拆分文件名和扩展名,当文件名中包含.
func SplitFilename(s string) (file, ext, newfileName string) {
	if strings.Contains(s, ".") {
		words := strings.Split(s, ".")
		ext = words[len(words)-1]
	}
	del := strings.Join([]string{".", ext}, "")
	file = strings.ReplaceAll(s, del, "")
	newfileName = strings.ReplaceAll(file, ".", "")
	return file, ext, newfileName
}
