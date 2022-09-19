package file

import (
	"log"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(fname string) {
	exist, err := PathExists(fname)
	if err != nil {
		log.Printf("get dir error![%v]\n", err)
		return
	}

	if exist {
		log.Printf("has dir![%v]\n", fname)
	} else {
		log.Printf("no dir![%v]\n", fname)
		// 创建文件夹
		err := os.Mkdir(fname, os.ModePerm)
		if err != nil {
			log.Printf("mkdir failed![%v]\n", err)
		} else {
			log.Printf("mkdir success!\n")
		}
	}
}
