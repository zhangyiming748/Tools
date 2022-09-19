package convert

import (
	"Tools/util/conf"
	"Tools/util/file"
	"log"
	"os/exec"
	"strings"
)

func WebpTo(src, fp string) {
	in := strings.Join([]string{src, fp}, "/")
	mid := file.LongNameGetFileName(in)
	//mid := strings.Split(in, ".")[0]
	pattern := conf.GetVal("location", "webpto")
	out := strings.Join([]string{mid, pattern}, ".")
	cmd := exec.Command("dwebp", in, "-o", out)
	log.Printf("生成的命令是:%s", cmd)
	//命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Printf("cmd.Run产生的错误:%v", err)
	}
	//fn := split(url)
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Printf("%s", string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Printf("命令运行期间产生的错误")
	}

}
