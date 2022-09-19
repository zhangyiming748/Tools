#!/usr/bin/env bash
echo 删除旧的日志文件
find . -type f -name "*.log" -exec rm {} \;
echo 格式化当前目录下go文件
find . ! -path "./vendor/*" -name "*.go" -exec gofmt -w {} \;
echo 删除多余隐藏文件
find . -name "*DS_Store*" -exec rm {} \;
echo 编译二进制文件forLinux32
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o forLinux32 main.go
echo 编译二进制文件forLinux64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o forLinux64 main.go
echo 编译二进制文件forRaspi
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o forRaspi main.go
echo 编译二进制文件forRaspi64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o forRaspi64 main.go
echo 编译二进制文件forWin32.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o forWin32.exe main.go
echo 编译二进制文件forWin64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o forWin64.exe main.go
echo 编译二进制文件forMac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o forMac main.go
echo 编译二进制文件forM1
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o forM1 main.go
echo 编译二进制文件forAndroid
CGO_ENABLED=0 GOOS=android GOARCH=arm64 go build -o forAndroid main.go
