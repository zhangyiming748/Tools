package BS4

import (
	"log"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFindByActress(t *testing.T) {
	url := "https://javdb40.com/actors/ky60"
	FindByActress(url)
}
func TestFindByActress2(t *testing.T) {
	url := "https://javdb40.com/actors/ky60"
	FindByActress2(url)
}
func TestFindByActress2All(t *testing.T) {
	for i := 2; i < 8; i++ {
		url := strings.Join([]string{"https://javdb.com/actors/ky60?page", strconv.Itoa(i)}, "=")
		t.Logf("url = %s\n", url)
		FindByActress2(url)
		for j := i * 5; j > 0; j-- {
			t.Logf("查询完当前页面,冷却%d秒后进入下一页\n", j)
			time.Sleep(time.Second)
		}
	}
}
func TestFindDetialMainPics(t *testing.T) {
	url := "https://javdb.com/v/rngN"
	pics := FindDetialMainPics(url)
	for i, pic := range pics {
		t.Logf("%d.%v\n", i, pic)
	}
}
func TestFindByActressAll(t *testing.T) {
	url := "https://javdb.com/actors/ky60"
	FindByActress(url)
	for j := 30; j > 0; j-- {
		t.Logf("查询完当前页面,冷却%d秒后进入下一页\n", j)
		time.Sleep(time.Second)
	}
	for i := 9; i < 23; i++ {
		url := strings.Join([]string{"https://javdb.com/actors/ky60?page", strconv.Itoa(i)}, "=")
		t.Logf("url = %s\n", url)
		FindByActress(url)
		for j := i * 5; j > 0; j-- {
			t.Logf("查询完当前页面,冷却%d秒后进入下一页\n", j)
			time.Sleep(time.Second)
		}

	}
}

func TestFindActresses(t *testing.T) {
	Furl := "https://javdb.com/actors/western"
	FindActresses(Furl)
	for i := 30; i > 0; i-- {
		log.Printf("查询完首页,冷却%d秒后进入下一页\n", i)
		time.Sleep(time.Second)
	}
	for i := 2; i <= 30; i++ {
		url := strings.Join([]string{"https://javdb.com/actors/western?page", strconv.Itoa(i)}, "=")
		log.Printf("正在处理的网页是%s\n", url)
		FindActresses(url)
		for j := i * 5; j > 0; j-- {
			t.Logf("查询完当前页面,冷却%d秒后进入下一页\n", j)
			time.Sleep(time.Second)
		}
		log.Printf("处理完成的网页是%s\n", url)
	}
}
