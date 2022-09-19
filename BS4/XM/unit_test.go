package XM

import (
	"testing"
	"time"
)

func TestFindByChannel(t *testing.T) {
	url := "https://xhamster.com/channels/straplezz"
	FindByChannel(url)
	for t_time := 30; t_time > 0; t_time-- {
		t.Logf("冷却时间,还有%d秒", t_time)
		time.Sleep(time.Second)
	}
	url1 := "https://xhamster.com/channels/straplezz/2"
	FindByChannel(url1)
	for t_time := 30; t_time > 0; t_time-- {
		t.Logf("冷却时间,还有%d秒", t_time)
		time.Sleep(time.Second)
	}
	url2 := "https://xhamster.com/channels/straplezz/3"
	FindByChannel(url2)
}
