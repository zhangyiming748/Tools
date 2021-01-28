# 视频文件转换为音频

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


def extractCommand(Source, Target, file):
    if not os.path.exists(Target):
        os.makedirs(Target)
    absSource = Source + '/' + file
    absTarget = Target + '/' + file

    files = file.split('.')
    file_name = files[0]
    extension_name = files[-1]
    print("abs: ", absSource)
    print("tag: ", absTarget)
    # ffmpeg -i 3.mp4 -vn -y -acodec copy 3.m4a
    prefix = 'ffmpeg -i '
    suffix = ' -vn -y -acodec copy '
    command = prefix + '\"' + absSource + '\"' + suffix + '\"' + Target + '/' + file_name + '.m4a' + '\"'
    print("转换的命令: %s" % command)
    os.system(command)


if __name__ == '__main__':
    Source = '/Users/zen//Downloads/Downie'
    Target = '/Users/zen//Downloads/Downie/Done'
    # /media/zen/Gloway720/process/work
    # Source = '/media/zen/Gloway720/process/work'
    # Target = '/media/zen/Gloway720/process/work/Done'
    relativeFiles = get_dir(Source)
    total = len(relativeFiles)
    print("共有 %d 个文件" % total)
    count = 0
    for file in relativeFiles:
        count+=1
        print("正在处理第 %d 个文件/共 %d 个文件" % (count, total))
        extractCommand(Source=Source, Target=Target, file=file)
        print("处理完成第 %d 个文件/共 %d 个文件" % (count, total))
