package get

//import (
//	"Tools/file/EXIF/decode"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func EXIF2GEO(fp string) string {
//	//defer func() {
//	//	if err := recover(); err != nil {
//	//		Debugln(err)
//	//	}
//	//}()
//	f, _ := os.Open(fp)
//	x, _ := decode.Decode(f)
//	x, _ := decode.Decode(f)
//	//fmt.Println(x)
//	Longitude, _ := x.Get(exif.GPSLongitude)
//	Latitude, _ := x.Get(exif.GPSLatitude)
//	//Debugln(Longitude)
//	//Debugln(Latitude)
//	Long, err1 := strconv.ParseFloat(fmt.Sprintf("%v", Longitude), 64)
//	Lat, err2 := strconv.ParseFloat(fmt.Sprintf("%v", Latitude), 64)
//	if err1 == nil || err2 == nil {
//		if Long < 180 || Lat < 90 {
//			ret := fmt.Sprintf("%f,%v", Long, Lat)
//			return ret
//		}
//	}
//	long := convert(Longitude)
//	lat := convert(Latitude)
//	ret := fmt.Sprintf("%f,%v", long, lat)
//	return ret
//}
//func convert(tag *tiff.Tag) float64 {
//	step1 := fmt.Sprintf("%v", tag)
//	step2 := strings.Replace(step1, "\"", "", -1)
//	step3 := strings.Replace(step2, "[", "", -1)
//	step4 := strings.Replace(step3, "]", "", -1)
//	step5 := strings.Split(step4, ",")
//	hs := strings.Split(step5[0], "/")[0]
//	ms := strings.Split(step5[1], "/")[0]
//	ss := strings.Split(step5[2], "/")[0]
//	h, _ := strconv.ParseFloat(hs, 64)
//	m, _ := strconv.ParseFloat(ms, 64)
//	s, _ := strconv.ParseFloat(ss, 64)
//	if strings.Split(step5[0], "/")[1] == "1000" || strings.Split(step5[1], "/")[1] == "1000" || strings.Split(step5[2], "/")[1] == "1000" {
//		ret := h/1000 + m/1000/60 + s/1000/3600
//		return ret
//	}
//	ret := h + m/60 + s/1000000/3600
//	return ret
//}
