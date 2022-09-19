#!/bin/bash
echo 编译Linux32
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o forLinux32 beautifulSoup.go
echo 编译Linux32结束
echo 编译Linux64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o forLinux64 beautifulSoup.go
echo 编译Linux64结束
echo 编译LinuxARM
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o forRaspi beautifulSoup.go
echo 编译LinuxARM结束
echo 编译Windows32
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o forWin32.exe beautifulSoup.go
echo 编译Windows32结束
echo 编译Windows64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o forWin64.exe beautifulSoup.go
echo 编译Windows64结束
echo 编译MacARM
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o forMac beautifulSoup.go
echo 编译MacARM结束
echo 编译Mac
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o forM1 beautifulSoup.go
echo 编译Mac结束