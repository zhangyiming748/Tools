package q10drfc

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFindByKeyword(t *testing.T) {
	rand.Seed(time.Now().Unix())
	seed := rand.Intn(10)
	original_url := "https://www.q10drfc.com/89152.html"
	for i := 1; i <= 11; i++ {
		if i == 1 {
			log.Printf("正在处理第%d页\n", i)
			url := original_url
			FindByKeyword(url)
			log.Printf("第%d页处理完成\n", i)
			for t_time := seed; t_time > 0; t_time-- {
				t.Logf("冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		} else {
			log.Printf("正在处理第%d页\n", i)
			url := strings.Join([]string{original_url, strconv.Itoa(i)}, "/")
			log.Printf("新拼接的url是%s\n", url)
			FindByKeyword(url)
			log.Printf("第%d页处理完成\n", i)
			for t_time := seed; t_time > 0; t_time-- {
				t.Logf("冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		}
	}
}

func TestMulti(t *testing.T) {
	start := 1
	end := 19
	code := 109214
	do(start, end, code)
}
func do(start, end, code int) {
	rand.Seed(time.Now().Unix())
	seed := rand.Intn(10)
	original_url := "https://www.q10drfc.com/89152.html"
	original_url = strings.Join([]string{"https://www.q10drfc.com/", strconv.Itoa(code), ".html"}, "")
	for i := start; i <= end; i++ {
		if i == 1 {
			log.Printf("正在处理第%d页\n", i)
			url := original_url
			FindByKeyword(url)
			log.Printf("第%d页处理完成\n", i)
			for t_time := seed; t_time > 0; t_time-- {
				log.Printf("冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		} else {
			log.Printf("正在处理第%d页\n", i)
			url := strings.Join([]string{original_url, strconv.Itoa(i)}, "/")
			log.Printf("新拼接的url是%s\n", url)
			FindByKeyword(url)
			log.Printf("第%d页处理完成\n", i)
			for t_time := seed; t_time > 0; t_time-- {
				log.Printf("冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		}
	}
}
func TestFindBySearch(t *testing.T) {
	rand.Seed(time.Now().Unix())
	seed := rand.Intn(10)
	original_url := "https://www.q10drfc.com/107550.html"
	for i := 1; i <= 6; i++ {
		if i == 1 {
			log.Printf("正在处理第%d页\n", i)
			url := original_url
			FindBySearch(url)
			log.Printf("第%d页处理完成\n", i)
			for t_time := seed; t_time > 0; t_time-- {
				t.Logf("冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		} else {
			log.Printf("正在处理第%d页\n", i)
			url := strings.Join([]string{original_url, strconv.Itoa(i)}, "/")
			log.Printf("新拼接的url是%s\n", url)
			FindBySearch(url)
			log.Printf("第%d页处理完成\n", i)
			for t_time := seed; t_time > 0; t_time-- {
				t.Logf("冷却时间,还有%d秒", t_time)
				time.Sleep(time.Second)
			}
		}
	}
}
