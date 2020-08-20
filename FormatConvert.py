# 利用ffmpeg实现文件批量转换
# 先获取工作目录文件列表的全路径,然后将获取到符合条件的文件名拼接命令执行
import os


def webp2png(files):
    for file in files:
        ext = file.split('.')
        if ext[-1] == 'webp':
            print(file + " :is a webp file")
            prefix = 'ffmpeg -i '
            # file
            suffix = ' ' + file +'.png'
            command = prefix + file + suffix
            print("命令: %s" % command)
            os.system(command)


def get_file_list() -> list:
    dirPath = os.getcwd()
    print("当前的工作目录是: %s" % dirPath)
    files = os.listdir(dirPath)
    print("目录下的文件列表: %s" % files)
    absPath = []
    for file in files:
        # print(file)
        # 拼接全路径
        f = os.path.join(dirPath, file)
        absPath.append(f)
    # print(absPath)
    return absPath


if __name__ == "__main__":
    files = get_file_list()
    webp2png(files)
