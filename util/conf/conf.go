package conf

import (
	"github.com/zhangyiming748/goini"
	"log"
)

const confPath = "./conf.ini"

var (
	RunMode string
	conf    *goini.Config
)

/*
*
  - 初始化
    init函数的主要作用：
    初始化不能采用初始化表达式初始化的变量。
    程序运行前的注册。
    实现sync.Once功能。
    其他
*/
func init() {
	initConfig()
}

func initConfig() {
	conf = goini.SetConfig(confPath)
	log.Println(confPath)
	RunMode, err := conf.GetValue("runmode", "mode")
	if err != nil {
		log.Println("not found key: runMode")
	} else {
		log.Printf("initialization get runMode:%s\n", RunMode)
	}
}

/**
 * 获取环境变量
 */
func GetEnv() string {
	if RunMode == "" {
		initConfig()
	}
	return RunMode
}

/**
 * 根据键获取值
 */
func GetVal(section, name string) string {
	if section == "" {
		section = GetEnv()
	}
	val, err := conf.GetValue(section, name)
	if err != nil {
		return "not found"
	}
	return val
}
