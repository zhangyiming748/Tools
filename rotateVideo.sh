#!/bin/bash
# 顺时针旋转当前目录下的视频
# 其中transpose取值:
# 0 = 90CounterCLockwise and Vertical Flip (default)
# 1 = 90Clockwise
# 2 = 90CounterClockwise
# 3 = 90Clockwise and Vertical Flip
for file in /Users/zen/Downloads/Downie/*
	do

	echo $file
	ffmpeg -i $file -vf "transpose=1" $file".mp4"
	done
