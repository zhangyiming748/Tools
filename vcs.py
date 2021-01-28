import sys
import time


## 简便的记录信息工具
def vcs(s):
    prefix = timeNow()
    with open('./log.txt', mode='a', encoding='utf-8')as f:
        f.write(prefix)
        f.write(' ')
        for suffix in s:
            f.write(str(suffix))
            f.write(' ')
        f.write('\n')


def timeNow():
    MyTime = str(time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time())))
    return MyTime


if __name__ == '__main__':
    if len(sys.argv) <= 1:
        print('参数输入要添加的信息')
        sys.exit(-1)
    words = sys.argv[1:]
    vcs(words)
    print(len(sys.argv))
