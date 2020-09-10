# 重命名文件替换中文符号
# 在macOS/Linux下正常运行
import os


# 获取文件名(列表)

def getFilesName(path):
    fileList = os.listdir(path)
    count = len(fileList)
    print("共有%d个文件"%count)
    return fileList


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
    # print("生成的文件名是 %s"%oldName)
    newName = oldName
    return newName


if __name__ == '__main__':
    path = '/Volumes/Samsung/ts/'
    names = getFilesName(path=path)
    count=0
    for name in names:
        count+=1
        print('获得的旧文件名 %s' % name)
        print(type(name))
        new = replace(name)
        print('获得的新文件名 %s' % new)
        print(type(new))
        # os.rename(name,new)
        os.rename(os.path.join(path, name), os.path.join(path, new))
        print("正在处理第 %d 个文件"%count)
# 有时间尝试 try catch