package composite

import (
	"Tools/util/threads"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

/*
当单集多P和多单集混在一起
按照多单集转换
异常文件名使用
find . -name "*.json" | xargs grep "MV" | tee find.txt
查找
重新使用多P命令转换
*/
type info struct {
	video string
	audio string
	title string
}
type entry struct {
	MediaType                  int    `json:"media_type"`
	HasDashAudio               bool   `json:"has_dash_audio"`
	IsCompleted                bool   `json:"is_completed"`
	TotalBytes                 int    `json:"total_bytes"`
	DownloadedBytes            int    `json:"downloaded_bytes"`
	Title                      string `json:"title"`
	TypeTag                    string `json:"type_tag"`
	Cover                      string `json:"cover"`
	VideoQuality               int    `json:"video_quality"`
	PreferedVideoQuality       int    `json:"prefered_video_quality"`
	GuessedTotalBytes          int    `json:"guessed_total_bytes"`
	TotalTimeMilli             int    `json:"total_time_milli"`
	DanmakuCount               int    `json:"danmaku_count"`
	TimeUpdateStamp            int64  `json:"time_update_stamp"`
	TimeCreateStamp            int64  `json:"time_create_stamp"`
	CanPlayInAdvance           bool   `json:"can_play_in_advance"`
	InterruptTransformTempFile bool   `json:"interrupt_transform_temp_file"`
	QualityPithyDescription    string `json:"quality_pithy_description"`
	QualitySuperscript         string `json:"quality_superscript"`
	CacheVersionCode           int    `json:"cache_version_code"`
	PreferredAudioQuality      int    `json:"preferred_audio_quality"`
	AudioQuality               int    `json:"audio_quality"`
	Avid                       int    `json:"avid"`
	Spid                       int    `json:"spid"`
	SeasionId                  int    `json:"seasion_id"`
	Bvid                       string `json:"bvid"`
	OwnerId                    int    `json:"owner_id"`
	OwnerName                  string `json:"owner_name"`
	OwnerAvatar                string `json:"owner_avatar"`
	PageData                   struct {
		Cid              int    `json:"cid"`
		Page             int    `json:"page"`
		From             string `json:"from"`
		Part             string `json:"part"`
		Link             string `json:"link"`
		RichVid          string `json:"rich_vid"`
		Vid              string `json:"vid"`
		HasAlias         bool   `json:"has_alias"`
		Weblink          string `json:"weblink"`
		Offsite          string `json:"offsite"`
		Tid              int    `json:"tid"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		Rotate           int    `json:"rotate"`
		DownloadTitle    string `json:"download_title"`
		DownloadSubtitle string `json:"download_subtitle"`
	} `json:"page_data"`
}

// 批量转换Android端哔哩哔哩下载文件

//适用于许多单集
func ForAllSingle(dir string) {
	var infos []info
	up := getDir(dir)
	//log.Printf("单集根目录是%v\n", up)
	middle := []string{}
	//end := []string{}
	//bottom := []string{}
	for _, v := range up {
		if strings.Contains(v, ".") {
			continue
		}
		if strings.Contains(v, "DS") {
			continue
		}
		m := strings.Join([]string{dir, v}, "/")
		middle = append(middle, m)
	}
	for _, m := range middle {
		var i info
		//log.Printf("拼接一级目录后的路径:%v\n", m)
		e := getDir(m)
		end := strings.Join([]string{m, e[0]}, "/")
		//log.Printf("拼接二级目录后的路径:%v\n", end)
		jackson := strings.Join([]string{end, "entry.json"}, "/")
		title := readEntry(jackson).Title
		i.title = replace(title)
		//log.Printf("视频标题:%v\n", i.title)
		b := getDir(end)
		bottom := strings.Join([]string{end, b[0]}, "/")
		//log.Printf("拼接三级目录后的路径:%v\n", bottom)
		i.audio = strings.Join([]string{bottom, "audio.m4s"}, "/")
		i.video = strings.Join([]string{bottom, "video.m4s"}, "/")
		//log.Printf("视频的路径是%s\n音频的路径是%s\n", i.video, i.audio)
		infos = append(infos, i)
	}
	for _, info := range infos {
		log.Printf("最终获取到的全部结构体%+v\n", info)
	}
	a2v(infos)

}

//适用于一个单集很多P
func ForMulti(dir string) {
	var infos []info
	up := getDir(dir)
	middle := []string{}
	end := []string{}
	for _, v := range up {
		if strings.Contains(v, ".") {
			continue
		}
		if strings.Contains(v, "DS") {
			continue
		}
		m := strings.Join([]string{dir, v}, "/")
		middle = append(middle, m)

	}
	//fmt.Printf("middle is %v\n", middle)
	end = getDir(middle[0])
	//fmt.Printf("end is %v\n", end)
	fullEnd := []string{}
	for _, val := range end {
		f := strings.Join([]string{middle[0], val}, "/")
		fullEnd = append(fullEnd, f)
	}
	//fmt.Printf("full end is %v\n", len(fullEnd))
	for _, value := range fullEnd {
		var i info
		//fmt.Printf("NO.%d `s full is %v\n", index, value)
		j := strings.Join([]string{value, "entry.json"}, "/")
		i.title = readEntry(j).PageData.Part
		//fmt.Printf("json is %v\n", i.title)
		source := getDir(value)
		s := strings.Join([]string{value, source[0]}, "/")
		//fmt.Printf("s dir is %v\n", s)
		i.video = strings.Join([]string{s, "video.m4s"}, "/")
		i.audio = strings.Join([]string{s, "audio.m4s"}, "/")
		//fmt.Printf("strtuc is %v\n", i)
		infos = append(infos, i)
	}
	a2v(infos)
}
func getDir(pwd string) (partname []string) {
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Panicln(err)
	}
	//fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		//fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
		partname = append(partname, fileInfoList[i].Name())
	}
	return partname
}
func readEntry(dir string) (e entry) {
	//var e entry
	bytes, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}

	err = json.Unmarshal(bytes, &e)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return
	}
	log.Printf("获取到的partname:%s\n", e.PageData.Part)
	log.Printf("获取到的title:%s\n", e.Title)

	return e
}
func a2v(infos []info) {
	t := threads.Threads()
	for i, v := range infos {
		fmt.Printf("正在处理第%d个文件\n", i+1)
		v.title = replace(v.title)
		//fmt.Printf("文件名是%s\n", v.title)
		//fmt.Printf("audio路径:%s\n", v.audio)
		//fmt.Printf("video路径:%s\n", v.video)
		log.Printf("最终得到的结构体:%v\n", v)
		fname := strings.Join([]string{v.title, "mp4"}, ".")
		//ffmpeg -i video.m4s -i audio.m4s -codec copy multi.mp4
		cmd := exec.Command("ffmpeg", "-threads", t, "-i", v.video, "-i", v.audio, "-codec", "copy", "-threads", t, fname)
		log.Printf("生成的命令是:%s", cmd)
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
}
func replace(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "，", ",", -1)
	//str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "《", "", -1)
	str = strings.Replace(str, "》", "", -1)
	str = strings.Replace(str, "【", "", -1)
	str = strings.Replace(str, "】", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "_", "", -1)
	str = strings.Replace(str, "·", "", -1)
	str = strings.Replace(str, "\uE000", "", -1)
	str = strings.Replace(str, "、", "", -1)
	str = strings.Replace(str, "\uE000", "", -1)
	//	，
	//	:/usr/local/bin/ffmpeg -threads 3 -i download/207257026/c_386723432/80/video.m4s -i download/207257026/c_386723432/80/audio.m4s -codec copy -thread新三国29曹操真是奸诈无比，连自己的发小许攸，都一骗再骗.mp4

	return str
}
