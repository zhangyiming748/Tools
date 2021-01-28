import os
import time
import datetime


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
        # 如果后缀为webm 加入选定列表
        #################如果是隐藏文件的点
        if suffix == 'webm' and len(prefix) > 0:
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
    absTarget = Target + '/' + file + '.mkv'
    print(absTarget)
    # print("abs: ", absSource)
    # print("tag: ", absTarget)
    # ffmpeg -i input.webm -c copy output.mp4
    prefix = 'ffmpeg -i '
    # suffix = ' -c copy '
    suffix = ' '
    command = prefix + '\"' + absSource + '\"' + suffix + '\"' + absTarget + '\"'
    print("转换的命令: %s" % command)
    s = str()

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


def toLog(total, current, name, status):
    thisTime = str(time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time())))
    with open('/Users/zen/Desktop/log.txt', 'a+', encoding='utf-8')as f:
        myLog = str('当前时间%s\t共有%s个文件\t正在转换第%s个\t文件名:%s\n' % (thisTime, total, current, name))
        myLogDone = str('当前时间%s\t共有%s个文件\t转换完成第%s个\t转换用时:%s\n' % (thisTime, total, current, name))
        if status == 'start':
            f.write(myLog)
        if status == 'end':
            f.write(myLogDone)


def Duration(oldtime, newtime):
    min = (newtime - oldtime).seconds // 60
    sec = (newtime - oldtime).seconds % 60
    result = str('%s分钟零%s秒' % (min, sec))
    return result


if __name__ == '__main__':
    Source = '/Users/zen/Videos'
    Target = '/Users/zen/Videos'
    relativeFiles = get_dir(Source)
    total = len(relativeFiles)
    print("共有 %d 个文件" % total)
    count = 1
    for file in relativeFiles:
        notification(total, count)
        toLog(total, count, file, 'start')
        oldTime = datetime.datetime.now()
        convertCommand(Source=Source, Target=Target, file=file)
        newTime = datetime.datetime.now()
        duration = Duration(oldTime, newTime)
        toLog(total, count, duration, 'end')
        count = count + 1
