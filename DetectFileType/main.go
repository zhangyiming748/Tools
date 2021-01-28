package main

import (
	"DetectFileType/getFileType"
	"fmt"
	"os"
)
//探测文件类型,需要放到被探测文件同级目录
func main() {
	 //for i,v:=range os.Args{
	 //	fmt.Printf("第%d个参数是%s\n",i+1,v)
	 //}
	 if len(os.Args)<=1{
	 	fmt.Println("需要输入文件名作为参数")
	 	os.Exit(1)
	 }
	 for i:=1;i<len(os.Args);i++{
	 	t:=getFileType.Master(os.Args[i])
	 	fmt.Printf("第 %v 个文件 %v 的文件类型是 %s\n",i,os.Args[i],t)
	 }
	 //getFileType.Master(os.Args[1])
}
