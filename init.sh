#!/usr/bin/env bash
echo 删除旧的日志文件
find . -type f -name "*.log" -exec rm {} \;
echo 格式化当前目录下go文件
find . ! -path "./vendor/*" -name "*.go" -exec gofmt -w {} \;
echo 删除多余隐藏文件
find . -name "*DS_Store*" -exec rm {} \;