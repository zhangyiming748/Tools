package convert

import (
	"Tools/util/threads"
	"log"
	"os/exec"
	"strings"
)

func Frame(src, dst, file, f string) {
	in := strings.Join([]string{src, file}, "/")
	out := strings.Join([]string{dst, file}, "/")
	log.Printf("src:%s\tfile:%s\nin:%s\tout:%s\n", src, file, in, out)
	t := threads.Threads()
	cmd := exec.Command("ffmpeg", "-threads", t, "-i", in, "-r", f, "-threads", t, out)
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
	log.Printf("当前的dst是%s\tfile是%s\n", dst, file)
	//当前的dst是/Users/zen/Downloads/Downie/done	file是4K【维拉Villa松岛】松岛一期.webm
}
