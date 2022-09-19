# 重命名文件替换中文符号
# 在macOS/Linux下正常运行
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
        if suffix != 'py':
            files.append(multi)
        else:
            print('跳过 %s 文件' % multi)
    return files


# 替换中文符号生成新文件名
def replace(oldName):  # 【上午】
    oldName = oldName.replace('，', ',')
    oldName = oldName.replace('。', '.')
    oldName = oldName.replace('（', '(')
    oldName = oldName.replace('）', ')')
    oldName = oldName.replace('“', '\"')
    oldName = oldName.replace('”', '\"')
    oldName = oldName.replace('：', ':')
    oldName = oldName.replace('；', ';')
    oldName = oldName.replace('？', '?')
    oldName = oldName.replace('！', '!')
    oldName = oldName.replace('《', '<')
    oldName = oldName.replace('》', '>')
    oldName = oldName.replace('【', '[')
    oldName = oldName.replace('】', ']')
    oldName = oldName.replace('、', '\\')
    oldName = oldName.replace('～', '~')
    oldName = oldName.replace(' ','')
    # print("生成的文件名是 %s"%oldName)
    newName = oldName
    return newName


def renameCommand(old, new):

    cmd = 'mv ' + old + ' ' + new
    print("运行前的命令: %s"%cmd)
    os.system(cmd)


if __name__ == '__main__':
    path = '/Users/zen/Downloads/dirc/13'
    relativeFiles = get_dir(path)
    total = len(relativeFiles)
    print("共有 %d 个文件" % total)
    for file in relativeFiles:
        oldABS = '\"' + path + '/' + file + '\"'
        newABS = '\"' + path + '/' + replace(file) + '\"'
        print("旧文件名是: %s,新文件名是%s" % (oldABS, newABS))
        renameCommand(oldABS,newABS)
