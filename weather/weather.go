package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	key = "3e469a45a72b31c120b0baf6609ae877"
)

const (
	北京市石景山区     = "110107"
	黑龙江省大庆市让胡路区 = "230604"
	海南省琼海市      = "469002"
)

func Weather() {
	m := make(map[string]string)
	//m["全国"] = "100000"
	m["北京市石景山区"] = "110107"
	m["黑龙江省大庆市让胡路区"] = "230604"
	m["海南省琼海市"] = "469002"
	for k, v := range m {
		log.Printf("%s的天气信息如下\n", k)
		format(v)
	}
}
func format(adcode string) {
	var m map[string]interface{}
	url := "https://restapi.amap.com/v3/weather/weatherInfo?output=json&extensions=all&city=" + adcode + "&key=" + key
	response, _ := http.Get(url)
	defer response.Body.Close() //在回复后必须关闭回复的主体
	body, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(body, &m)
	reporttime := m["forecasts"].([]interface{})[0].(map[string]interface{})["reporttime"]
	log.Printf("%v发布\n", reporttime)
	today_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["date"]
	today_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["week"]
	switch today_week {
	case "1":
		today_week = "一"
	case "2":
		today_week = "二"
	case "3":
		today_week = "三"
	case "4":
		today_week = "四"
	case "5":
		today_week = "五"
	case "6":
		today_week = "六"
	case "7":
		today_week = "日"
	}
	today_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["dayweather"]
	today_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nightweather"]
	today_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["daytemp"]
	today_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nighttemp"]
	today_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["daywind"]
	today_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nightwind"]
	today_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["daypower"]
	today_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nightpower"]
	log.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C %v风%v级\n", today_date, today_week, today_dayweather, today_daytemp, today_daywind, today_daypower, today_nightweather, today_nighttemp, today_nightwind, today_nightpower)
	after_24_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["date"]
	after_24_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["week"]
	switch after_24_week {
	case "1":
		after_24_week = "一"
	case "2":
		after_24_week = "二"
	case "3":
		after_24_week = "三"
	case "4":
		after_24_week = "四"
	case "5":
		after_24_week = "五"
	case "6":
		after_24_week = "六"
	case "7":
		after_24_week = "日"
	}
	after_24_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["dayweather"]
	after_24_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nightweather"]
	after_24_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["daytemp"]
	after_24_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nighttemp"]
	after_24_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["daywind"]
	after_24_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nightwind"]
	after_24_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["daypower"]
	after_24_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nightpower"]
	log.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C %v风%v级\n", after_24_date, after_24_week, after_24_dayweather, after_24_daytemp, after_24_daywind, after_24_daypower, after_24_nightweather, after_24_nighttemp, after_24_nightwind, after_24_nightpower)
	after_48_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["date"]
	after_48_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["week"]
	switch after_48_week {
	case "1":
		after_48_week = "一"
	case "2":
		after_48_week = "二"
	case "3":
		after_48_week = "三"
	case "4":
		after_48_week = "四"
	case "5":
		after_48_week = "五"
	case "6":
		after_48_week = "六"
	case "7":
		after_48_week = "日"
	}
	after_48_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["dayweather"]
	after_48_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nightweather"]
	after_48_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["daytemp"]
	after_48_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nighttemp"]
	after_48_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["daywind"]
	after_48_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nightwind"]
	after_48_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["daypower"]
	after_48_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nightpower"]
	log.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C %v风%v级\n", after_48_date, after_48_week, after_48_dayweather, after_48_daytemp, after_48_daywind, after_48_daypower, after_48_nightweather, after_48_nighttemp, after_48_nightwind, after_48_nightpower)

	after_72_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["date"]
	after_72_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["week"]
	switch after_72_week {
	case "1":
		after_72_week = "一"
	case "2":
		after_72_week = "二"
	case "3":
		after_72_week = "三"
	case "4":
		after_72_week = "四"
	case "5":
		after_72_week = "五"
	case "6":
		after_72_week = "六"
	case "7":
		after_72_week = "日"
	}
	after_72_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["dayweather"]
	after_72_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nightweather"]
	after_72_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["daytemp"]
	after_72_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nighttemp"]
	after_72_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["daywind"]
	after_72_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nightwind"]
	after_72_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["daypower"]
	after_72_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nightpower"]
	log.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C %v风%v级\n\n", after_72_date, after_72_week, after_72_dayweather, after_72_daytemp, after_72_daywind, after_72_daypower, after_72_nightweather, after_72_nighttemp, after_72_nightwind, after_72_nightpower)
}

const (
	bajiaozhongli = "Beijing+Shijingshan+bajiao"
	qihoo         = "Beijing+Chaoyang+Wangjing"
)

// wget wttr.in/Beijing+Chaoyang+Jiuxianqiao.png
func WeatherPNG() {
	locations := []string{bajiaozhongli, qihoo}
	for _, location := range locations {
		cmd(location)
	}

}
func cmd(s string) {
	defer log.Println("生成图片进程结束")
	perfix := strings.Join([]string{"wttr.in", s}, "/")
	suffix := strings.Join([]string{perfix, "png"}, ".")
	cmd := exec.Command("wget", suffix)
	log.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
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
