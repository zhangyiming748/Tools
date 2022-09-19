package rar

import (
	"fmt"
	"log"
	"os/exec"
)

func UnRar(rarpath string, pass string) {
	cmd := exec.Command("unrar", "e", "-p"+pass, rarpath) //解压到当前文件夹
	defer func() {
		if err := recover(); err != nil {
			log.Printf("运行中产生错误%v", err)
		}
	}()
	if err := cmd.Run(); err != nil {
		if err.Error() == "exit status 255" {
			fmt.Printf("file : %s\tpassword is :%s\n", rarpath, pass)
		}
		//log.Printf("runcmd err:%v\n", err)
	}
}
