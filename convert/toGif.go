package convert

import (
	"Tools/util/threads"
	"log"
	"os/exec"
	"strings"
)

func ToGif(src, file string) {
	t := threads.Threads()
	in := strings.Join([]string{src, file}, "/")
	fname := strings.Split(file, ".")[0]
	fname = strings.Join([]string{fname, "gif"}, ".")
	out := strings.Join([]string{src, fname}, "/")
	cmd := exec.Command("ffmpeg", "-threads", t, "-i", in, "-threads", t, out)
	log.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Println("命令执行中有错误产生", err)
	}
}
