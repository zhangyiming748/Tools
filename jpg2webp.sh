#!/bin/bash
# webp最大宽高不得超过16383像素
#!/bin/bash
# get all filename in specified path

path=$1
files=$(ls *.png $path)
for filename in $files
	do
		cwebp $filename -o ${filename%.png}.webp
	done