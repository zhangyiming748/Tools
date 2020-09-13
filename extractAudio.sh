#!/bin/bash

filename=$(ls *.mp4)
echo $filename

ffmpeg -i $filename -vn -y -acodec copy "media.m4a"
rm $filename
