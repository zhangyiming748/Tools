package Telegraph

import (
	"Tools/util/file"
	"strings"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	urls := []string{
		"https://telegra.ph/%E7%A7%80%E4%BA%BA%E7%BD%91-NO3128-%E6%9D%A8%E6%99%A8%E6%99%A8sugar54P-04-14",
		"https://telegra.ph/%E7%A7%80%E4%BA%BA%E7%BD%91-NO3129-%E9%99%88%E5%B0%8F%E5%96%B583P-04-15",
	}
	for i, url := range urls {
		t.Logf("正在下载第%d个网页\n", i+1)
		GetLong(url)
		for t_time := 30; t_time > 0; t_time-- {
			t.Logf("下载网页完毕,冷却时间,还有%d秒", t_time)
			time.Sleep(time.Second)
		}
	}

	// GetLong("https://telegra.ph/%E6%9F%9A%E6%9C%A856-58-03-15")
}
func TestGetShort(t *testing.T) {
	url := "https://telegra.ph/Qfq2Ar-03-13"
	GetShort(url)
}
func TestGetShorts(t *testing.T) {
	urls := file.ReadLink("todo.list")
	for i, url := range urls {
		GetShort(url)
		for t_time := 3 * i; t_time > 0; t_time-- {
			t.Logf("冷却时间,还有%d秒", t_time)
			time.Sleep(time.Second)
		}
	}
}

func TestMultiUrl(t *testing.T) {
	urls := []string{
		"https://telegra.ph/qMNzy2-04-01",
		"",
		"",
		"",
		"",
		"",
	}
	for i, url := range urls {
		if strings.Contains(url, "%") {
			t.Logf("正在下载第%d个网页\n", i+1)
			GetLong(url)
			for t_time := 30; t_time > 0; t_time-- {
				t.Logf("下载网页完毕,冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		} else if url != "" {
			t.Logf("正在下载第%d个网页\n", i+1)
			GetShort(url)
			for t_time := 30; t_time > 0; t_time-- {
				t.Logf("下载网页完毕,冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		} else {
			continue
		}
	}
}
func TestRunshell(t *testing.T) {
	runShell("test.sh")
}
