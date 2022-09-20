package main

import (
	"Tools/convert"
	"Tools/file/HASH"
	"Tools/file/download"
	"Tools/file/format"
	"Tools/file/mediaInfo"
	"Tools/getTime"
	"Tools/net"
	"Tools/unzip"
	"Tools/util/conf"
	util "Tools/util/file"
	"Tools/util/log"
	"Tools/util/threads"
	"Tools/weather"
	"github.com/zhangyiming748/AVmerger/merge"
	"github.com/zhangyiming748/rotateVideo/rotate"
	conv "github.com/zhangyiming748/video2h265mp4"
	"github.com/zhangyiming748/youtube-dl-bat/ytd"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	files        []string
	pattern      = conf.GetVal("location", "pattern")
	src          = conf.GetVal("location", "src")
	dst          = conf.GetVal("location", "dst")
	delDone      = conf.GetVal("main", "delAfterDone")
	MaxGoroutine = conf.GetVal("youtube-dl", "goroutine")
	fp           = conf.GetVal("youtube-dl", "fp")
	addr         = conf.GetVal("youtube-dl", "addr")
	port         = conf.GetVal("youtube-dl", "port")
	target       = conf.GetVal("youtube-dl", "target")
	proxy        = conf.GetVal("youtube-dl", "isproxy")
	URL          = conf.GetVal("hey", "url")
	Requests     = conf.GetVal("hey", "Requests")
	Concurrent   = conf.GetVal("hey", "Concurrent")
	videoPrefix  = conf.GetVal("title", "start")
	multiLinks   = conf.GetVal("download", "multi")
)

func init() {
	// initialization.Initialization()
	//log.SetFlags(log.Lshortfile)
	defer func() {
		if err := recover(); err != nil {
			log.Warn.Println("我又救了你一命")
		}
	}()
	if isIllegal(src, dst) {
		log.Warn.Panicf("参数不合法:src=%v,dst=%v\n", src, dst)
	}
	cmd := exec.Command("bash", "-c", "init.sh")
	cmd.Run()

}
func main() {
	start := time.Now()
	log.Debug.Println("程序开始时间:", time.Now().Format("2006-01-02 15:04:05"))
	defer func() {
		end := time.Now()
		log.Debug.Println("程序结束时间:", time.Now().Format("2006-01-02 15:04:05"))
		sub := end.Sub(start)
		log.Debug.Println("程序用时:", sub)
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Warn.Printf("程序运行过程中有错误产生:%v", err)
		}
	}()
	files = util.GetFiles(src, pattern)
	fn := conf.GetVal("main", "function")
	switch fn {
	case "MediaInfo":
		if exists("Solution.sh") {
			if err := os.Remove("Solution.sh"); err != nil {
				log.Warn.Panicln("删除上次生成的脚本文件失败")
			}
			log.Debug.Println("删除上次生成的脚本文件")
		}
		if exists("report.md") {
			if err := os.Remove("report.md"); err != nil {
				log.Warn.Panicln("删除上次生成的报告文件失败")
			}
			log.Debug.Println("删除上次生成的报告文件")
		}
		title := strings.Join([]string{"|", "文件名", "|", "isHEVC", "|"}, "")
		form := strings.Join([]string{"|:---:|:---:|"}, "")
		util.Writeline("report.md", title)
		util.Writeline("report.md", form)
		for _, f := range files {
			//log.Printf("需要测试的文件:%s\n", f)
			mediaInfo.MediaInfo(src, f)
		}
	case "ToMp4":
		for index, file := range files {
			log.Debug.Printf("准备好进行转换的文件:%v", file)
			convert.ToMp4(src, file, index, len(files))
			if delDone == "true" {
				s := strings.Join([]string{src, file}, "/")
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Warn.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	case "ToWebm":
		for _, f := range files {
			log.Debug.Printf("准备好进行转换的文件:%v", f)
			convert.ToWebm(src, f)
			if delDone == "true" {
				s := strings.Join([]string{src, f}, "/")
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Warn.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	case "frame":
		f := conf.GetVal("frame", "fps")
		for _, file := range files {
			log.Debug.Printf("当前正在处理的文件:%v", f)
			convert.Frame(src, dst, file, f)
			if delDone == "true" {
				s := strings.Join([]string{src, file}, "/")
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Info.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	case "resolution":
		p := conf.GetVal("resolution", "p")
		for _, file := range files {
			log.Debug.Printf("当前正在处理的文件:%v", file)
			convert.Resolution(src, dst, file, p)
			if delDone == "true" {
				s := strings.Join([]string{src, file}, "/")
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Warn.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	//case "ToH265":
	//	log.Emergency.Println("fn")
	//	for index, file := range files {
	//		log.Debug.Printf("准备好进行转换的文件:%v", file)
	//		convert.ToH265(src, dst, file, index, len(files))
	//		if delDone == "true" {
	//			log.Info.Printf("删除源文件:%s\n", file)
	//			fp := strings.Join([]string{src, file}, "/")
	//			if err := os.Remove(fp); err == nil {
	//				log.Info.Printf("删除转换完成的文件\"%v\"成功\n", fp)
	//			} else {
	//				log.Warn.Printf("删除转换完成的文件\"%v\"发生错误\n", fp)
	//			}
	//		} else {
	//			log.Info.Printf("保留源文件:%s\n", file)
	//		}
	//		runtime.GC()
	//	}
	//case "rotate":
	//	if direct := conf.GetVal("rotate", "direction"); direct == "ToRight" {
	//		for _, f := range files {
	//			log.Debug.Printf("准备好进行转换的文件:%v", f)
	//			rotateVideo.ToRight(src, dst, f)
	//			if delDone == "true" {
	//				s := strings.Join([]string{src, f}, "/")
	//				if err := os.Remove(s); err == nil {
	//					log.Info.Printf("删除转换完成的文件%v成功\n", s)
	//				} else {
	//					log.Warn.Printf("删除转换完成的文件%v发生错误\n", s)
	//				}
	//			}
	//		}
	//	} else {
	//		for _, f := range files {
	//			log.Debug.Printf("准备好进行转换的文件:%v", f)
	//			rotateVideo.ToLeft(src, dst, f)
	//			if delDone == "true" {
	//				s := strings.Join([]string{src, f}, "/")
	//				if err := os.Remove(s); err == nil {
	//					log.Info.Printf("删除转换完成的文件%v成功\n", s)
	//				} else {
	//					log.Warn.Printf("删除转换完成的文件%v发生错误\n", s)
	//				}
	//			}
	//		}
	//	}
	case "ToWebp":
		log.Info.Println("webp最大宽高不得超过16383像素")
		for _, f := range files {
			convert.ToWebp(src, f)
		}
	case "WebpTo":
		for _, f := range files {
			convert.WebpTo(src, f)
		}
	case "ToGif":
		for _, f := range files {
			convert.ToGif(src, f)
		}
	case "Unzip":
		keyfile := conf.GetVal("unzip", "passwd")
		crack := make(chan string, 1)
		passwords := util.ReadLine(keyfile)
		go func() {
			for _, passwd := range passwords {
				unzip.UnZip(src, dst, passwd, crack)
			}
		}()
		if v, ok := <-crack; ok {
			log.Debug.Printf("密码有可能是%v", v)
			break
		}
	case "wget":
		list := conf.GetVal("download", "wget")
		files := util.ReadLink(list)
		for _, f := range files {
			download.WGet(f, dst)
		}
	case "Weather":
		weather.Weather()
	case "ExtractAudio":
		files := util.GetAllFiles(src)
		for i, f := range files {
			convert.ExtractAudio(src, f)
			log.Info.Printf("处理第(%d / %d)个文件\n", i+1, len(files))
			if delDone == "true" {
				s := strings.Join([]string{src, f}, "/")
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Warn.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	case "GetTime":
		getTime.GetTime()
	case "youtube-dl":
		ytd.Master(fp, addr, port, target, MaxGoroutine, proxy)
	case "hey":
		net.Hey(URL, Requests, Concurrent)
	case "WeatherPNG":
	case "HASH":
		if isDir(src) {
			files = util.GetFiles(src, pattern)
			HASH.SHA1(files...)
			HASH.SHA256(files...)
			HASH.MD5(files...)
		} else {
			if sameFile(src, dst) {
				log.Debug.Printf("%v和%v是同一个文件\n", src, dst)
			} else {
				log.Debug.Printf("%v和%v不是同一个文件\n", src, dst)
			}
		}
	case "duplicate":
		util.Duplicate(src, dst)
	case "decode":
		files = util.GetFiles(src, pattern)
		for _, f := range files {
			fullpath := strings.Join([]string{src, f}, "/")
			log.Debug.Printf("文件路径为%v\n", f)
			format.GB18030ToUtf8(fullpath)
		}
	case "cut":
		files = util.GetFiles(src, pattern)
		for _, f := range files {
			convert.Cut(src, f, videoPrefix)
		}
	case "multi":
		links := util.ReadLink(multiLinks)
		for _, link := range links {
			url := strings.Split(link, "|")[0]
			name := strings.Split(link, "|")[1]
			log.Debug.Printf("获取到的下载链接:%s\n保存的文件名是%s\n", url, name)
			err := download.MultiDownload(url, name, runtime.NumCPU(), true)
			if err != nil {
				continue
			}
		}
	case "ToFlac":
		files := util.GetFiles(src, pattern)
		for _, f := range files {
			convert.ToFlac(src, f)
		}
	case "ToMp3":
		files := util.GetFiles(src, pattern)
		for _, f := range files {
			convert.ToMp3(src, f)
		}
	case "bilibili":
		way := conf.GetVal("bilibili", "way")
		switch way {
		case "Single":
			merge.Single(src, dst)
		case "Multi":
			merge.Multi(src, dst)
		default:
			log.Info.Println("没有选择视频类型")
			return
		}
	case "ToH265":
		conv.ConvToH265(src, dst, pattern, threads.Threads())
	case "Rotate":
		direction := conf.GetVal("rotate", "direction")
		rotate.Rotate(src, pattern, direction, dst, threads.Threads())
	//case "amr":
	//	ConvertMP3(src, dst)
	case "default":
		log.Info.Println("不运行功能")
	case "panic":
		log.Info.Panicln("手动panic")
	case "printLog":
		log.Info.Println("仅打印日志")
		log.CMD.Println("仅保存到文件")
		log.Debug.Println("打印并保存")
		log.Warn.Println("发生了错误")
	default:
		log.Warn.Panicln("没有指定程序功能或错误的拼写")
	}
}

// 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func sameFile(src, dst string) bool {
	s1f1, _ := HASH.SHA1File(src)
	s256f1, _ := HASH.SHA256File(src)
	md5f1, _ := HASH.MD5File(src)
	s1f2, _ := HASH.SHA1File(dst)
	s256f2, _ := HASH.SHA256File(dst)
	md5f2, _ := HASH.MD5File(dst)
	if s1f1 == s1f2 && s256f1 == s256f2 && md5f1 == md5f2 {
		return true
	}
	return false
}

// 验证基本参数是否合法
func isIllegal(src, dst string) bool {
	if src == dst {
		log.Warn.Println("输入输出目录相同\n")
		return true
	}
	if !exists(src) {
		log.Warn.Println("src目录不存在\n")
		return true
	}
	if !exists(dst) {
		log.Warn.Println("dst目录不存在\n")
		return true
	}
	if !isDir(src) {
		log.Warn.Println("src不是目录\n")
		return true
	}
	if !isDir(dst) {
		log.Warn.Println("dst不是目录\n")
		return true
	}
	return false
}
