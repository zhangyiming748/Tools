## 探测当前目录下文本文件的编码,转换为UTF-8,并删除原文件
## 需要拷贝到文件同级文件夹
from chardet.universaldetector import UniversalDetector
import os


# 探测当前文件编码
def encoding_detector(filepath):
    detector = UniversalDetector()
    detector.reset()
    for each in open(filepath, 'rb'):
        detector.feed(each)
        if detector.done:
            break
    detector.close()
    file_encoding = detector.result['encoding']
    confidence = detector.result['confidence']
    # print(file_encoding, confidence)
    # print(type(file_encoding))
    return file_encoding


# 转换为新文件
def convert_encoding(fp, old):
    f = open(fp, 'r', encoding=old, errors='replace')
    text = f.read()
    f.close()
    # names = fp.split('.')
    # # print(names)
    # prefix = names[0]
    # suffix = names[1]
    # new = prefix + 'utf8.' + suffix
    new = replace(fp)
    delete_old_file(fp)
    # if new==fp:
    #     new = 'rename' + new
    fi = open(new, 'w', encoding='utf-8', errors='replace')
    fi.write(text)
    fi.close()


# 获取目录下文件
def get_dir(dir):
    multi_files = os.listdir(dir)
    files = []
    for multi in multi_files:
        names = multi.split('.')
        # print(names)
        prefix = names[0]
        suffix = names[1]
        if suffix == 'm3u' and len(prefix)>0:
            files.append(multi)
        else:
            print('跳过 %s 文件' % multi)
    return files


# 删除旧文件
def delete_old_file(path_file):
    os.remove(path_file)


# 替换中文符号生成新文件名
def replace(oldName):
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
    oldName = oldName.replace('《', '')
    oldName = oldName.replace('》', '')
    oldName = oldName.replace('【', '[')
    oldName = oldName.replace('】', ']')
    oldName = oldName.replace('、', '\\')
    oldName = oldName.replace('～', '~')
    # print("生成的文件名是 %s"%oldName)
    newName = oldName
    return newName


if __name__ == '__main__':
    # fp = '/Users/zen/PycharmProjects/convertCode/name.txt'
    # ret = encoding_detector(fp)
    # print(ret, type(ret))
    # convert_encoding(fp, ret)
    dirfiles = get_dir('./')
    count = 0
    for dirfile in dirfiles:
        print('共 %d 个文件' % len(dirfiles))
        count += 1
        print('正在处理第 %d 个文件' % count)
        convert_encoding(dirfile, encoding_detector(dirfile))
        # delete_old_file(dirfile)
