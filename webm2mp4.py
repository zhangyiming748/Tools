import os


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
        # 如果后缀为webm 且不是文件挂载点则加入选定列表
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
    absTarget = Target + '/' + file + '.mp4'
    print(absTarget)
    print("abs: ", absSource)
    print("tag: ", absTarget)
    # ffmpeg -i input.webm -c copy output.mp4
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


if __name__ == '__main__':
    Source = '/Users/zen/Downloads/Test'
    Target = '/Users/zen/Downloads/Test/Done'
    relativeFiles = get_dir(Source)
    total = len(relativeFiles)
    print("共有 %d 个文件" % total)
    count = 1
    for file in relativeFiles:
        notification(total, count)
        convertCommand(Source=Source, Target=Target, file=file)
        count += 1
