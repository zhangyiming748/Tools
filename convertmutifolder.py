# 转换目录下多个子文件夹
import os
import time

def getFolders(dir):
    path = os.listdir(dir)
    for p in path:
        fullPath = dir + p
        if os.path.isdir(fullPath):
            # print(p)
            subfolder = fullPath
            print('子文件夹: %s' % subfolder)
            files = getFiles(subfolder)
            webp2jpg(files)


def getFiles(subfolder):
    files = os.listdir(subfolder)
    print('子文件夹的文件列表: %s' % files)
    f = []
    for file in files:
        print('file is %s' % file)
        full = os.path.join(subfolder, file)
        # print('转换前的全路径: %s' % full)
        f.append(full)

        print(type(full))
    print(type(f))
    return f


def webp2jpg(files):
    for file in files:
        print('ffmpeg输入的文件是 %s'%file)
        ext = file.split('.')
        if ext[-1] == 'webp':
            print(file + " :is a webp file")
            prefix = 'ffmpeg -i '
            # file
            thisTime=str(time.time())
            suffix = ' ' + file + thisTime +'.png'
            command = prefix + file + suffix
            print("命令: %s" % command)
            os.system(command)
            deleteWebp='rm '+file
            os.system(deleteWebp)


if __name__ == '__main__':
    root = r'/Users/zen/Pictures/Sticker/'
    getFolders(root)
