package PHUB

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

//https://www.pornhub.com/channels/flexy-teens/videos?page=8
func TestFindByChannel(t *testing.T) {
	channel_url_name := "dollhouse-films"
	index := strings.Join([]string{"https://www.pornhub.com/channels", channel_url_name, "videos"}, "/")
	//index := "https://www.pornhub.com/channels/nylon-up/videos"
	FindByChannel(index)
	for t_time := 30; t_time > 0; t_time-- {
		t.Logf("冷却时间,还有%d秒", t_time)
		time.Sleep(time.Second)
	}
	prefix := strings.Join([]string{index, "page"}, "?")
	for i := 2; i <= 3; i++ {
		url := strings.Join([]string{prefix, strconv.Itoa(i)}, "=")
		t.Logf("即将获取的地址是:%s\n", url)
		FindByChannelAfter2(url)
		t.Logf("第%d页获取完毕\n", i)
		for t_time := 30; t_time > 0; t_time-- {
			t.Logf("冷却时间,还有%d秒", t_time)
			time.Sleep(time.Second)
		}
	}
}
func TestFindBySingleChannel(t *testing.T) {
	//https://www.pornhub.com/channels/sarina-valentina/videos
	FindByChannel("https://www.pornhub.com/channels/sarina-valentina/videos")
}
func TestFindByStarUpload(t *testing.T) {
	index := "https://www.pornhub.com/pornstar/miss-ellie-mouse/videos/upload"
	FindByStarUpload(index)
	time.Sleep(30 * time.Second)
	prefix := strings.Join([]string{index, "page"}, "?")
	for i := 2; i <= 5; i++ {
		url := strings.Join([]string{prefix, strconv.Itoa(i)}, "=")
		t.Logf("即将获取的地址是:%s\n", url)
		FindByStarUpload(url)
		t.Logf("第%d页获取完毕\n", i)
		for t_time := 30; t_time > 0; t_time-- {
			t.Logf("冷却时间,还有%d秒", t_time)
			time.Sleep(time.Second)
		}
	}

}
func TestFindByModel2(t *testing.T) {
	prefix := "https://www.pornhub.com/model/theoryofsex/videos?page"
	for i := 2; i <= 4; i++ {
		url := strings.Join([]string{prefix, strconv.Itoa(i)}, "=")
		t.Logf("即将获取的地址是:%s\n", url)
		FindByModel(url)
		for t_time := 30; t_time > 0; t_time-- {
			t.Logf("冷却时间,还有%d秒", t_time)
			time.Sleep(time.Second)
		}
	}
	url := "https://www.pornhub.com/model/theoryofsex/videos"
	FindByModel(url)
}

func TestFindByChannelAfter2(t *testing.T) {
	url := "https://www.pornhub.com/channels/straplezz/videos?page=4"
	FindByChannelAfter2(url)
}
