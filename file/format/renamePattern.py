import os


# 获取文件名(列表)

def get_dir(dictionary):
    multi_files = os.listdir(dictionary)
    files = []
    for multi in multi_files:
        names = multi.split('.')
        # print(names)
        prefix = names[0]
        suffix = names[-1]
        if suffix == 'mp3':
            files.append(multi)
        else:
            print('跳过 %s 文件' % multi)
    return files


# 按照分隔符分割
def asDelimiter(file):
    chars = file.split(']')
    print(chars)
    newName = str(chars[-1])
    # for i in chars[6:]:
    #     newName+=i
    #
    print("-------------",chars[-1])
    print("newname = %s" %newName)
    return newName


def renameCommand(old, new):
    cmd = 'mv ' + old + ' ' + new
    print("运行前的命令: %s" % cmd)
    os.system(cmd)


if __name__ == '__main__':
    path = '/Users/zen/Documents/rename/pattern'
    relativeFiles = get_dir(path)
    total = len(relativeFiles)
    print("共有 %d 个文件" % total)
    for file in relativeFiles:
        perfix = asDelimiter(file)
        suffix = '.mp3'
        oldABS = '\"' + path + '/' + file + '\"'
        newABS = '\"' + path + '/' + perfix  + '\"'
        print("旧文件名是: %s,新文件名是%s" % (oldABS, newABS))
        renameCommand(oldABS, newABS)
