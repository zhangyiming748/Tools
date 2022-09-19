package rotateVideo

import (
	absfile "Tools/util/file"
	"Tools/util/log"
	"Tools/util/threads"
	"os/exec"
	"strings"
)

func ToRight(src, dst, file string) {
	in := strings.Join([]string{src, file}, "/")
	//todo 未测试
	justname := absfile.ShortNameGetFileName(file)
	justname = strings.Replace(justname, ".", "", -1)
	justname = strings.Replace(justname, "+", "", -1)

	newFileNmae := strings.Join([]string{justname, "mp4"}, ".")

	out := strings.Join([]string{dst, newFileNmae}, "/")
	log.Debug.Printf("src:%s\tfile:%s\nin:%s\tout:%s\n", src, file, in, out)
	t := threads.Threads()
	cmd := exec.Command("ffmpeg", "-threads", t, "-i", in, "-vf", "transpose=1", "-threads", t, out)
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
	log.CMD.Printf("完成当前文件的处理:dst是%s\tfile是%s\n", dst, file)
}
func ToLeft(src, dst, file string) {
	in := strings.Join([]string{src, file}, "/")
	//todo 未测试
	justname := absfile.ShortNameGetFileName(file)
	justname = strings.Replace(justname, ".", "", -1)
	justname = strings.Replace(justname, "+", "", -1)

	newFileNmae := strings.Join([]string{justname, "mp4"}, ".")

	out := strings.Join([]string{dst, newFileNmae}, "/")
	log.Debug.Printf("src:%s\tfile:%s\nin:%s\tout:%s\n", src, file, in, out)
	t := threads.Threads()
	cmd := exec.Command("ffmpeg", "-threads", t, "-i", in, "-vf", "transpose=2", "-threads", t, out)
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
	log.CMD.Printf("完成当前文件的处理:dst是%s\tfile是%s\n", dst, file)
}
