#!/usr/bin/python
# -*- coding: UTF-8 -*-
from glob import glob


def zh_ch2en_us(files):
    for file in files:
        print(file)
        f = open(file, 'r+')
        all_the_lines = f.readlines()
        f.seek(0)
        f.truncate()
        for line in all_the_lines:
            line = line.replace('，', ',')
            line = line.replace('。', '.')
            line = line.replace('（', '(')
            line = line.replace('）', ')')
            line = line.replace('“', '\"')
            line = line.replace('”', '\"')
            line = line.replace('：', ':')
            line = line.replace('；', ';')
            line = line.replace('？', '?')
            line = line.replace('！', '!')
            line = line.replace('《', '<')
            line = line.replace('》', '>')
            line = line.replace('【', '[')
            line = line.replace('】', ']')
            line = line.replace('、', '\\')
            line = line.replace('～', '~')
            f.write(line)
        f.close()


if __name__ == '__main__':
    path = glob('*.md')
    print(path)
    zh_ch2en_us(path)
