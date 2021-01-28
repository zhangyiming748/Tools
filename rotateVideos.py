import os


# 获取视频文件列表并返回
def get_dir(dictionary):
    multi_files = os.listdir(dictionary)
    files = []
    for multi in multi_files:
        names = multi.split('.')
        # print(names)
        prefix = names[0]
        suffix = names[-1]
        if suffix == 'mp4':
            files.append(multi)
        else:
            print('跳过 %s 文件' % multi)
    return files


def rotateCommand(Source, Target, file):
    if not os.path.exists(Target):
        os.makedirs(Target)
    absSource = Source + '/' + file
    absTarget = Target + '/' + file

    print("abs: ", absSource)
    print("tag: ", absTarget)
    # ffmpeg -i $file -vf "transpose=1" $file".mp4"
    prefix = 'ffmpeg -i '
    suffix = ' -vf "transpose=2" '
    command = prefix + '\"' + absSource + '\"' + suffix + '\"' + absTarget + '\"'
    print("转换的命令: %s" % command)
    os.system(command)


if __name__ == '__main__':
    Source = '/Users/zen/Downloads/Downie'
    Target = '/Users/zen/Downloads/Downie/Done'
    relativeFiles = get_dir(Source)
    total = len(relativeFiles)
    print("共有 %d 个文件" % total)
    for file in relativeFiles:
        rotateCommand(Source=Source, Target=Target, file=file)
