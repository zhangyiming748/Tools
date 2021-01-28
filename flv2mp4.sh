#!/bin/bash

for file in /home/pi/Share/ffmpeg/flv/*
	do
		echo $file
		ffmpeg -i $file $file".mp4"