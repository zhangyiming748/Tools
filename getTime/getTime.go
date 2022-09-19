package getTime

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func init() {
	log.SetFlags(2 | 16)
}

func GetTime() {
	var (
		start string
		end   string
	)
	log.Println("input start time")
	if _, err := fmt.Scanln(&start); err != nil {
		panic("输入格式有误(HH:MM:SS)")
	}
	log.Println("input end time")
	if _, err := fmt.Scanln(&end); err != nil {
		panic("输入格式有误(HH:MM:SS)")
	}
	start = addColon(start)
	end = addColon(end)
	//time1 := "2015-03-20 08:50:29"
	//time2 := "2015-03-20 09:04:25"
	log.Printf("添加冒号之后的开始时间 = %s\n添加冒号之后的结束时间 = %s\n", start, end)
	tstart := strings.Join([]string{"2015-03-20", start}, " ") //01:03:44
	tend := strings.Join([]string{"2015-03-20", end}, " ")     //01:05:43
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", tstart)
	t2, err := time.Parse("2006-01-02 15:04:05", tend)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		log.Println("true")
	} else {
		log.Println("开始时间晚于结束时间")
	}
	t3 := t2.Sub(t1)
	t := t3.String()
	log.Printf("t=%v", t)
	log.Printf("%T", t)
	res := formatTime(t)
	log.Printf("after process t = %v", res)

}
func formatTime(s string) string {
	sb := []byte(s)
	resb := []byte{}

	for _, v := range sb {
		if strings.Contains(s, "h") { //带小时
			if v == 'h' || v == 'm' {
				resb = append(resb, ':')
			}
			if v != 'h' && v != 'm' {
				resb = append(resb, v)
			}
			if v == 's' {
				continue
			}
		} else { //不带小时
			if strings.Contains(s, "m") { //带分钟
				if v == 'm' {
					resb = append(resb, ':')
				}
				if v != 'm' {
					resb = append(resb, v)
				}
				if v == 's' {
					continue
				}
			} else { //不带分钟
				if v != 's' {
					resb = append(resb, v)
				}
			}
		}
	}
	return string(resb)
}
func addColon(s string) string {
	sb := []byte(s)
	resb := []byte{}
	for i, v := range sb {
		resb = append(resb, v)
		if i%2 != 0 && i != len(sb) {
			resb = append(resb, ':')
		}
	}
	log.Println(resb)
	resb = resb[:len(resb)-1]
	return string(resb)
}
