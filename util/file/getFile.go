package file

import (
	"Tools/util/log"
	"io/ioutil"
	"strings"
)

// 输入目录路径和扩展名,返回符合的相对路径文件名切片列表
// todo 改进:使用后缀判断,避免文件名中间含有半角句号
func GetFiles(dir, pattern string) []string {
	files, _ := ioutil.ReadDir(dir)
	var aim []string
	types := strings.Split(pattern, ";") //"wmv;rm"
	for _, f := range files {
		//fmt.Println(f.Name())
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			//log.Info.Printf("有效的文件:%v\n", f.Name())
			for _, v := range types {
				if strings.HasSuffix(f.Name(), v) {
					log.Debug.Printf("有效的目标文件:%v\n", f.Name())
					//absPath := strings.Join([]string{dir, f.Name()}, "/")
					//log.Printf("目标文件的绝对路径:%v\n", absPath)
					aim = append(aim, f.Name())
				}
			}
		}
	}
	return aim
}
func GetAllFiles(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	var aim []string
	for _, f := range files {
		//fmt.Println(f.Name())
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			//log.Printf("有效的文件:%v\n", f.Name())
			//if strings.HasSuffix(f.Name(), pattern) {
			//	log.Printf("有效的目标文件:%v\n", f.Name())
			//	//absPath := strings.Join([]string{dir, f.Name()}, "/")
			//	//log.Printf("目标文件的绝对路径:%v\n", absPath)
			aim = append(aim, f.Name())
			//}
		}
	}
	return aim
}
