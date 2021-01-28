#!/bin/bash
# webp最大宽高不得超过16383像素
for x in ls *.jpg
	do
		cwebp $x -o ${x%.jpg}.webp
	done