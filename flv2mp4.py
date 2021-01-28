import os

import datetime
import time


# webm批量转换mp4
# 获取视频文件列表并返回
def get_dir(dictionary):
    # 遍历目录下所有文件
    multi_files = os.listdir(dictionary)
    files = []
    for multi in multi_files:
        names = multi.split('.')
        # print(names)
        prefix = names[0]
        suffix = names[-1]
        # 如果后缀为flv 且不是文件挂载点则加入选定列表
        if suffix == 'flv' and len(prefix) > 0:
            files.append(multi)
        else:
            print('跳过 %s 文件' % multi)
    return files


def convertCommand(Source, Target, file):
    # 如果文件夹不存在则新建
    if not os.path.exists(Target):
        os.makedirs(Target)
    absSource = Source + '/' + file
    files = file.rsplit('.', 1)
    print(files)
    file = files[0]
    absTarget = Target + '/' + file + '.mp4'
    print(absTarget)
    print("abs: ", absSource)
    print("tag: ", absTarget)
    # ffmpeg -i input.flv  output.mp4
    prefix = 'ffmpeg -i '
    suffix = ' '
    command = prefix + '\"' + absSource + '\"' + suffix + '\"' + absTarget + '\"'
    print("转换的命令: %s" % command)
    os.system(command)


def notification(total, current):
    msg = str('共有%s个文件,正在转换第%s个' % (total, current))
    print('\033[1;47m%s\033[0m' % msg)
    print('\033[5;41m%s\033[0m' % msg)
    print('\033[7;40m%s\033[0m' % msg)
    print('\033[1;47m%s\033[0m' % msg)
    print('\033[5;41m%s\033[0m' % msg)
    print('\033[7;40m%s\033[0m' % msg)
    print('\033[1;47m%s\033[0m' % msg)
    print('\033[5;41m%s\033[0m' % msg)
    print('\033[7;40m%s\033[0m' % msg)
def timeReport(msg):
    print('\033[1;47m%s\033[0m' % msg)
    print('\033[5;41m%s\033[0m' % msg)
    print('\033[7;40m%s\033[0m' % msg)
    print('\033[1;47m%s\033[0m' % msg)
    print('\033[5;41m%s\033[0m' % msg)
    print('\033[7;40m%s\033[0m' % msg)
    print('\033[1;47m%s\033[0m' % msg)
    print('\033[5;41m%s\033[0m' % msg)
    print('\033[7;40m%s\033[0m' % msg)

def timer(ot, nt):
    ot = datetime.datetime.now()
    print(ot)
    time.sleep(1)
    nt = datetime.datetime.now()
    print(nt)
    print('：%s' % (nt - ot))
    print('用时：%s微秒' % (nt - ot).microseconds)
    print('用时：%s秒' % (nt - ot).seconds)
    print('用时：%s分' % (nt - ot).seconds*60)


if __name__ == '__main__':
    Source = '/Users/zen/Downloads/Test'
    Target = '/Users/zen/Downloads/Test/Done'
    relativeFiles = get_dir(Source)
    total = len(relativeFiles)
    print("共有 %d 个文件" % total)
    count = 1
    start_time=datetime.datetime.now()
    for file in relativeFiles:
        notification(total, count)
        part_start=datetime.datetime.now()
        convertCommand(Source=Source, Target=Target, file=file)
        print("当前部分用时")
        part_end=datetime.datetime.now()
        timer(part_start, part_end)
        count += 1
    print("整个过程用时")
    end_time=datetime.datetime.now()
