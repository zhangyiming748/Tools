# 删除macOS系统读取文件后自动生成的挂载点文件
import os

# 遍历文件
def getFiles(dir):
    fileList = os.listdir(dir)
    print(fileList)
    return fileList

# 获得前缀判断删除
def judge(file):
    print('%s 是挂载点文件,删除' % file)
    os.remove(os.path.join(path, file))

if __name__ == '__main__':
    path = r'C:\Users\zen\PycharmProjects\delete4K\testfile\\'
    files = getFiles(path)
    for file in files:
        judge(file)
    print("Done!")