package convert

import (
	absfile "Tools/util/file"
	"Tools/util/log"
	"Tools/util/threads"
	"os/exec"
	"strings"
)

func ToH265(src, dst, file string, index, total int) {

	in := strings.Join([]string{src, file}, "/")
	log.Emergency.Printf("开始处理文件:%v", in)
	justname := absfile.ShortNameGetFileName(file)
	justname = strings.Replace(justname, "+", "", -1)

	newFilename := strings.Join([]string{justname, "mp4"}, ".")
	out := strings.Join([]string{dst, newFilename}, "/")

	log.Info.Printf("src:%s\tfile:%s\nin:%s\tout:%s\n", src, file, in, out)
	t := threads.Threads()
	cmd := exec.Command("ffmpeg", "-threads", t, "-i", in, "-c:v", "libx265", "-threads", t, out)
	log.CMD.Printf("开始处理文件%s\t生成的命令是:%s", file, cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		log.Info.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Warn.Println("命令执行中有错误产生", err)
	}
	log.Emergency.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
}
