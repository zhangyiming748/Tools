#!/bin/bash
# 从视频中提取音频

filename=$(ls *.mp4)
echo $filename

ffmpeg -i $filename -vn -y -acodec copy "name.m4a"
rm $filename