package spankbang

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFindByChannel(t *testing.T) {
	url := "https://spankbang.com/f4/channel/flexy+teens/"
	FindByChannel(url)
}
func TestFindByChannelAll(t *testing.T) {
	base_url := "https://spankbang.com/f4/channel/flexy+teens/"
	FindByChannel(base_url)
	for t_time := 30; t_time > 0; t_time-- {
		t.Logf("冷却时间,还有%d秒", t_time)
		time.Sleep(time.Second)
	}
	for i := 2; i <= 3; i++ {
		url := strings.Join([]string{base_url, strconv.Itoa(i), "/"}, "")
		FindByChannel(url)
		for t_time := 30 * i; t_time > 0; t_time-- {
			t.Logf("冷却时间,还有%d秒", t_time)
			time.Sleep(time.Second)
		}
	}
}
func TestFindByKeyword(t *testing.T) {
	url := "https://spankbang.com/s/agentredgirl/"
	FindBySearch(url)
}
