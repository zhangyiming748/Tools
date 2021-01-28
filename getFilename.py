import os


def get_dir(dictionary):
    multi_files = os.listdir(dictionary)
    files = []
    for multi in multi_files:
        names = multi.split('.')
        # print(names)
        prefix = names[0]
        suffix = names[-1]
        if suffix != 'a':

            files.append(multi)
            # renameCommand(dictionary+'/'+multi,dictionary+'/'+onlyNum(multi))
        else:
            print('跳过 %s 文件' % multi)
    return files


def justGetlist(dictionary):
    files = os.listdir(dictionary)
    for file in files:
        if file != '.DS_Store':
            full = dictionary + '/' + file
            print(full)
            write2file(dictionary, full)


def onlyNum(file):
    names = file.split('-')
    newName = names[0]
    # print("切割的文件名是: %s"%newName)
    return newName


def renameCommand(old, new):
    cmd = 'mv ' + '\"' + old + '\"' + ' ' + '\"' + new + '.avi\"'
    print("运行前的命令: %s" % cmd)
    os.system(cmd)


def write2file(dictionary, line):
    target = dictionary + '/file.txt'
    with open(target, "a+", encoding="utf-8") as f:
        f.write('file ')
        # f.write('\"')
        f.write(line)
        # f.write('\"')
        f.write('\n')


if __name__ == '__main__':
    dir = '/Users/zen/Downloads/dirc/1'
    nums = justGetlist(dir)
    # for num in nums:
    #     onlyNum(num)
    # print(nums)
