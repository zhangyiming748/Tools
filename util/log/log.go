package log

import (
	"Tools/util/conf"
	"io"
	"log"
	"os"
	"strings"
)

var (
	Info      *log.Logger // 仅打印到屏幕
	CMD       *log.Logger // 仅保存到文件
	Debug     *log.Logger // 打印屏幕并保存到文件
	Warn      *log.Logger // if err handle的问题,打印屏幕并保存到文件
	Emergency *log.Logger // 仅保存关键信息到独立日志文件,用于灾后重建
	MyError   *log.Logger // 未启用
)

func init() {
	function := conf.GetVal("main", "function")
	file := strings.Join([]string{function, "log"}, ".")
	emergence := strings.Join([]string{function, "log"}, ".")
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("打开日志文件错误")
	}

	emergencyf, err := os.OpenFile(emergence, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	ip := strings.Join([]string{function, "INFO :"}, " ")
	dp := strings.Join([]string{function, "DEBUG :"}, " ")
	wp := strings.Join([]string{function, "WARN :"}, " ")
	cp := strings.Join([]string{function, "CMD :"}, " ")
	ep := strings.Join([]string{function, "EMERGENCY :"}, " ")

	Info = log.New(os.Stdout, ip, log.Ltime)
	Debug = log.New(io.MultiWriter(logf, os.Stdout), dp, log.Ltime|log.Lshortfile)
	Warn = log.New(io.MultiWriter(logf, os.Stdout), wp, log.Ltime|log.Lshortfile)
	CMD = log.New(logf, cp, log.Ltime)
	Emergency = log.New(emergencyf, ep, log.Ltime)
	//MyError = log.New(io.MultiWriter(elog, os.Stdout), prefix, log.Ltime|log.Lshortfile)

}
