package convert

import (
	"time"
)

/*
Sep 23, 2021 06:42:47.778473844 UTC
Sep 18, 2021 17:45:16.460053843 CST
*/
var (
	moon = map[string]string{
		"Jan":  "01",
		"Feb":  "01",
		"Mar":  "01",
		"Apr":  "01",
		"May":  "01",
		"Jun":  "01",
		"Jul":  "01",
		"Aug":  "01",
		"Sept": "01",
		"Oct":  "10",
		"Nov":  "01",
		"Dec":  "01",
	}
)

func readLocalTime(s string) string {
	parse, _ := time.Parse("Jan 02, 2006 15:04:05.000000000 MST", s)
	parse.Location()
	return parse.Local().String()
}
