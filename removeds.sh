#!/bin/bash
echo 删除多余隐藏文件
find . -name "*DS_Store*" -exec rm {} \;