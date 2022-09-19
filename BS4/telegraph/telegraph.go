package Telegraph

import (
	"Tools/BS4/soup"
	"Tools/util/file"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//func GetOne() {
//resp, err := soup.GetWithProxy("https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-04/1.created.html", "http://127.0.0.1:8889")
//if err != nil {
//	os.Exit(1)
//}
//doc := soup.HTMLParse(resp)
//books := doc.Find("div", "class", "book")
//book := books.Find("div", "class", "book-body")
//inner := book.Find("div", "class", "body-inner")
//wrapper := inner.Find("div", "class", "page-wrapper")
//page := wrapper.Find("div", "class", "page-inner")
//normal := page.Find("section", "class", "normal")
//h3 := normal.Find("h3").Text()
//log.Printf("%+v", h3)
//}

func GetLong(url string) {
	cmd := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	title := doc.Find("header").Find("h1").Text()
	dname := replace(title)

	head := "#!/opt/homebrew/bin/bash"
	cmd = append(cmd, head)
	log.Printf("dir name is %v\n", dname)
	contents := doc.FindAll("figure")
	for i, content := range contents {
		img := content.Find("img").Attrs()["src"]
		src := strings.Join([]string{"https://telegra.ph", img}, "")
		ext := strings.Join([]string{strconv.Itoa(i), "jpg"}, ".")
		fullname := strings.Join([]string{dname, ext}, "/")
		log.Printf("第%d张照片%s\n", i, src)
		line := strings.Join([]string{"wget", "-c", src, "-O", fullname}, " ")
		cmd = append(cmd, line)
	}
	// mkdir
	file.CreateDir(dname)
	shell := strings.Join([]string{dname, "sh"}, ".")
	file.WriteLines(shell, cmd)

	//run
	//var c *exec.Cmd
	//c = exec.Command("bash", "-c", Universal)
	//stdout, err := c.StdoutPipe()
	//c.Stderr = c.Stdout
	//if err != nil {
	//	log.Printf("cmd.StdoutPipe产生的错误:%v", err)
	//}
	//if err = c.Start(); err != nil {
	//	log.Printf("cmd.Run产生的错误:%v", err)
	//}
	//for {
	//	tmp := make([]byte, 1024)
	//	_, err := stdout.Read(tmp)
	//
	//	log.Printf("文件输出:%s", string(tmp))
	//	if err != nil {
	//		break
	//	}
	//}
	//if err = c.Wait(); err != nil {
	//	log.Printf("%v对应文件:%v\n", err, url)
	//}
	//log.Printf("下载文件%v完成\n", url)
	//wait := time.Duration(rand.Intn(3))
	//time.Sleep(wait * time.Second)
}

func replace(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "《", "", -1)
	str = strings.Replace(str, "》", "", -1)
	str = strings.Replace(str, "【", "", -1)
	str = strings.Replace(str, "】", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "/", "", -1)
	str = strings.Replace(str, "&", "", -1)
	return str
}
func GetShort(url string) {
	cmd := make([]string, 0)

	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	log.Printf("doc is %v\n", doc)
	title := doc.Find("header").Find("h1").Text()
	dname := replace(title)
	log.Printf("dir name is %v\n", dname)
	head := "#!/opt/homebrew/bin/bash"
	cmd = append(cmd, head)
	imgs := doc.FindAll("img")
	for i, img := range imgs {
		img := img.Attrs()["src"]
		src := strings.Join([]string{"https://telegra.ph", img}, "")
		ext := strings.Join([]string{strconv.Itoa(i), "jpg"}, ".")
		fullname := strings.Join([]string{dname, ext}, "/")
		log.Printf("第%d张照片%s\n", i, src)
		line := strings.Join([]string{"wget", "--no-check-certificate", "-c", src, "-O", fullname}, " ")
		cmd = append(cmd, line)
	}
	// mkdir
	file.CreateDir(dname)
	shell := strings.Join([]string{dname, "sh"}, ".")
	file.WriteLines(shell, cmd)
	//runShell(Universal)
	//absPath := strings.Join([]string{"/Users/zen/Gitlab/Tools/BS4/telegraph", shell}, "/")
	//Universal.RunShell(absPath)
}
func runShell(sname string) {
	absPath := strings.Join([]string{"/Users/zen/Github/Tools/BS4/telegraph", sname}, "/")
	log.Printf("即将运行的shell:%s\n", absPath)
	sh := exec.Command("bash", "-c", absPath)
	output, err := sh.Output()
	if err != nil {
		log.Panicf("运行shell发生错误:%v\n", err)
	} else {
		log.Printf("%s\n", output)
	}
}
